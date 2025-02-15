package service

import (
	"errors"
	"fmt"
	"supermoney/models"
	"supermoney/repositories"
	"sync"
	"time"
)

type ConsumerService struct {
	topicRepo    *repositories.TopicRepository
	consumerRepo *repositories.ConsumerRepository
}

func NewConsumerService(topicRepo *repositories.TopicRepository, consumerRepo *repositories.ConsumerRepository) *ConsumerService {
	return &ConsumerService{
		topicRepo:    topicRepo,
		consumerRepo: consumerRepo,
	}
}

func (s *ConsumerService) RegisterConsumer(consumerName, topicName string) error {
	topic, err := s.topicRepo.GetTopic(topicName)
	if err != nil {
		return err
	}

	consumer := models.NewConsumer(consumerName, topic)
	return s.consumerRepo.AddConsumer(consumer)
}

func (s *ConsumerService) ConsumeMessages(consumerName string) error {
	consumer, err := s.consumerRepo.GetConsumer(consumerName)
	if err != nil {
		return err
	}

	topic := consumer.GetTopic()
	if topic == nil {
		return errors.New("consumer is not subscribed to a topic")
	}

	messages := topic.GetMessages()
	consumerOffset := consumer.GetOffset()

	if consumerOffset >= len(messages) {
		fmt.Printf("Consumer %s: No new messages to consume.\n", consumerName)
		return nil
	}

	var wg sync.WaitGroup
	for i := consumerOffset; i < len(messages); i++ {
		wg.Add(1)
		go func(msg models.Message, index int) {
			defer wg.Done()
			fmt.Printf("Consumer %s consumed message: %s (Timestamp: %s)\n",
				consumer.GetName(), msg.Content, msg.Timestamp.Format(time.RFC3339))
			consumer.SetOffset(index + 1)
		}(*messages[i], i)
	}

	wg.Wait()
	return nil
}

func (s *ConsumerService) ResetConsumerOffset(consumerName string, newOffset int) error {
	consumer, err := s.consumerRepo.GetConsumer(consumerName)
	if err != nil {
		return err
	}

	if newOffset < 0 || newOffset >= len(consumer.GetTopic().GetMessages()) {
		return errors.New("invalid offset value")
	}

	consumer.SetOffset(newOffset)
	fmt.Printf("Consumer %s reset offset to %d\n", consumer.GetName(), newOffset)
	return nil
}

func (s *ConsumerService) RemoveConsumer(consumerName string) error {
	return s.consumerRepo.RemoveConsumer(consumerName)
}

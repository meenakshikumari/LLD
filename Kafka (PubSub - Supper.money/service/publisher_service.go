package service

import (
	"errors"
	"fmt"
	"supermoney/models"
	"supermoney/repositories"
	"sync"
)

type PublisherService struct {
	topicRepo     *repositories.TopicRepository
	publisherRepo *repositories.PublisherRepository
}

func NewPublisherService(topicRepo *repositories.TopicRepository, publisherRepo *repositories.PublisherRepository) *PublisherService {
	return &PublisherService{
		topicRepo:     topicRepo,
		publisherRepo: publisherRepo,
	}
}

func (s *PublisherService) RegisterPublisher(name, topicName string) error {
	topic, err := s.topicRepo.GetTopic(topicName)
	if err != nil {
		return err
	}

	publisher := models.NewPublisher(name, topic)
	return s.publisherRepo.AddPublisher(publisher)
}

func (s *PublisherService) PublishMessage(publisherName, message string) error {
	publisher, err := s.publisherRepo.GetPublisher(publisherName)
	if err != nil {
		return err
	}

	topic := publisher.GetTopic()
	if topic == nil {
		return errors.New("publisher is not assigned to a topic")
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := topic.AddMessage(message)
		if err != nil {
			return
		}
		fmt.Printf("Publisher %s published message: %s\n", publisher.GetName(), message)
	}()

	wg.Wait()
	return nil
}

func (s *PublisherService) RemovePublisher(publisherName string) error {
	//TODO functionality for future
	return s.publisherRepo.RemovePublisher(publisherName)
}

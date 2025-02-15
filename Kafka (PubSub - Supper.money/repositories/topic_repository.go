package repositories

import (
	"errors"
	"supermoney/models"
	"sync"
	"time"
)

type TopicRepository struct {
	Topics map[string]*models.Topic
	Mutex  sync.Mutex
}

func NewTopicRepository() *TopicRepository {
	return &TopicRepository{
		Topics: make(map[string]*models.Topic),
	}
}

func (r *TopicRepository) CreateTopic(name string, retentionTime time.Duration) (*models.Topic, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	if _, exists := r.Topics[name]; exists {
		return nil, errors.New("topic already exists")
	}
	topic := models.NewTopic(name, retentionTime)
	r.Topics[name] = topic
	return topic, nil
}

func (r *TopicRepository) GetTopic(name string) (*models.Topic, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	topic, exists := r.Topics[name]
	if !exists || topic.Deleted {
		return nil, errors.New("topic not found or deleted")
	}

	return topic, nil
}

func (r *TopicRepository) DeleteTopic(name string) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	topic, err := r.GetTopic(name)
	if err != nil {
		return err
	}

	topic.DeleteTopic()

	for _, consumer := range topic.Consumers {
		consumer.HandleTopicDeleted()
	}

	delete(r.Topics, name)
	return nil
}

func (r *TopicRepository) AddConsumerToTopic(topicName string, consumer *models.Consumer) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	topic, err := r.GetTopic(topicName)
	if err != nil {
		return err
	}

	topic.AddConsumer(consumer)
	return nil
}

func (r *TopicRepository) GetAllTopics() []*models.Topic {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	var allTopics []*models.Topic
	for _, topic := range r.Topics {
		allTopics = append(allTopics, topic)
	}
	return allTopics
}

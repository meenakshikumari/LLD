package service

import (
	"supermoney/models"
	"supermoney/repositories"
	"time"
)

type TopicService struct {
	Repo *repositories.TopicRepository
}

func NewTopicService(repo *repositories.TopicRepository) *TopicService {
	return &TopicService{Repo: repo}
}

func (s *TopicService) CreateTopic(name string, retentionTimeInSeconds int) (*models.Topic, error) {
	retentionDuration := time.Duration(retentionTimeInSeconds) * time.Second
	return s.Repo.CreateTopic(name, retentionDuration)
}

func (s *TopicService) DeleteTopic(name string) error {
	return s.Repo.DeleteTopic(name)
}

func (s *TopicService) GetTopic(name string) (*models.Topic, error) {
	return s.Repo.GetTopic(name)
}

func (s *TopicService) AddConsumerToTopic(topicName string, consumer *models.Consumer) error {
	return s.Repo.AddConsumerToTopic(topicName, consumer)
}

func (s *TopicService) CleanupExpiredMessages() {
	for _, topic := range s.Repo.GetAllTopics() {
		topic.CleanupOldMessages()
	}
}

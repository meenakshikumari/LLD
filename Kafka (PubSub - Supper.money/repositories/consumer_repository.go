package repositories

import (
	"errors"
	"supermoney/models"
	"sync"
)

type ConsumerRepository struct {
	Consumers map[string]*models.Consumer
	Mutex     sync.Mutex
}

func NewConsumerRepository() *ConsumerRepository {
	return &ConsumerRepository{
		Consumers: make(map[string]*models.Consumer),
	}
}

func (r *ConsumerRepository) AddConsumer(consumer *models.Consumer) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	if consumer == nil {
		return errors.New("consumer cannot be nil")
	}

	_, exists := r.Consumers[consumer.Name]
	if exists {
		return errors.New("consumer already exists")
	}

	r.Consumers[consumer.Name] = consumer
	return nil
}

func (r *ConsumerRepository) GetConsumer(consumerName string) (*models.Consumer, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	consumer, exists := r.Consumers[consumerName]
	if !exists {
		return nil, errors.New("consumer not found")
	}

	return consumer, nil
}

func (r *ConsumerRepository) RemoveConsumer(consumerName string) error {
	//TODO Can be implemented
	delete(r.Consumers, consumerName)
	return nil
}

func (r *ConsumerRepository) UpdateOffset(consumerName string, offset int) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	consumer, err := r.GetConsumer(consumerName)
	if err != nil {
		return err
	}

	consumer.SetOffset(offset)
	return nil
}

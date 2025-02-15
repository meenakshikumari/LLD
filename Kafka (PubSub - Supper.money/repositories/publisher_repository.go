package repositories

import (
	"errors"
	"supermoney/models"
	"sync"
)

type PublisherRepository struct {
	publishers map[string]*models.Publisher
	mutex      sync.Mutex
}

func NewPublisherRepository() *PublisherRepository {
	return &PublisherRepository{
		publishers: make(map[string]*models.Publisher),
	}
}

func (r *PublisherRepository) AddPublisher(publisher *models.Publisher) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if publisher == nil {
		return errors.New("publisher cannot be nil")
	}

	if _, exists := r.publishers[publisher.GetName()]; exists {
		return errors.New("publisher already exists")
	}

	r.publishers[publisher.Name] = publisher
	return nil
}

func (r *PublisherRepository) GetPublisher(name string) (*models.Publisher, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	publisher, exists := r.publishers[name]
	if !exists {
		return nil, errors.New("publisher not found")
	}

	return publisher, nil
}

func (r *PublisherRepository) RemovePublisher(name string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.publishers[name]; !exists {
		return errors.New("publisher not found")
	}

	delete(r.publishers, name)
	return nil
}

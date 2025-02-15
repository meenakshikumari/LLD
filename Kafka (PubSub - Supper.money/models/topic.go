package models

import (
	"errors"
	"sync"
	"time"
)

type Topic struct {
	Name          string
	Messages      []*Message
	Consumers     map[string]*Consumer
	Publishers    map[string]*Publisher
	Mutex         sync.Mutex
	RetentionTime time.Duration
	CreatedAt     time.Time
	Deleted       bool
	DeletedAt     time.Time
}

func NewTopic(name string, retentionTime time.Duration) *Topic {
	return &Topic{
		Name:          name,
		Messages:      []*Message{},
		Consumers:     make(map[string]*Consumer),
		Publishers:    make(map[string]*Publisher),
		RetentionTime: retentionTime,
		CreatedAt:     time.Now(),
		Deleted:       false,
	}
}

func (t *Topic) AddMessage(msg string) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	if t.Deleted {
		return errors.New("cannot add message, topic is deleted")
	}

	t.Messages = append(t.Messages, NewMessage(msg))

	t.CleanupOldMessages() // TODO: we can even do this periodically instead of this func. But doing here for MVP
	return nil
}

func (t *Topic) GetMessages() []*Message {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	return t.Messages
}

func (t *Topic) CleanupOldMessages() {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	expirationTime := time.Now().Add(-t.RetentionTime)
	var filteredMessages []*Message

	for _, msg := range t.Messages {
		if msg.Timestamp.After(expirationTime) {
			filteredMessages = append(filteredMessages, msg)
		}
	}

	t.Messages = filteredMessages
}

func (t *Topic) DeleteTopic() {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	t.Deleted = true
	t.DeletedAt = time.Now()
	t.Messages = nil
}

func (t *Topic) AddConsumer(consumer *Consumer) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Consumers[consumer.Name] = consumer
}

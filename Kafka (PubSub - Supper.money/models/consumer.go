package models

import (
	"errors"
	"fmt"
	"sync"
)

type Consumer struct {
	Name    string
	Topic   *Topic
	Offset  int
	Mutex   sync.Mutex
	Stopped bool
}

func NewConsumer(name string, topic *Topic) *Consumer {
	consumer := &Consumer{Name: name, Topic: topic, Offset: 0, Stopped: false}
	topic.Consumers[name] = consumer
	return consumer
}

func (c *Consumer) GetName() string {
	return c.Name
}

func (c *Consumer) GetTopic() *Topic {
	return c.Topic
}

func (c *Consumer) GetOffset() int {
	return c.Offset
}

func (c *Consumer) SetOffset(offset int) {
	c.Offset = offset
}

func (c *Consumer) Consume() (string, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.Stopped {
		return "", errors.New("consumer has stopped consuming due to topic deletion")
	}

	if c.Offset >= len(c.Topic.Messages) {
		return "", errors.New("no new messages")
	}

	message := c.Topic.Messages[c.Offset]
	c.Offset++
	return message.Content, nil
}

func (c *Consumer) HandleTopicDeleted() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	fmt.Printf("Consumer %s stopped: Topic %s was deleted.\n", c.Name, c.Topic.Name)
	c.Stopped = true
}

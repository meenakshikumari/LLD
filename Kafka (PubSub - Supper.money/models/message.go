package models

import "time"

type Message struct {
	Content   string
	Timestamp time.Time
}

func NewMessage(content string) *Message {
	return &Message{Content: content, Timestamp: time.Now()}
}

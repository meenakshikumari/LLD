package models

type Publisher struct {
	Name  string
	Topic *Topic
}

func NewPublisher(name string, topic *Topic) *Publisher {
	return &Publisher{Name: name, Topic: topic}
}

func (p *Publisher) GetName() string {
	return p.Name
}

func (p *Publisher) SetName(name string) {
	p.Name = name
}

func (p *Publisher) GetTopic() *Topic {
	return p.Topic
}

func (p *Publisher) SetTopic(topic *Topic) {
	p.Topic = topic
}

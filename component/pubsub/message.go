package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	channel   Topic
	data      any
	createdAt time.Time
}

func NewMessage(data any) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (evt *Message) String() string {
	return fmt.Sprintf("Message %s", evt.channel)
}

func (evt *Message) Channel() Topic {
	return evt.channel
}

func (evt *Message) SetChannel(channel Topic) {
	evt.channel = channel
}

func (evt *Message) Data() any {
	return evt.data
}

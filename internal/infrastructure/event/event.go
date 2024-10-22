package event

import (
	"time"
)

type Event struct {
	Type      string
	Timestamp time.Time
	Data      interface{}
}

type EventBus struct {
	Subscribers map[string][]chan<- Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		Subscribers: make(map[string][]chan<- Event),
	}
}

// Subscribe adds a new subscriber for a given event type
func (eb *EventBus) Subscribe(eventType string, subscriber chan<- Event) {
	eb.Subscribers[eventType] = append(eb.Subscribers[eventType], subscriber)
}

// Publish sends an event to all subscribers of a given event type
func (eb *EventBus) Publish(event Event) {
	subscribers := eb.Subscribers[event.Type]
	for _, subscriber := range subscribers {
		subscriber <- event
	}
}

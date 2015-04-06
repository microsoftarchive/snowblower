package main

// EventPublisher ...
type EventPublisher interface {
	publish(event *Event)
}

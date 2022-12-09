package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/variant64/server/pkg/bus"
)

// pub represents the models message bus.
type pub[T any] interface {
	Publish(topic uuid.UUID, message T) error
}

// UpdatePublisher represents the models message publisher.
// Publisher sends messages to the message bus.
type UpdatePublisher[T any] struct {
	Pub      pub[T]
	EntityID uuid.UUID
}

// Publish pushes a message to message bus.
func (u *UpdatePublisher[T]) Publish(update T) {
	u.Pub.Publish(u.EntityID, update)
}

// NewUpdatePub returns a new message publisher for the given topic.
func NewUpdatePub[T any](id uuid.UUID, p *bus.Bus[T]) (*UpdatePublisher[T], error) {
	err := p.NewTopic(id)
	if err != nil {
		return nil, err
	}

	u := &UpdatePublisher[T]{
		Pub:      bus.NewPub(p),
		EntityID: id,
	}

	return u, nil
}

// EventWriter write messages to downstream, for example performing
// WriteMessage() on a websocket connection.
type EventWriter interface {
	WriteMessage(messageType int, data []byte) error
}

// MockEventWriter creates a mock event writer,
// event writers writes message to downstream.
type MockEventWriter struct {
	SentMessages []string
}

// WriteMessage handles the incoming message.
func (m *MockEventWriter) WriteMessage(messageType int, data []byte) error {
	m.SentMessages = append(m.SentMessages, string(data))

	return nil
}

// NewMockEventWriter creates a new mock message publisher.
func NewMockEventWriter() *MockEventWriter {
	return &MockEventWriter{
		SentMessages: make([]string, 0),
	}
}

// subscriber represent the models subscription object for message bus.
type subscriber[T any] struct {
	eventWriter EventWriter
}

// NewMessageSubscriber returns a new subscriber for message bus.
func NewMessageSubscriber[T any](eventWriter EventWriter) *subscriber[T] {
	return &subscriber[T]{
		eventWriter: eventWriter,
	}
}

// OnMessage handles subscribers incoming messages from message bus.
func (s *subscriber[T]) OnMessage(update T) error {
	message, err := json.Marshal(update)
	if err != nil {
		return err
	}
	s.eventWriter.WriteMessage(1, message)
	return nil
}

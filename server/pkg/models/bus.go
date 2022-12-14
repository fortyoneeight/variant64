package models

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/variant64/server/pkg/bus"
)

// UpdateMessage represents an update message.
type UpdateMessage[T any] struct {
	Channel string     `json:"channel"`
	Data    T          `json:"data"`
	Type    UpdateType `json:"type"`
}

// NewUpdateBus creates a new update message bus.
func NewUpdateBus[T any]() *bus.Bus[UpdateMessage[T]] {
	return bus.NewBus[UpdateMessage[T]]([]uuid.UUID{})
}

// Publish pushes a message to message UpdatePublisher.
func (u *UpdatePublisher[T]) Publish(update UpdateMessage[T]) {
	u.Pub.Publish(u.EntityID, update)
}

// pub represents the models message bus.
type pub[T any] interface {
	Publish(topic uuid.UUID, message T) error
}

// UpdatePublisher represents the models message publisher.
// Publisher sends messages to the message bus.
type UpdatePublisher[T any] struct {
	Pub      pub[UpdateMessage[T]]
	EntityID uuid.UUID
}

// NewUpdatePub returns a new message publisher for the given topic.
func NewUpdatePub[T any](id uuid.UUID, p *bus.Bus[UpdateMessage[T]]) (*UpdatePublisher[T], error) {
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

// LastMessage returns the last sent message.
func (m *MockEventWriter) LastMessage() string {
	return m.SentMessages[len(m.SentMessages)-1]
}

// NewMockEventWriter creates a new mock message publisher.
func NewMockEventWriter() *MockEventWriter {
	return &MockEventWriter{
		SentMessages: make([]string, 0),
	}
}

// subscriber represent the models subscription object for message bus.
type subscriber[T any] struct {
	channel     string
	eventWriter EventWriter
}

// NewMessageSubscriber returns a new subscriber for message bus.
func NewMessageSubscriber[T any](channel string, eventWriter EventWriter) *subscriber[T] {
	return &subscriber[T]{
		channel:     channel,
		eventWriter: eventWriter,
	}
}

// OnMessage handles subscribers incoming messages from message bus.
func (s *subscriber[any]) OnMessage(update any) error {
	message, err := json.Marshal(update)
	if err != nil {
		return err
	}
	s.eventWriter.WriteMessage(1, message)
	return nil
}

type UpdateType int32

const (
	UpdateType_NONE UpdateType = iota
	UpdateType_SNAPSHOT
	UpdateType_DELTA
)

func (t UpdateType) String() string {
	switch t {
	case UpdateType_NONE:
		return "none"
	case UpdateType_DELTA:
		return "delta"
	case UpdateType_SNAPSHOT:
		return "snapshot"
	default:
		return ""
	}
}

func (t UpdateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// Subscribe subscribes to the bus.Bus.
func Subscribe[T any](b *bus.Bus[UpdateMessage[T]], topic uuid.UUID, channel string, pub EventWriter) {
	b.Subscribe(topic, NewMessageSubscriber[UpdateMessage[T]](channel, pub))
}

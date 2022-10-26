package bus

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// Bus represents a set of topics and their subscribers.
type Bus[T any] struct {
	topics map[uuid.UUID]chan T
	mux sync.RWMutex

	subs map[uuid.UUID][]Sub[T]
}

func NewBus[T any](topics []uuid.UUID) *Bus[T] {
	bus := &Bus[T]{
		topics: make(map[uuid.UUID]chan T),
		mux: sync.RWMutex{},
		subs: make(map[uuid.UUID][]Sub[T]),
	}
	for _, t := range topics {
		bus.topics[t] = make(chan T)
	}
	return bus
}

// Sub represents a subscriber to a Bus topic.
type Sub[T any] interface {
	OnMessage(t T) error
}

// Pub sends messages to a topic on a Bus.
type Pub[T any] struct {
	bus *Bus[T]
}

// Start starts the Bus publishing routine.
func (b *Bus[T]) Start() {
	for topic, c := range b.topics {
		go b.publish(topic, c)
	}
}

// publish reads messages off the topicChan and sends them to required Subs.
func (b *Bus[T]) publish(topic uuid.UUID, topicChan chan T) {
	for {
		msg := <- topicChan
		b.mux.RLock()
		for _, s := range b.subs[topic] {
			s.OnMessage(msg)
		}
		b.mux.RUnlock()
	}
}

// Subscribe adds the Sub to the list of Subs for the provided topic.
func (b *Bus[T]) Subscribe(topic uuid.UUID, s Sub[T]) error {
	b.mux.Lock()
	defer b.mux.Unlock()

	if _, ok := b.topics[topic]; ok {
		if _, ok := b.subs[topic]; !ok {
			b.subs[topic] = make([]Sub[T], 0)
		}
		b.subs[topic] = append(b.subs[topic], s)
		return nil
	}
	return fmt.Errorf("topic does not exist: %s", topic)
}

// Publish sends a message on the provided topic.
func (p *Pub[T]) Publish(topic uuid.UUID, message T) error {
	p.bus.mux.RLock()
	defer p.bus.mux.RUnlock()

	if t, ok := p.bus.topics[topic]; ok {
		t <- message
		return nil
	}
	return fmt.Errorf("topic does not exist: %s", topic)
}

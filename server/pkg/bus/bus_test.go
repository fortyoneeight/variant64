package bus

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBus(t *testing.T) {
	topic1 := uuid.New()
	topic2 := uuid.New()

	testcases := []struct {
		name           string
		bus            *Bus[string]
		subs           map[uuid.UUID][]*mockSub[string]
		messages       map[uuid.UUID][]string
		subscribeError error
		publishError   error
	}{
		{
			name: "One topic, one sub, one message.",
			bus:  NewBus[string]([]uuid.UUID{topic1}),
			subs: map[uuid.UUID][]*mockSub[string]{
				topic1: {
					&mockSub[string]{},
				},
			},
			messages: map[uuid.UUID][]string{
				topic1: {
					"test1",
				},
			},
			subscribeError: nil,
			publishError:   nil,
		},
		{
			name: "One topic, two subs, two messages.",
			bus:  NewBus[string]([]uuid.UUID{topic1}),
			subs: map[uuid.UUID][]*mockSub[string]{
				topic1: {
					&mockSub[string]{},
					&mockSub[string]{},
				},
			},
			messages: map[uuid.UUID][]string{
				topic1: {
					"test1",
					"test2",
				},
			},
			subscribeError: nil,
			publishError:   nil,
		},
		{
			name: "Two topics, two subs per topic, two messages per topic.",
			bus:  NewBus[string]([]uuid.UUID{topic1, topic2}),
			subs: map[uuid.UUID][]*mockSub[string]{
				topic1: {
					&mockSub[string]{},
					&mockSub[string]{},
				},
				topic2: {
					&mockSub[string]{},
					&mockSub[string]{},
				},
			},
			messages: map[uuid.UUID][]string{
				topic1: {
					"test1",
					"test2",
				},
				topic2: {
					"test3",
					"test4",
				},
			},
			subscribeError: nil,
			publishError:   nil,
		},
		{
			name: "No topics, one sub, subscribe error.",
			bus:  NewBus[string]([]uuid.UUID{}),
			subs: map[uuid.UUID][]*mockSub[string]{
				topic1: {
					&mockSub[string]{},
				},
			},
			messages:       map[uuid.UUID][]string{},
			subscribeError: fmt.Errorf("topic does not exist: %s", topic1),
			publishError:   nil,
		},
		{
			name: "No topics, no pubs, publish error.",
			bus:  NewBus[string]([]uuid.UUID{}),
			subs: map[uuid.UUID][]*mockSub[string]{},
			messages: map[uuid.UUID][]string{
				topic1: {
					"test1",
				},
			},
			subscribeError: nil,
			publishError:   fmt.Errorf("topic does not exist: %s", topic1),
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.bus.Start()
			pub := &Pub[string]{bus: testcase.bus}

			for topic, subs := range testcase.subs {
				for _, s := range subs {
					error := testcase.bus.Subscribe(topic, s)
					if testcase.subscribeError != nil {
						assert.Equal(t, testcase.subscribeError, error)
					}
				}
			}

			for topic, messages := range testcase.messages {
				for _, m := range messages {
					error := pub.Publish(topic, m)
					if testcase.publishError != nil {
						assert.Equal(t, testcase.publishError, error)
					}
				}
			}

			for topic, subs := range testcase.subs {
				for _, s := range subs {
					for _, message := range s.receivedMessages {
						assert.Contains(t, testcase.messages[topic], message)
					}
				}
			}
		})
	}
}

type mockSub[T any] struct {
	receivedMessages []T
}

func (m *mockSub[T]) OnMessage(message T) error {
	m.receivedMessages = append(m.receivedMessages, message)
	return nil
}

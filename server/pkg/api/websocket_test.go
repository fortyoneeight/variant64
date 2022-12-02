package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/pkg/api/command"
)

func TestReadAndHandleMessages(t *testing.T) {
	testcases := []struct {
		description        string
		mockCommandHandler *mockCommandHandler
		messages           []string
		expectedCommands   []command.Command
	}{
		{
			"No commands sent.",
			&mockCommandHandler{
				receivedCommands: make([]command.Command, 0),
			},
			[]string{},
			[]command.Command{},
		},
		{
			"One command sent.",
			&mockCommandHandler{
				receivedCommands: make([]command.Command, 0),
			},
			[]string{"{\"command\":\"subscribe\"}"},
			[]command.Command{
				command.Command{Command: command.Subscribe},
			},
		},
		{
			"Multiple commands sent.",
			&mockCommandHandler{
				receivedCommands: make([]command.Command, 0),
			},
			[]string{
				"{\"command\":\"subscribe\"}",
				"{\"command\":\"unsubscribe\"}",
				"{\"command\":\"subscribe\"}",
				"{\"command\":\"unsubscribe\"}",
			},
			[]command.Command{
				command.Command{Command: command.Subscribe},
				command.Command{Command: command.Unsubscribe},
				command.Command{Command: command.Subscribe},
				command.Command{Command: command.Unsubscribe},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			// Setup mock server with mockCommandHandler.
			server := httptest.NewServer(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						conn, err := upgrader.Upgrade(w, r, nil)
						if err != nil {
							return
						}
						defer conn.Close()

						readAndHandleMessages(conn, tc.mockCommandHandler)
					}),
			)

			// Connect to mock server.
			serverURL := "ws" + strings.TrimPrefix(server.URL, "http")
			conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
			if err != nil {
				t.Error(err)
			}

			// Send messages.
			for _, message := range tc.messages {
				conn.WriteMessage(websocket.TextMessage, []byte(message))
			}

			// Wait for processing and check expected.
			for i := 0; i < 5; i += 1 {
				if len(tc.expectedCommands) > len(tc.mockCommandHandler.receivedCommands) {
					time.Sleep(1 * time.Second)
				}
			}
			assert.Equal(t, tc.expectedCommands, tc.mockCommandHandler.receivedCommands)

			server.Close()
			conn.Close()
		})
	}
}

type mockCommandHandler struct {
	receivedCommands []command.Command
}

func (m *mockCommandHandler) HandleCommand(command command.Command, message []byte) error {
	m.receivedCommands = append(m.receivedCommands, command)
	return nil
}

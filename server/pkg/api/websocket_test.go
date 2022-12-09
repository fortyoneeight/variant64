package api

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/pkg/models/game"
	"github.com/variant64/server/pkg/models/room"
)

func TestReadAndHandleMessages(t *testing.T) {
	testcases := []struct {
		description               string
		mockCommandHandler        *mockCommandHandler
		messages                  []string
		expectedCommands          []WebSocketRequest
		expectedChannelRequestMap map[string]int
	}{
		{
			"all commands.",
			newMockCommandHandler(),
			[]string{
				"{\"channel\":\"game\", \"command\":\"subscribe\"}",
				"{\"channel\":\"game\", \"command\":\"unsubscribe\"}",
				"{\"channel\":\"room\", \"command\":\"subscribe\"}",
				"{\"channel\":\"room\", \"command\":\"unsubscribe\"}",
			},
			[]WebSocketRequest{
				{Channel: game.MessageChannel, Command: game.GameSubscribe},
				{Channel: game.MessageChannel, Command: game.GameUnsubscribe},
				{Channel: room.MessageChannel, Command: room.RoomSubscribe},
				{Channel: room.MessageChannel, Command: room.RoomUnsubscribe},
			},
			map[string]int{
				game.MessageChannel: 2,
				room.MessageChannel: 2,
			},
		},
		{
			"No commands sent.",
			newMockCommandHandler(),
			[]string{},
			[]WebSocketRequest{},
			make(map[string]int),
		},
		{
			"One command sent.",
			newMockCommandHandler(),
			[]string{"{\"channel\":\"game\", \"command\":\"subscribe\"}"},
			[]WebSocketRequest{
				{Command: game.GameSubscribe},
			},
			map[string]int{
				game.MessageChannel: 1,
			},
		},
		{
			"Multiple commands sent.",
			newMockCommandHandler(),
			[]string{
				"{\"channel\":\"game\", \"command\":\"subscribe\"}",
				"{\"channel\":\"game\", \"command\":\"unsubscribe\"}",
				"{\"channel\":\"game\", \"command\":\"subscribe\"}",
				"{\"channel\":\"game\", \"command\":\"unsubscribe\"}",
			},
			[]WebSocketRequest{
				{Channel: game.MessageChannel, Command: game.GameSubscribe},
				{Channel: game.MessageChannel, Command: game.GameUnsubscribe},
				{Channel: game.MessageChannel, Command: game.GameSubscribe},
				{Channel: game.MessageChannel, Command: game.GameUnsubscribe},
			},
			map[string]int{
				game.MessageChannel: 4,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			// Setup mock server with mockCommandHandler.
			server := httptest.NewServer(NewWebSocketHandleFunc(tc.mockCommandHandler))

			// Connect to mock server.
			wsServer := NewWebSocketServer(server.URL)

			// Send messages.
			for _, message := range tc.messages {
				wsServer.WriteMessage(websocket.TextMessage, []byte(message))
			}

			// Wait for processing and check expected.
			for i := 0; i < 5; i += 1 {
				if len(tc.expectedCommands) > tc.mockCommandHandler.totalCommandsReceived {
					time.Sleep(1 * time.Second)
				}
			}
			assert.Equal(t, len(tc.expectedCommands), tc.mockCommandHandler.totalCommandsReceived)

			for channel, expectedChannelCount := range tc.expectedChannelRequestMap {
				assert.Equal(t, len(tc.mockCommandHandler.receivedCommandsMap[channel]), expectedChannelCount)
			}

			server.Close()
			wsServer.Close()
		})
	}
}

type mockCommandHandler struct {
	receivedCommandsMap   map[string][]WebSocketRequest
	totalCommandsReceived int
}

func (m *mockCommandHandler) HandleCommand(command WebSocketRequest) error {
	m.receivedCommandsMap[command.Channel] = append(m.receivedCommandsMap[command.Channel], command)
	m.totalCommandsReceived++
	return nil
}

func newMockCommandHandler() *mockCommandHandler {
	return &mockCommandHandler{
		receivedCommandsMap: make(map[string][]WebSocketRequest),
	}
}

func TestChannelHandlerMap(t *testing.T) {
	wsHandler := &WSHandler{
		handlerMap: make(map[string]channelHandlerFunc),
	}
	// Setup mock server with mockCommandHandler.
	server := httptest.NewServer(NewWebSocketHandleFunc(wsHandler))

	// Connect to mock server.
	wsServer := NewWebSocketServer(server.URL)
	wsHandler.SetWebsocketConn(wsServer.Conn)

	RegisterChannelHandlers(wsHandler)

	assert.Equal(t, len(wsHandler.AvailableChannels()), 2)
}

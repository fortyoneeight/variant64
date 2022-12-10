package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/variant64/server/pkg/models"
	"github.com/variant64/server/pkg/models/game"
	"github.com/variant64/server/pkg/models/room"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// channelHandlerFunc is the handler for a specific channel.
type channelHandlerFunc func(conn models.EventWriter, command, body string) error

// WSHandler handles incoming WS messages from a client.
type WSHandler struct {
	conn       *websocket.Conn
	handlerMap map[string]channelHandlerFunc
}

// RegisterChannelHandler registers [channel, handler] to WSHandler.handlerMap.
func (h *WSHandler) RegisterChannelHandler(channel string, handleFunc channelHandlerFunc) {
	h.handlerMap[channel] = handleFunc
}

// AvailableChannels returns a list of available websocket channels.
func (h *WSHandler) AvailableChannels() []string {
	channels := []string{}

	for channel := range h.handlerMap {
		channels = append(channels, channel)
	}

	return channels
}

// SetWebsocketConn sets the connection.
func (h *WSHandler) SetWebsocketConn(conn *websocket.Conn) {
	h.conn = conn
}

// NewWSHandler creates and returns a new WSHandler.
func NewWSHandler(conn *websocket.Conn) *WSHandler {
	return &WSHandler{
		conn:       conn,
		handlerMap: make(map[string]channelHandlerFunc),
	}
}

// WebSocketRequest is the request structure for incoming web socket messages.
type WebSocketRequest struct {
	Channel string `json:"channel"`
	Command string `json:"command"`
	Body    string `json:"body"`
}

// commandHandler represents a handler of incoming client commands.
type commandHandler interface {
	HandleCommand(command WebSocketRequest) error
}

// HandleCommand handles a specific WebsocketRequest from a provided client message.
func (w *WSHandler) HandleCommand(command WebSocketRequest) error {
	handleFunc, ok := w.handlerMap[command.Channel]
	if !ok {
		return errors.New("invalid or missing channel command handler")
	}

	return handleFunc(w.conn, command.Command, command.Body)
}

// websocketHandler upgrades and handles an incoming websocket connection request.
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("error during connection upgrade: ", err)
		return
	}
	defer conn.Close()

	handler := NewWSHandler(conn)
	RegisterChannelHandlers(handler)
	readAndHandleMessages(conn, handler)
}

// readAndHandleMessages continuously reads client messages and handles them.
func readAndHandleMessages(conn *websocket.Conn, handler commandHandler) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("error during message reading: ", err)
			break
		}

		command := &WebSocketRequest{}
		err = json.Unmarshal(message, command)
		if err == nil {
			handler.HandleCommand(*command)
		}
	}
}

// RegisterChannelHandlers registers app websocket channels.
func RegisterChannelHandlers(h *WSHandler) {
	h.RegisterChannelHandler(game.MessageChannel, game.HandleCommand)
	h.RegisterChannelHandler(room.MessageChannel, room.HandleCommand)
}

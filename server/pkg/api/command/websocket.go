package command

import (
	"encoding/json"
	"errors"

	"github.com/gorilla/websocket"
	"github.com/variant64/server/pkg/models/game"
)

// WSHandler handles incoming WS messages from a client.
type WSHandler struct {
	conn *websocket.Conn
}

// NewWSHandler creates and returns a new WSHandler.
func NewWSHandler(conn *websocket.Conn) *WSHandler {
	return &WSHandler{
		conn: conn,
	}
}

// HandleCommand handles a specific Command from a provided client message.
func (w *WSHandler) HandleCommand(command Command, message []byte) error {
	switch command.Command {
	case GameSubscribe:
		commandGameSubscribe := &CommandGameSubscribe{}
		err := json.Unmarshal(message, commandGameSubscribe)
		if err != nil {
			return err
		}
		return w.handleGameSubscribe(commandGameSubscribe)
	default:
		return errors.New("invalid or missing command")
	}
}

// handleGameSubscribe handles a CommandSubscribe from a client.
func (w *WSHandler) handleGameSubscribe(command *CommandGameSubscribe) error {
	gameUpdateBus := game.GetGameUpdateBus()
	gameUpdateBus.Subscribe(command.GameID, &gameUpdateSubscriber{conn: w.conn})
	return nil
}

// gameUpdateSubscriber subscribes a websocket.Conn to entity.GameUpdates.
type gameUpdateSubscriber struct {
	conn *websocket.Conn
}

// OnMessage forwards entity.GameUpdates to the associated websocket.Conn.
func (g *gameUpdateSubscriber) OnMessage(update game.GameUpdate) error {
	message, err := json.Marshal(update)
	if err != nil {
		return nil
	}
	g.conn.WriteMessage(1, message)
	return nil
}

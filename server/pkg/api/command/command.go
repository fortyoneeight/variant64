package command

import (
	"github.com/google/uuid"
)

// APICommand enumerates the client command types.
type APICommand string

const (
	GameSubscribe   APICommand = "game_subscribe"
	GameUnsubscribe APICommand = "game_unsubscribe"

	RoomSubscribe   APICommand = "room_subscribe"
	RoomUnsubscribe APICommand = "room_unsubscribe"
)

// Command represents a top level command.
type Command struct {
	Command APICommand `json:"command"`
}

// CommandGameSubscribe represents a game subscribe command.
type CommandGameSubscribe struct {
	Command
	GameID uuid.UUID `json:"game_id"`
}

// CommandGameUnsubscribe represents an game unsubscribe command.
type CommandGameUnsubscribe struct {
	Command
	GameID uuid.UUID `json:"game_id"`
}

// CommandRoomSubscribe represents a game subscribe command.
type CommandRoomSubscribe struct {
	Command
	RoomID uuid.UUID `json:"room_id"`
}

// CommandRoomUnsubscribe represents an room unsubscribe command.
type CommandRoomUnsubscribe struct {
	Command
	RoomID uuid.UUID `json:"room_id"`
}

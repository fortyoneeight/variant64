package command

import (
	"github.com/google/uuid"
)

// APICommand enumerates the client command types.
type APICommand string

const (
	Subscribe   APICommand = "subscribe"
	Unsubscribe APICommand = "unsubscribe"
)

// Command represents a top level command.
type Command struct {
	Command APICommand `json:"command"`
}

// CommandSubscribe represents a subscribe command.
type CommandSubscribe struct {
	Command
	GameID uuid.UUID `json:"game_id"`
}

// CommandUnsubscribe represents an unsubscribe command.
type CommandUnsubscribe struct {
	Command
	GameID uuid.UUID `json:"game_id"`
}

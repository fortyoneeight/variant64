package models

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
)

// ErrInvalidCommand represents a invalid command.
type ErrInvalidCommand struct{}

func (e ErrInvalidCommand) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e ErrInvalidCommand) Error() string {
	return "Websocket action not defined"
}

// ErrFailedCommandMarshall represents a failed command marshall.
type ErrFailedCommandMarshall struct{}

func (e ErrFailedCommandMarshall) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e ErrFailedCommandMarshall) Error() string {
	return "Websocket action failed to marshall"
}

type ErrFailedUpdatePub struct {
	entityType string
}

func (e ErrFailedUpdatePub) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e ErrFailedUpdatePub) Error() string {
	return fmt.Sprintf("%s error: cannot create new update pub", e.entityType)
}

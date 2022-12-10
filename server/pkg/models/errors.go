package models

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
)

// ErrInvalidCommand represents a invalid command.
var ErrInvalidCommand = errortypes.New(errortypes.BadRequest, "Websocket action not defined")

// ErrFailedCommandMarshall represents a failed command marshall.
var ErrFailedCommandMarshall = errortypes.New(errortypes.BadRequest, "Websocket action failed to marshall")

// ErrFailedUpdatePub represents a failed command marshall.
var ErrFailedUpdatePub = func(entityType string) errortypes.TypedError {
	return errortypes.New(errortypes.InternalError, fmt.Sprintf("%s error: cannot create new update pub", entityType))
}

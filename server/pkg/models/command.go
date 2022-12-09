package models

import (
	"encoding/json"

	"github.com/variant64/server/pkg/errortypes"
)

// commander represent an command object that can perform an action.
type commander interface {
	PerformAction() errortypes.TypedError
}

// Command represents a command type.
type Command struct {
	Command string `json:"command"`
}

// HandleCommand performs the command action.
func HandleCommand(c commander, err errortypes.TypedError) errortypes.TypedError {
	if err != nil {
		return err
	}

	return c.PerformAction()
}

// MarshallCommand generically marshalls a command to a type.
func MarshallCommand[T any](command string, t *T) (*T, errortypes.TypedError) {
	err := json.Unmarshal([]byte(command), t)
	if err != nil {
		return nil, ErrFailedCommandMarshall{}
	}

	return t, nil
}

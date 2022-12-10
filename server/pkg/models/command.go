package models

import (
	"encoding/json"
)

// commander represent an command object that can perform an action.
type commander interface {
	PerformAction() error
}

// Command represents a command type.
type Command struct {
	Command string `json:"command"`
}

// HandleCommand performs the command action.
func HandleCommand(c commander, err error) error {
	if err != nil {
		return err
	}

	return c.PerformAction()
}

// MarshallCommand generically marshalls a command to a type.
func MarshallCommand[T any](command string, t *T) (*T, error) {
	err := json.Unmarshal([]byte(command), t)
	if err != nil {
		return nil, ErrFailedCommandMarshall
	}

	return t, nil
}

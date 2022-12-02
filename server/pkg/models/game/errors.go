package game

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
)

type errGameNotFound struct{}

func (e errGameNotFound) GetType() errortypes.Type {
	return errortypes.NotFound
}

func (e errGameNotFound) Error() string {
	return "Player error: not found"
}

type errInvalidPlayersNumber struct {
	number int
}

func (i errInvalidPlayersNumber) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (i errInvalidPlayersNumber) Error() string {
	return fmt.Sprintf("Player error: invalid number of players: %d", i.number)
}

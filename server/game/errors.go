package game

import (
	"fmt"

	"github.com/variant64/server/errortypes"
)

type errGameNotFound struct{}

func (e errGameNotFound) GetType() errortypes.Type {
	return errortypes.NotFound
}

func (e errGameNotFound) Error() string {
	return "Game error: not found"
}

type errInvalidPlayersNumber struct {
	number int
}

func (i errInvalidPlayersNumber) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (i errInvalidPlayersNumber) Error() string {
	return fmt.Sprintf("Game error: invalid number of players: %d", i.number)
}

type errPlayerNotInGame struct{}

func (e errPlayerNotInGame) GetType() errortypes.Type {
	return errortypes.NotFound
}

func (e errPlayerNotInGame) Error() string {
	return "Game error: player not found in game"
}

type errGameFinsished struct{}

func (e errGameFinsished) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e errGameFinsished) Error() string {
	return "Game error: game already finished"
}

type errGameStarted struct{}

func (e errGameStarted) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e errGameStarted) Error() string {
	return "Game error: game already started"
}

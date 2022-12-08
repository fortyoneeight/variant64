package game

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/variant64/server/pkg/errortypes"
)

var errGameNotFound = errortypes.New(errortypes.NotFound, "Game error: not found")

var errInvalidPlayersNumber = func(numberOfPlayers int) errortypes.TypedError {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Game error: invalid number of players: %d", numberOfPlayers))
}

var errPlayerNotInGame = errortypes.New(errortypes.NotFound, "Game error: player not found in game")

var errIncorrectGameState = func(required, current gameState) errortypes.TypedError {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Game error: game is %s not %s", current, required))
}

var errUnableToCreateBoard = errortypes.New(errortypes.BadRequest, "Game error: unable to create game board")

var errInvalidMove = func(error error) errortypes.TypedError {
	return errortypes.New(errortypes.BadRequest, errors.Wrap(error, "Game error: invalid move ").Error())
}

var errNotPlayersTurn = func(playerID string) errortypes.TypedError {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Game error: incorrect player, not their turn %s", playerID))
}

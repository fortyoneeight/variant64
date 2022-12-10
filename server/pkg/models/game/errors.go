package game

import (
	"fmt"

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

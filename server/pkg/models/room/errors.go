package room

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
)

var errRoomNotFound = errortypes.New(errortypes.NotFound, "Room error: not found")

var errPlayerLimit = errortypes.New(errortypes.BadRequest, fmt.Sprintf("Room error: room has reached player_limit %d.", PLAYER_LIMIT_DEFAULT))

var errMissingName = errortypes.New(errortypes.BadRequest, "Room error: room_name is required")

var errNameTooLong = errortypes.New(errortypes.BadRequest, fmt.Sprintf("Player error: room_name longer than limit: %d", NAME_MAX_LENGTH))

var errDuplicatePlayer = func(playerID string) errortypes.TypedError {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Room error: duplicate player %s", playerID))
}

var errPlayerNotInRoom = func(playerID string) errortypes.TypedError {
	return errortypes.New(errortypes.NotFound, fmt.Sprintf("Room error: player not in room %s", playerID))
}

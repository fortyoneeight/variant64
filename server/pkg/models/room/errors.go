package room

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
)

var errRoomNotFound = errortypes.New(errortypes.NotFound, "Room error: not found")
var errMissingName = errortypes.New(errortypes.BadRequest, "Room error: room_name is required")
var errDuplicatePlayer = func(playerID string) errortypes.TypedError {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Room error: duplicate player %s", playerID))
}

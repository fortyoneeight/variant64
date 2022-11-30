package room

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/variant64/server/errortypes"
)

type errRoomNotFound struct{}

func (e errRoomNotFound) GetType() errortypes.Type {
	return errortypes.NotFound
}

func (e errRoomNotFound) Error() string {
	return "Room error: not found"
}

type errMissingName struct{}

func (e errMissingName) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e errMissingName) Error() string {
	return "Room error: room_name is required"
}

type errDuplicatePlayer struct {
	playerID uuid.UUID
}

func (e errDuplicatePlayer) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e errDuplicatePlayer) Error() string {
	return fmt.Sprintf("Room error: duplicate player %s", e.playerID)
}

package player

import (
	"github.com/variant64/server/errortypes"
)

type errPlayerNotFound struct{}

func (e errPlayerNotFound) GetType() errortypes.Type {
	return errortypes.NotFound
}

func (e errPlayerNotFound) Error() string {
	return "Player error: not found"
}

type missingDisplayName struct {
	err error
}

func (i missingDisplayName) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (i missingDisplayName) Error() string {
	return "Player error: missing display_name"
}

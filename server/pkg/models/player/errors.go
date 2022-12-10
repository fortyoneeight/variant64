package player

import (
	"github.com/variant64/server/pkg/errortypes"
)

var errPlayerNotFound = errortypes.New(errortypes.NotFound, "Player error: not found")
var errMissingDisplayName = errortypes.New(errortypes.BadRequest, "Player error: missing display_name")

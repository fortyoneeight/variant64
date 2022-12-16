package player

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
)

var errPlayerNotFound = errortypes.New(errortypes.NotFound, "Player error: not found")

var errMissingDisplayName = errortypes.New(errortypes.BadRequest, "Player error: missing display_name")

var errDisplayNameTooLong = errortypes.New(errortypes.BadRequest, fmt.Sprintf("Player error: display_name longer than limit: %d", DISPLAY_NAME_MAX_LENGTH))

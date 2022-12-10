package classic

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
	"github.com/variant64/server/pkg/models/board"
)

var errPieceNotFound = errortypes.New(errortypes.NotFound, "Board error: piece not found")
var errApplyingMove = errortypes.New(errortypes.BadRequest, "Board error: unable to apply move")
var errMoveNotAllowed = errortypes.New(errortypes.BadRequest, "Board error: move is not allowed")
var errInvalidColor = func(color board.Color) errortypes.TypedError {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Board error: color is invalid %s", color.String()))
}
var errNotAllowedToCastle = func(color board.Color) errortypes.TypedError {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Board error: player %s is not allowed to castle", color.String()))
}

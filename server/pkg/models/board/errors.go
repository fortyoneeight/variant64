package board

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
)

var errCannotHandleMoveType = func(m MoveType) error {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Move error: cannot handle move type %s", m))
}

var errSourcePieceNotFound = errortypes.New(errortypes.BadRequest, "Move error: source piece not found.")

var errInvalidColor = func(color Color) error {
	return errortypes.New(errortypes.BadRequest, fmt.Sprintf("Move error: color is invalid %s", color.String()))
}

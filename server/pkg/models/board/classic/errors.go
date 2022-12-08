package classic

import (
	"fmt"

	"github.com/variant64/server/pkg/errortypes"
	"github.com/variant64/server/pkg/models/board"
)

type errPieceNotFound struct{}

func (e errPieceNotFound) GetType() errortypes.Type {
	return errortypes.NotFound
}

func (e errPieceNotFound) Error() string {
	return "Board error: piece not found"
}

type errApplyingMove struct{}

func (e errApplyingMove) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e errApplyingMove) Error() string {
	return "Board error: unable to apply move"
}

type errMoveNotAllowed struct{}

func (e errMoveNotAllowed) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e errMoveNotAllowed) Error() string {
	return "Board error: move is not allowed"
}

type errInvalidColor struct {
	color board.Color
}

func (e errInvalidColor) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e errInvalidColor) Error() string {
	return fmt.Sprintf("Board error: color is invalid %s", e.color.String())
}

type errNotAllowedToCastle struct {
	color board.Color
}

func (e errNotAllowedToCastle) GetType() errortypes.Type {
	return errortypes.BadRequest
}

func (e errNotAllowedToCastle) Error() string {
	return fmt.Sprintf("Board error: player %s is not allowed to castle", e.color.String())
}

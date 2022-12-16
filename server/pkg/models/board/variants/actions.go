package variants

import "github.com/variant64/server/pkg/models/board"

type RequestNewClassicBoard struct{}

func (r *RequestNewClassicBoard) PerformAction() (*board.Board, error) {
	return NewClassicBoard(), nil
}

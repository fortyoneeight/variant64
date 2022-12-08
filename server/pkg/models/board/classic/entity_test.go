package classic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/pkg/models/board"
)

func TestCheckPawnMoves(t *testing.T) {
	bounds := board.Bounds{Rank: 8, File: 8}
	classicBoard := &ClassicBoard{
		Bounds: bounds,
		locations: map[int]map[int]piece{
			1: {0: &board.Pawn{Color: board.WHITE}},
			6: {0: &board.Pawn{Color: board.BLACK}},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "White pawn start.",
			position: board.Position{Rank: 1, File: 0},
			piece:    &board.Pawn{Color: board.WHITE},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 2, File: 0},
				},
				board.PAWN_DOUBLE_PUSH: []board.Position{
					{Rank: 3, File: 0},
				},
				board.CAPTURE:          []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
		{
			name:     "Black pawn start.",
			position: board.Position{Rank: 6, File: 0},
			piece:    &board.Pawn{Color: board.BLACK},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 5, File: 0},
				},
				board.PAWN_DOUBLE_PUSH: []board.Position{
					{Rank: 4, File: 0},
				},
				board.CAPTURE:          []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			legalMoves := classicBoard.getLegalMoves(test.position)
			assert.Equal(t, test.expectedLegalMoves, legalMoves)
		})
	}
}

func TestCheckKnightMoves(t *testing.T) {
	classicBoard := &ClassicBoard{
		Bounds: board.Bounds{Rank: 8, File: 8},
		locations: map[int]map[int]piece{
			3: {3: &board.Knight{Color: board.WHITE}},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "Knight in the middle of the board.",
			position: board.Position{Rank: 3, File: 3},
			piece:    &board.Knight{Color: board.WHITE},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 5, File: 4},
					{Rank: 5, File: 2},
					{Rank: 4, File: 5},
					{Rank: 4, File: 1},
					{Rank: 2, File: 5},
					{Rank: 2, File: 1},
					{Rank: 1, File: 4},
					{Rank: 1, File: 2},
				},
				board.CAPTURE:          []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			legalMoves := classicBoard.getLegalMoves(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
		})
	}
}

func TestCheckQueenMoves(t *testing.T) {
	bounds := board.Bounds{Rank: 8, File: 8}
	classicBoard := &ClassicBoard{
		Bounds: bounds,
		locations: map[int]map[int]piece{
			3: {3: &board.Queen{Bounds: bounds, Color: board.WHITE}},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "Queen in the middle of the board.",
			position: board.Position{Rank: 3, File: 3},
			piece:    &board.Queen{Color: board.WHITE},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 0, File: 3},
					{Rank: 1, File: 3},
					{Rank: 2, File: 3},
					{Rank: 4, File: 3},
					{Rank: 5, File: 3},
					{Rank: 6, File: 3},
					{Rank: 7, File: 3},
					{Rank: 3, File: 0},
					{Rank: 3, File: 1},
					{Rank: 3, File: 2},
					{Rank: 3, File: 4},
					{Rank: 3, File: 5},
					{Rank: 3, File: 6},
					{Rank: 3, File: 7},
					{Rank: 0, File: 0},
					{Rank: 1, File: 1},
					{Rank: 2, File: 2},
					{Rank: 4, File: 4},
					{Rank: 5, File: 5},
					{Rank: 6, File: 6},
					{Rank: 7, File: 7},
					{Rank: 6, File: 0},
					{Rank: 5, File: 1},
					{Rank: 4, File: 2},
					{Rank: 2, File: 4},
					{Rank: 1, File: 5},
					{Rank: 0, File: 6},
				},
				board.CAPTURE:          []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			legalMoves := classicBoard.getLegalMoves(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
		})
	}
}

func TestCheckKingMoves(t *testing.T) {
	bounds := board.Bounds{Rank: 8, File: 8}
	classicBoard := &ClassicBoard{
		Bounds:                      bounds,
		whiteAllowedKingsideCastle:  true,
		whiteAllowedQueensideCastle: true,
		blackAllowedKingsideCastle:  true,
		blackAllowedQueensideCastle: true,
		locations: map[int]map[int]piece{
			0: {
				0: &board.Rook{Bounds: bounds, Color: board.WHITE},
				4: &board.King{Color: board.WHITE},
				7: &board.Rook{Bounds: bounds, Color: board.WHITE},
			},
			4: {
				4: &board.King{Color: board.WHITE},
			},
			7: {
				0: &board.Rook{Bounds: bounds, Color: board.BLACK},
				4: &board.King{Color: board.BLACK},
				7: &board.Rook{Bounds: bounds, Color: board.BLACK},
			},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "King in the middle of the board.",
			position: board.Position{Rank: 4, File: 4},
			piece:    &board.King{Color: board.WHITE},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 5, File: 5},
					{Rank: 5, File: 4},
					{Rank: 5, File: 3},
					{Rank: 4, File: 5},
					{Rank: 4, File: 3},
					{Rank: 3, File: 5},
					{Rank: 3, File: 4},
					{Rank: 3, File: 3},
				},
				board.CAPTURE:          []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
		{
			name:     "White king in the original position.",
			position: board.Position{Rank: 0, File: 4},
			piece:    &board.King{Color: board.WHITE},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 0, File: 3},
					{Rank: 0, File: 5},
					{Rank: 1, File: 3},
					{Rank: 1, File: 4},
					{Rank: 1, File: 5},
				},
				board.CAPTURE: []board.Position{},
				board.KINGSIDE_CASTLE: []board.Position{
					{Rank: 0, File: 6},
				},
				board.QUEENSIDE_CASTLE: []board.Position{
					{Rank: 0, File: 2},
				},
			},
		},
		{
			name:     "Black king in the original position.",
			position: board.Position{Rank: 7, File: 4},
			piece:    &board.King{Color: board.BLACK},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 7, File: 3},
					{Rank: 7, File: 5},
					{Rank: 6, File: 3},
					{Rank: 6, File: 4},
					{Rank: 6, File: 5},
				},
				board.CAPTURE: []board.Position{},
				board.KINGSIDE_CASTLE: []board.Position{
					{Rank: 7, File: 6},
				},
				board.QUEENSIDE_CASTLE: []board.Position{
					{Rank: 7, File: 2},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			legalMoves := classicBoard.getLegalMoves(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
			for _, move := range tc.expectedLegalMoves[board.CAPTURE] {
				assert.Contains(t, legalMoves[board.CAPTURE], move)
			}
			for _, move := range tc.expectedLegalMoves[board.KINGSIDE_CASTLE] {
				assert.Contains(t, legalMoves[board.KINGSIDE_CASTLE], move)
			}
			for _, move := range tc.expectedLegalMoves[board.QUEENSIDE_CASTLE] {
				assert.Contains(t, legalMoves[board.QUEENSIDE_CASTLE], move)
			}
		})
	}
}

func TestCheckRookMoves(t *testing.T) {
	bounds := board.Bounds{Rank: 8, File: 8}
	classicBoard := &ClassicBoard{
		Bounds: bounds,
		locations: map[int]map[int]piece{
			3: {3: &board.Rook{Bounds: bounds, Color: board.WHITE}},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "Rook in the middle of the board.",
			position: board.Position{Rank: 3, File: 3},
			piece:    &board.Rook{Bounds: bounds, Color: board.WHITE},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 0, File: 3},
					{Rank: 1, File: 3},
					{Rank: 2, File: 3},
					{Rank: 4, File: 3},
					{Rank: 5, File: 3},
					{Rank: 6, File: 3},
					{Rank: 7, File: 3},
					{Rank: 3, File: 0},
					{Rank: 3, File: 1},
					{Rank: 3, File: 2},
					{Rank: 3, File: 4},
					{Rank: 3, File: 5},
					{Rank: 3, File: 6},
					{Rank: 3, File: 7},
				},
				board.CAPTURE:          []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			legalMoves := classicBoard.getLegalMoves(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
		})
	}
}

func TestCheckBishopMoves(t *testing.T) {
	bounds := board.Bounds{Rank: 8, File: 8}
	classicBoard := &ClassicBoard{
		Bounds: board.Bounds{Rank: 8, File: 8},
		locations: map[int]map[int]piece{
			3: {3: &board.Bishop{Bounds: bounds, Color: board.WHITE}},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "Bishop in the middle of the board.",
			position: board.Position{Rank: 3, File: 3},
			piece:    &board.Bishop{Bounds: bounds, Color: board.WHITE},
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 0, File: 0},
					{Rank: 1, File: 1},
					{Rank: 2, File: 2},
					{Rank: 4, File: 4},
					{Rank: 5, File: 5},
					{Rank: 6, File: 6},
					{Rank: 7, File: 7},
					{Rank: 6, File: 0},
					{Rank: 5, File: 1},
					{Rank: 4, File: 2},
					{Rank: 2, File: 4},
					{Rank: 1, File: 5},
					{Rank: 0, File: 6},
				},
				board.CAPTURE:          []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			legalMoves := classicBoard.getLegalMoves(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
		})
	}
}

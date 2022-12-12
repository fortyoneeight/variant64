package classic

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/pkg/models/board"
)

func TestCheckPawnMoves(t *testing.T) {
	bounds := board.Bounds{RankCount: 8, FileCount: 8}
	classicBoard := &ClassicBoard{
		Bounds: bounds,
		gameboardState: board.GameboardState{
			1: {0: board.NewPawn(board.WHITE)},
			6: {0: board.NewPawn(board.BLACK)},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              *board.Piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "White pawn start.",
			position: board.Position{Rank: 1, File: 0},
			piece:    board.NewPawn(board.WHITE),
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 2, File: 0},
				},
				board.PAWN_DOUBLE_PUSH: []board.Position{
					{Rank: 3, File: 0},
				},
				board.CAPTURE:          []board.Position{},
				board.RAY:              []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
		{
			name:     "Black pawn start.",
			position: board.Position{Rank: 6, File: 0},
			piece:    board.NewPawn(board.BLACK),
			expectedLegalMoves: board.MoveMap{
				board.NORMAL: []board.Position{
					{Rank: 5, File: 0},
				},
				board.PAWN_DOUBLE_PUSH: []board.Position{
					{Rank: 4, File: 0},
				},
				board.CAPTURE:          []board.Position{},
				board.RAY:              []board.Position{},
				board.KINGSIDE_CASTLE:  []board.Position{},
				board.QUEENSIDE_CASTLE: []board.Position{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			legalMoves := classicBoard.getMovesAtPosition(test.position)
			assert.Equal(t, test.expectedLegalMoves, legalMoves)
		})
	}
}

func TestCheckKnightMoves(t *testing.T) {
	classicBoard := &ClassicBoard{
		Bounds: board.Bounds{RankCount: 8, FileCount: 8},
		gameboardState: board.GameboardState{
			3: {3: board.NewKnight(board.WHITE)},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              *board.Piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "Knight in the middle of the board.",
			position: board.Position{Rank: 3, File: 3},
			piece:    board.NewKnight(board.WHITE),
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
			legalMoves := classicBoard.getMovesAtPosition(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
		})
	}
}

func TestCheckQueenMoves(t *testing.T) {
	bounds := board.Bounds{RankCount: 8, FileCount: 8}
	classicBoard := &ClassicBoard{
		Bounds: bounds,
		gameboardState: board.GameboardState{
			3: {3: board.NewQueen(board.WHITE, bounds)},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              *board.Piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "Queen in the middle of the board.",
			position: board.Position{Rank: 3, File: 3},
			piece:    board.NewQueen(board.WHITE, bounds),
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
			legalMoves := classicBoard.getMovesAtPosition(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
		})
	}
}

func TestCheckKingMoves(t *testing.T) {
	bounds := board.Bounds{RankCount: 8, FileCount: 8}
	classicBoard := &ClassicBoard{
		Bounds:                      bounds,
		whiteAllowedKingsideCastle:  true,
		whiteAllowedQueensideCastle: true,
		blackAllowedKingsideCastle:  true,
		blackAllowedQueensideCastle: true,
		gameboardState: board.GameboardState{
			0: {
				0: board.NewRook(board.WHITE, bounds),
				4: board.NewKing(board.WHITE),
				7: board.NewRook(board.WHITE, bounds),
			},
			4: {
				4: board.NewKing(board.WHITE),
			},
			7: {
				0: board.NewRook(board.BLACK, bounds),
				4: board.NewKing(board.BLACK),
				7: board.NewRook(board.BLACK, bounds),
			},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              *board.Piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "King in the middle of the board.",
			position: board.Position{Rank: 4, File: 4},
			piece:    board.NewKing(board.WHITE),
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
			piece:    board.NewKing(board.WHITE),
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
			piece:    board.NewKing(board.BLACK),
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
			legalMoves := classicBoard.getMovesAtPosition(tc.position)
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
	bounds := board.Bounds{RankCount: 8, FileCount: 8}
	classicBoard := &ClassicBoard{
		Bounds: bounds,
		gameboardState: board.GameboardState{
			3: {3: board.NewRook(board.WHITE, bounds)},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              *board.Piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "Rook in the middle of the board.",
			position: board.Position{Rank: 3, File: 3},
			piece:    board.NewRook(board.WHITE, bounds),
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
			legalMoves := classicBoard.getMovesAtPosition(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
		})
	}
}

func TestCheckBishopMoves(t *testing.T) {
	bounds := board.Bounds{RankCount: 8, FileCount: 8}
	classicBoard := &ClassicBoard{
		Bounds: board.Bounds{RankCount: 8, FileCount: 8},
		gameboardState: board.GameboardState{
			3: {3: board.NewBishop(board.WHITE, bounds)},
		},
	}

	tests := []struct {
		name               string
		position           board.Position
		piece              *board.Piece
		expectedLegalMoves board.MoveMap
	}{
		{
			name:     "Bishop in the middle of the board.",
			position: board.Position{Rank: 3, File: 3},
			piece:    board.NewBishop(board.WHITE, bounds),
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
			legalMoves := classicBoard.getMovesAtPosition(tc.position)
			for _, move := range tc.expectedLegalMoves[board.NORMAL] {
				assert.Contains(t, legalMoves[board.NORMAL], move)
			}
		})
	}
}

func TestHandleMove(t *testing.T) {
	bounds := board.Bounds{RankCount: 8, FileCount: 8}
	testCases := []struct {
		name          string
		moves         []board.Move
		initialBoard  ClassicBoard
		expectedBoard ClassicBoard
		expectedErr   error
	}{
		{
			name: "pawn normal move",
			moves: []board.Move{
				{
					Source:      board.Position{Rank: 0, File: 0},
					Destination: board.Position{Rank: 1, File: 0},
					MoveType:    board.NORMAL,
				},
				{
					Source:      board.Position{Rank: 6, File: 0},
					Destination: board.Position{Rank: 5, File: 0},
					MoveType:    board.NORMAL,
				},
			},
			initialBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  true,
				whiteAllowedQueensideCastle: true,
				blackAllowedKingsideCastle:  true,
				blackAllowedQueensideCastle: true,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						0: {
							0: board.NewPawn(board.WHITE),
						},
						6: {
							0: board.NewPawn(board.BLACK),
						},
					},
				),
			},
			expectedBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  true,
				whiteAllowedQueensideCastle: true,
				blackAllowedKingsideCastle:  true,
				blackAllowedQueensideCastle: true,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						1: {
							0: board.NewPawn(board.WHITE),
						},
						5: {
							0: board.NewPawn(board.BLACK),
						},
					},
				),
			},
			expectedErr: nil,
		},
		{
			name: "pawn double push",
			moves: []board.Move{
				{
					Source:      board.Position{Rank: 1, File: 0},
					Destination: board.Position{Rank: 3, File: 0},
					MoveType:    board.PAWN_DOUBLE_PUSH,
				},
				{
					Source:      board.Position{Rank: 6, File: 0},
					Destination: board.Position{Rank: 4, File: 0},
					MoveType:    board.PAWN_DOUBLE_PUSH,
				},
			},
			initialBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  true,
				whiteAllowedQueensideCastle: true,
				blackAllowedKingsideCastle:  true,
				blackAllowedQueensideCastle: true,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						1: {
							0: board.NewPawn(board.WHITE),
						},
						6: {
							0: board.NewPawn(board.BLACK),
						},
					},
				),
			},
			expectedBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  true,
				whiteAllowedQueensideCastle: true,
				blackAllowedKingsideCastle:  true,
				blackAllowedQueensideCastle: true,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						3: {
							0: board.NewPawn(board.WHITE),
						},
						4: {
							0: board.NewPawn(board.BLACK),
						},
					},
				),
			},
			expectedErr: nil,
		},
		{
			name: "kingside castle",
			moves: []board.Move{
				{
					Source:      board.Position{Rank: 0, File: 4},
					Destination: board.Position{Rank: 0, File: 6},
					MoveType:    board.KINGSIDE_CASTLE,
				},
				{
					Source:      board.Position{Rank: 7, File: 4},
					Destination: board.Position{Rank: 7, File: 6},
					MoveType:    board.KINGSIDE_CASTLE,
				},
			},
			initialBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  true,
				whiteAllowedQueensideCastle: true,
				blackAllowedKingsideCastle:  true,
				blackAllowedQueensideCastle: true,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						0: {
							4: board.NewKing(board.WHITE),
							7: board.NewRook(board.WHITE, bounds),
						},
						7: {
							4: board.NewKing(board.BLACK),
							7: board.NewRook(board.BLACK, bounds),
						},
					},
				),
			},
			expectedBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  false,
				whiteAllowedQueensideCastle: false,
				blackAllowedKingsideCastle:  false,
				blackAllowedQueensideCastle: false,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						0: {
							6: board.NewKing(board.WHITE),
							5: board.NewRook(board.WHITE, bounds),
						},
						7: {
							6: board.NewKing(board.BLACK),
							5: board.NewRook(board.BLACK, bounds),
						},
					},
				),
			},
			expectedErr: nil,
		},
		{
			name: "queenside castle",
			moves: []board.Move{
				{
					Source:      board.Position{Rank: 0, File: 4},
					Destination: board.Position{Rank: 0, File: 2},
					MoveType:    board.QUEENSIDE_CASTLE,
				},
				{
					Source:      board.Position{Rank: 7, File: 4},
					Destination: board.Position{Rank: 7, File: 2},
					MoveType:    board.QUEENSIDE_CASTLE,
				},
			},
			initialBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  true,
				whiteAllowedQueensideCastle: true,
				blackAllowedKingsideCastle:  true,
				blackAllowedQueensideCastle: true,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						0: {
							0: board.NewRook(board.WHITE, bounds),
							4: board.NewKing(board.WHITE),
						},
						7: {
							0: board.NewRook(board.BLACK, bounds),
							4: board.NewKing(board.BLACK),
						},
					},
				),
			},
			expectedBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  false,
				whiteAllowedQueensideCastle: false,
				blackAllowedKingsideCastle:  false,
				blackAllowedQueensideCastle: false,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						0: {
							2: board.NewKing(board.WHITE),
							3: board.NewRook(board.WHITE, bounds),
						},
						7: {
							2: board.NewKing(board.BLACK),
							3: board.NewRook(board.BLACK, bounds),
						},
					},
				),
			},
			expectedErr: nil,
		},
		{
			name: "pawn capture",
			moves: []board.Move{
				{
					Source:      board.Position{Rank: 1, File: 0},
					Destination: board.Position{Rank: 2, File: 1},
					MoveType:    board.CAPTURE,
				},
				{
					Source:      board.Position{Rank: 7, File: 0},
					Destination: board.Position{Rank: 6, File: 1},
					MoveType:    board.CAPTURE,
				},
			},
			initialBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  true,
				whiteAllowedQueensideCastle: true,
				blackAllowedKingsideCastle:  true,
				blackAllowedQueensideCastle: true,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						1: {
							0: board.NewPawn(board.WHITE),
						},
						2: {
							1: board.NewPawn(board.BLACK),
						},
						6: {
							1: board.NewPawn(board.WHITE),
						},
						7: {
							0: board.NewPawn(board.BLACK),
						},
					},
				),
			},
			expectedBoard: ClassicBoard{
				Bounds:                      bounds,
				whiteAllowedKingsideCastle:  true,
				whiteAllowedQueensideCastle: true,
				blackAllowedKingsideCastle:  true,
				blackAllowedQueensideCastle: true,
				gameboardState: NewGameboardState(
					bounds,
					board.GameboardState{
						2: {
							1: board.NewPawn(board.WHITE),
						},
						6: {
							1: board.NewPawn(board.BLACK),
						},
					},
				),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Create a new board for each test case to avoid modifying the initial board.
		tc.initialBoard.updateAvailableMoves()
		tc.expectedBoard.updateAvailableMoves()
		board := tc.initialBoard

		// Call the HandleMove method on the board.
		for _, move := range tc.moves {
			err := board.HandleMove(move)
			if err != nil {
				t.Errorf("Test case %s: expected no error but got %v", tc.name, err)
			}
		}

		// Check that the board is in the expected state.
		if !reflect.DeepEqual(board, tc.expectedBoard) {
			t.Errorf("Test case %s: expected board %v but got %v", tc.name, tc.expectedBoard, board)
		}
	}
}

package board

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPawnMoves(t *testing.T) {
	classicBoard := Build(
		WithGameboardState(
			GameboardState{
				1: {0: NewPawn(WHITE)},
				6: {0: NewPawn(BLACK)},
			},
		),
	)

	tests := []struct {
		name               string
		position           Position
		piece              *Piece
		expectedLegalMoves MoveMap
	}{
		{
			name:     "White pawn start.",
			position: Position{Rank: 1, File: 0},
			piece:    NewPawn(WHITE),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{
					{Rank: 2, File: 0},
				},
				PAWN_DOUBLE_PUSH: []Position{
					{Rank: 3, File: 0},
				},
				CAPTURE:          []Position{},
				JUMP:             []Position{},
				JUMP_CAPTURE:     []Position{},
				KINGSIDE_CASTLE:  []Position{},
				QUEENSIDE_CASTLE: []Position{},
			},
		},
		{
			name:     "Black pawn start.",
			position: Position{Rank: 6, File: 0},
			piece:    NewPawn(BLACK),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{
					{Rank: 5, File: 0},
				},
				PAWN_DOUBLE_PUSH: []Position{
					{Rank: 4, File: 0},
				},
				CAPTURE:          []Position{},
				JUMP:             []Position{},
				JUMP_CAPTURE:     []Position{},
				KINGSIDE_CASTLE:  []Position{},
				QUEENSIDE_CASTLE: []Position{},
			},
		},
	}

	for _, tc := range tests {
		availableMoves := classicBoard.GameboardState[tc.position.Rank][tc.position.File].AvailableMoves
		for moveType, expectedMovesByType := range tc.expectedLegalMoves {
			for _, move := range expectedMovesByType {
				assert.Contains(t, availableMoves[moveType], move)
			}
		}
	}
}

func TestKnightMoves(t *testing.T) {
	classicBoard := Build(
		WithGameboardState(
			GameboardState{
				3: {3: NewKnight(WHITE)},
			},
		),
	)

	tests := []struct {
		name               string
		position           Position
		piece              *Piece
		expectedLegalMoves MoveMap
	}{
		{
			name:     "Knight in the middle of the ",
			position: Position{Rank: 3, File: 3},
			piece:    NewKnight(WHITE),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{},
				JUMP: []Position{
					{Rank: 5, File: 4},
					{Rank: 5, File: 2},
					{Rank: 4, File: 5},
					{Rank: 4, File: 1},
					{Rank: 2, File: 5},
					{Rank: 2, File: 1},
					{Rank: 1, File: 4},
					{Rank: 1, File: 2},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			availableMoves := classicBoard.GameboardState[tc.position.Rank][tc.position.File].AvailableMoves
			for moveType, expectedMovesByType := range tc.expectedLegalMoves {
				for _, move := range expectedMovesByType {
					assert.Contains(t, availableMoves[moveType], move)
				}
			}
		})
	}
}

func TestQueenMoves(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	classicBoard := Build(
		WithBounds(bounds),
		WithGameboardState(
			GameboardState{
				3: {3: NewQueen(WHITE, bounds)},
			},
		),
	)

	tests := []struct {
		name               string
		position           Position
		piece              *Piece
		expectedLegalMoves MoveMap
	}{
		{
			name:     "Queen in the middle of the ",
			position: Position{Rank: 3, File: 3},
			piece:    NewQueen(WHITE, bounds),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{
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
				JUMP:             []Position{},
				JUMP_CAPTURE:     []Position{},
				CAPTURE:          []Position{},
				PAWN_DOUBLE_PUSH: []Position{},
				KINGSIDE_CASTLE:  []Position{},
				QUEENSIDE_CASTLE: []Position{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			availableMoves := classicBoard.GameboardState[tc.position.Rank][tc.position.File].AvailableMoves
			for moveType, expectedMovesByType := range tc.expectedLegalMoves {
				for _, move := range expectedMovesByType {
					assert.Contains(t, availableMoves[moveType], move)
				}
			}
		})
	}
}

func TestKingMoves(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	classicBoard := Build(
		WithBounds(bounds),
		WithGameboardState(
			GameboardState{
				0: {
					0: NewRook(WHITE, bounds),
					4: NewKing(WHITE),
					7: NewRook(WHITE, bounds),
				},
				4: {
					4: NewKing(WHITE),
				},
				7: {
					0: NewRook(BLACK, bounds),
					4: NewKing(BLACK),
					7: NewRook(BLACK, bounds),
				},
			},
		),
	)

	tests := []struct {
		name               string
		position           Position
		piece              *Piece
		expectedLegalMoves MoveMap
	}{
		{
			name:     "King in the middle of the ",
			position: Position{Rank: 4, File: 4},
			piece:    NewKing(WHITE),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{
					{Rank: 5, File: 5},
					{Rank: 5, File: 4},
					{Rank: 5, File: 3},
					{Rank: 4, File: 5},
					{Rank: 4, File: 3},
					{Rank: 3, File: 5},
					{Rank: 3, File: 4},
					{Rank: 3, File: 3},
				},
				CAPTURE:          []Position{},
				KINGSIDE_CASTLE:  []Position{},
				QUEENSIDE_CASTLE: []Position{},
			},
		},
		{
			name:     "White king in the original position.",
			position: Position{Rank: 0, File: 4},
			piece:    NewKing(WHITE),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{
					{Rank: 0, File: 3},
					{Rank: 0, File: 5},
					{Rank: 1, File: 3},
					{Rank: 1, File: 4},
					{Rank: 1, File: 5},
				},
				CAPTURE: []Position{},
				KINGSIDE_CASTLE: []Position{
					{Rank: 0, File: 6},
				},
				QUEENSIDE_CASTLE: []Position{
					{Rank: 0, File: 2},
				},
			},
		},
		{
			name:     "Black king in the original position.",
			position: Position{Rank: 7, File: 4},
			piece:    NewKing(BLACK),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{
					{Rank: 7, File: 3},
					{Rank: 7, File: 5},
					{Rank: 6, File: 3},
					{Rank: 6, File: 4},
					{Rank: 6, File: 5},
				},
				CAPTURE: []Position{},
				KINGSIDE_CASTLE: []Position{
					{Rank: 7, File: 6},
				},
				QUEENSIDE_CASTLE: []Position{
					{Rank: 7, File: 2},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			availableMoves := classicBoard.GameboardState[tc.position.Rank][tc.position.File].AvailableMoves
			for moveType, expectedMovesByType := range tc.expectedLegalMoves {
				for _, move := range expectedMovesByType {
					assert.Contains(t, availableMoves[moveType], move)
				}
			}
		})
	}
}

func TestRookMoves(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	classicBoard := Build(
		WithBounds(bounds),
		WithGameboardState(
			GameboardState{
				3: {3: NewRook(WHITE, bounds)},
			},
		),
	)

	tests := []struct {
		name               string
		position           Position
		piece              *Piece
		expectedLegalMoves MoveMap
	}{
		{
			name:     "Rook in the middle of the ",
			position: Position{Rank: 3, File: 3},
			piece:    NewRook(WHITE, bounds),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{
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
				CAPTURE:          []Position{},
				KINGSIDE_CASTLE:  []Position{},
				QUEENSIDE_CASTLE: []Position{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			availableMoves := classicBoard.GameboardState[tc.position.Rank][tc.position.File].AvailableMoves
			for moveType, expectedMovesByType := range tc.expectedLegalMoves {
				for _, move := range expectedMovesByType {
					assert.Contains(t, availableMoves[moveType], move)
				}
			}
		})
	}
}

func TestBishopMoves(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	classicBoard := Build(
		WithBounds(bounds),
		WithGameboardState(
			GameboardState{
				3: {3: NewBishop(WHITE, bounds)},
			},
		),
	)

	tests := []struct {
		name               string
		position           Position
		piece              *Piece
		expectedLegalMoves MoveMap
	}{
		{
			name:     "Bishop in the middle of the ",
			position: Position{Rank: 3, File: 3},
			piece:    NewBishop(WHITE, bounds),
			expectedLegalMoves: MoveMap{
				NORMAL: []Position{
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
				CAPTURE:          []Position{},
				KINGSIDE_CASTLE:  []Position{},
				QUEENSIDE_CASTLE: []Position{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			availableMoves := classicBoard.GameboardState[tc.position.Rank][tc.position.File].AvailableMoves
			for moveType, expectedMovesByType := range tc.expectedLegalMoves {
				for _, move := range expectedMovesByType {
					assert.Contains(t, availableMoves[moveType], move)
				}
			}
		})
	}
}

func TestHandleMove(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	testCases := []struct {
		name          string
		moves         []Move
		initialBoard  *Board
		expectedBoard *Board
		expectedErr   error
	}{
		{
			name: "pawn normal move",
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 1, File: 0},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 6, File: 0},
					Destination: Position{Rank: 5, File: 0},
					MoveType:    NORMAL,
				},
			},
			initialBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						0: {
							0: NewPawn(WHITE),
						},
						6: {
							0: NewPawn(BLACK),
						},
					},
				),
			),
			expectedBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						1: {
							0: NewPawn(WHITE),
						},
						5: {
							0: NewPawn(BLACK),
						},
					},
				),
			),
			expectedErr: nil,
		},
		{
			name: "pawn double push",
			moves: []Move{
				{
					Source:      Position{Rank: 1, File: 0},
					Destination: Position{Rank: 3, File: 0},
					MoveType:    PAWN_DOUBLE_PUSH,
				},
				{
					Source:      Position{Rank: 6, File: 0},
					Destination: Position{Rank: 4, File: 0},
					MoveType:    PAWN_DOUBLE_PUSH,
				},
			},
			initialBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						1: {
							0: NewPawn(WHITE),
						},
						6: {
							0: NewPawn(BLACK),
						},
					},
				),
			),
			expectedBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						3: {
							0: NewPawn(WHITE),
						},
						4: {
							0: NewPawn(BLACK),
						},
					},
				),
			),
			expectedErr: nil,
		},
		{
			name: "kingside castle",
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 0, File: 6},
					MoveType:    KINGSIDE_CASTLE,
				},
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 7, File: 6},
					MoveType:    KINGSIDE_CASTLE,
				},
			},
			initialBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						0: {
							4: NewKing(WHITE),
							7: NewRook(WHITE, bounds),
						},
						7: {
							4: NewKing(BLACK),
							7: NewRook(BLACK, bounds),
						},
					},
				),
			),
			expectedBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						0: {
							6: NewKing(WHITE),
							5: NewRook(WHITE, bounds),
						},
						7: {
							6: NewKing(BLACK),
							5: NewRook(BLACK, bounds),
						},
					},
				),
				WithCastlingState(
					createCastlingState(false, false, false, false),
				),
				WithMoveFilter(
					NewMoveFilter(
						&FilterOutOfBounds{Bounds: bounds},
						&FilterPieceCollision{},
						&FilterFriendlyCapture{},
						&FilterInvalidPawnDoublePush{},
						&FilterIllegalKingsideCastle{
							createCastlingState(false, false, false, false),
						},
						&FilterIllegalQueensideCastle{
							createCastlingState(false, false, false, false),
						},
					),
				),
			),
			expectedErr: nil,
		},
		{
			name: "queenside castle",
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    QUEENSIDE_CASTLE,
				},
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 7, File: 2},
					MoveType:    QUEENSIDE_CASTLE,
				},
			},
			initialBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						0: {
							0: NewRook(WHITE, bounds),
							4: NewKing(WHITE),
						},
						7: {
							0: NewRook(BLACK, bounds),
							4: NewKing(BLACK),
						},
					},
				),
			),
			expectedBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						0: {
							2: NewKing(WHITE),
							3: NewRook(WHITE, bounds),
						},
						7: {
							2: NewKing(BLACK),
							3: NewRook(BLACK, bounds),
						},
					},
				),
				WithCastlingState(
					createCastlingState(false, false, false, false),
				),
				WithMoveFilter(
					NewMoveFilter(
						&FilterOutOfBounds{Bounds: bounds},
						&FilterPieceCollision{},
						&FilterFriendlyCapture{},
						&FilterInvalidPawnDoublePush{},
						&FilterIllegalKingsideCastle{
							createCastlingState(false, false, false, false),
						},
						&FilterIllegalQueensideCastle{
							createCastlingState(false, false, false, false),
						},
					),
				),
			),
			expectedErr: nil,
		},
		{
			name: "pawn capture",
			moves: []Move{
				{
					Source:      Position{Rank: 1, File: 0},
					Destination: Position{Rank: 2, File: 1},
					MoveType:    CAPTURE,
				},
				{
					Source:      Position{Rank: 7, File: 0},
					Destination: Position{Rank: 6, File: 1},
					MoveType:    CAPTURE,
				},
			},
			initialBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						1: {
							0: NewPawn(WHITE),
						},
						2: {
							1: NewPawn(BLACK),
						},
						6: {
							1: NewPawn(WHITE),
						},
						7: {
							0: NewPawn(BLACK),
						},
					},
				),
			),
			expectedBoard: Build(
				WithBounds(bounds),
				WithGameboardState(
					GameboardState{
						2: {
							1: NewPawn(WHITE),
						},
						6: {
							1: NewPawn(BLACK),
						},
					},
				),
			),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Create a new board for each test case to avoid modifying the initial
		board := tc.initialBoard

		// Call the HandleMove method on the
		for _, move := range tc.moves {
			err := board.HandleMove(move)
			if err != nil {
				t.Errorf("Test case %s: expected no error but got %v", tc.name, err)
			}
		}

		// Check that the board is in the expected state.
		if !reflect.DeepEqual(board.CastlingState, tc.expectedBoard.CastlingState) {
			t.Errorf("Test case %s: expected board %v but got %v", tc.name, tc.expectedBoard.CastlingState, board.CastlingState)
		}
	}
}

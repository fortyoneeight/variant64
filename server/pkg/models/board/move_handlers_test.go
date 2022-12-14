package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglePieceMoveApplicator(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	moveApplicator := SinglePieceMoveApplicator{}
	testcases := []struct {
		name          string
		move          Move
		state         GameboardState
		expectedState GameboardState
		expectedError error
	}{
		{
			name: "normal move",
			move: Move{
				Source:      Position{Rank: 6, File: 0},
				Destination: Position{Rank: 5, File: 0},
				MoveType:    NORMAL,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					6: {
						0: NewPawn(BLACK),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					5: {
						0: NewPawn(BLACK),
					},
				},
			),
		},
		{
			name: "pawn double push move",
			move: Move{
				Source:      Position{Rank: 6, File: 0},
				Destination: Position{Rank: 4, File: 0},
				MoveType:    PAWN_DOUBLE_PUSH,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					6: {
						0: NewPawn(BLACK),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					4: {
						0: NewPawn(BLACK),
					},
				},
			),
		},
		{
			name: "capture move",
			move: Move{
				Source:      Position{Rank: 3, File: 3},
				Destination: Position{Rank: 1, File: 2},
				MoveType:    CAPTURE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					3: {
						3: NewKnight(BLACK),
					},
					1: {
						2: NewKnight(WHITE),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					1: {
						2: NewKnight(BLACK),
					},
				},
			),
		},
		{
			name: "unsupport MoveType",
			move: Move{
				Source:      Position{Rank: 0, File: 4},
				Destination: Position{Rank: 0, File: 6},
				MoveType:    KINGSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						4: NewKing(BLACK),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						4: NewKing(BLACK),
					},
				},
			),
			expectedError: errCannotHandleMoveType(KINGSIDE_CASTLE),
		},
		{
			name: "nil piece at source",
			move: Move{
				Source:      Position{Rank: 0, File: 0},
				Destination: Position{Rank: 0, File: 1},
				MoveType:    NORMAL,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: nil,
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: nil,
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := moveApplicator.ApplyMove(tc.move, tc.state)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedState, tc.state)
		})
	}
}

func TestKingsideCastleMoveApplicator(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	moveApplicator := KingsideCastleMoveApplicator{}
	testcases := []struct {
		name          string
		move          Move
		state         GameboardState
		expectedState GameboardState
		expectedError error
	}{
		{
			name: "white kingside castle",
			move: Move{
				Source:      Position{Rank: 0, File: 4},
				Destination: Position{Rank: 0, File: 6},
				MoveType:    KINGSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: NewKing(WHITE),
						7: NewRook(WHITE, bounds),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						6: NewKing(WHITE),
						5: NewRook(WHITE, bounds),
					},
				},
			),
		},
		{
			name: "black kingside castle",
			move: Move{
				Source:      Position{Rank: 7, File: 4},
				Destination: Position{Rank: 7, File: 6},
				MoveType:    KINGSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						4: NewKing(BLACK),
						7: NewRook(BLACK, bounds),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						6: NewKing(BLACK),
						5: NewRook(BLACK, bounds),
					},
				},
			),
		},
		{
			name: "unsupported MoveType",
			move: Move{
				Source:      Position{Rank: 6, File: 0},
				Destination: Position{Rank: 5, File: 0},
				MoveType:    NORMAL,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					6: {
						0: NewPawn(BLACK),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					6: {
						0: NewPawn(BLACK),
					},
				},
			),
			expectedError: errCannotHandleMoveType(NORMAL),
		},
		{
			name: "white kingside castle, nil king piece",
			move: Move{
				Source:      Position{Rank: 0, File: 4},
				Destination: Position{Rank: 0, File: 6},
				MoveType:    KINGSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: nil,
						7: NewRook(WHITE, bounds),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: nil,
						7: NewRook(WHITE, bounds),
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
		{
			name: "white kingside castle, nil rook piece",
			move: Move{
				Source:      Position{Rank: 0, File: 4},
				Destination: Position{Rank: 0, File: 6},
				MoveType:    KINGSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: NewKing(WHITE),
						7: nil,
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: NewKing(WHITE),
						7: nil,
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
		{
			name: "black kingside castle, nil king piece",
			move: Move{
				Source:      Position{Rank: 7, File: 4},
				Destination: Position{Rank: 7, File: 6},
				MoveType:    KINGSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						4: nil,
						7: NewRook(BLACK, bounds),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						4: nil,
						7: NewRook(BLACK, bounds),
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
		{
			name: "black kingside castle, nil rook piece",
			move: Move{
				Source:      Position{Rank: 7, File: 4},
				Destination: Position{Rank: 7, File: 6},
				MoveType:    KINGSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						4: NewKing(BLACK),
						7: nil,
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						4: NewKing(BLACK),
						7: nil,
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := moveApplicator.ApplyMove(tc.move, tc.state)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedState, tc.state)
		})
	}
}

func TestQueensideCastleMoveApplicator(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	moveApplicator := QueensideCastleMoveApplicator{}
	testcases := []struct {
		name          string
		move          Move
		state         GameboardState
		expectedState GameboardState
		expectedError error
	}{
		{
			name: "white queenside castle",
			move: Move{
				Source:      Position{Rank: 0, File: 4},
				Destination: Position{Rank: 0, File: 2},
				MoveType:    QUEENSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: NewKing(WHITE),
						0: NewRook(WHITE, bounds),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						2: NewKing(WHITE),
						3: NewRook(WHITE, bounds),
					},
				},
			),
		},
		{
			name: "black queenside castle",
			move: Move{
				Source:      Position{Rank: 7, File: 4},
				Destination: Position{Rank: 7, File: 2},
				MoveType:    QUEENSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						4: NewKing(BLACK),
						0: NewRook(BLACK, bounds),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						2: NewKing(BLACK),
						3: NewRook(BLACK, bounds),
					},
				},
			),
		},
		{
			name: "unsupported MoveType",
			move: Move{
				Source:      Position{Rank: 6, File: 0},
				Destination: Position{Rank: 5, File: 0},
				MoveType:    NORMAL,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					6: {
						0: NewPawn(BLACK),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					6: {
						0: NewPawn(BLACK),
					},
				},
			),
			expectedError: errCannotHandleMoveType(NORMAL),
		},
		{
			name: "white queenside castle, nil king piece",
			move: Move{
				Source:      Position{Rank: 0, File: 4},
				Destination: Position{Rank: 0, File: 2},
				MoveType:    QUEENSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
						4: nil,
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
						4: nil,
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
		{
			name: "white queenside castle, nil rook piece",
			move: Move{
				Source:      Position{Rank: 0, File: 4},
				Destination: Position{Rank: 0, File: 2},
				MoveType:    QUEENSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: nil,
						4: NewKing(WHITE),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: nil,
						4: NewKing(WHITE),
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
		{
			name: "black queenside castle, nil king piece",
			move: Move{
				Source:      Position{Rank: 7, File: 4},
				Destination: Position{Rank: 7, File: 2},
				MoveType:    QUEENSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						0: NewRook(BLACK, bounds),
						4: nil,
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						0: NewRook(BLACK, bounds),
						4: nil,
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
		{
			name: "black queenside castle, nil rook piece",
			move: Move{
				Source:      Position{Rank: 7, File: 4},
				Destination: Position{Rank: 7, File: 2},
				MoveType:    QUEENSIDE_CASTLE,
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						0: nil,
						4: NewKing(BLACK),
					},
				},
			),
			expectedState: NewGameboardState(
				bounds,
				GameboardState{
					7: {
						0: nil,
						4: NewKing(BLACK),
					},
				},
			),
			expectedError: errSourcePieceNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := moveApplicator.ApplyMove(tc.move, tc.state)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedState, tc.state)
		})
	}
}

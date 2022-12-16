package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCastleState(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	boardState := NewGameboardState(
		bounds,
		GameboardState{
			0: {
				0: NewRook(WHITE, bounds),
				4: NewKing(WHITE),
				7: NewRook(WHITE, bounds),
			},
			7: {
				0: NewRook(BLACK, bounds),
				4: NewKing(BLACK),
				7: NewRook(BLACK, bounds),
			},
		},
	)
	testcases := []struct {
		name                  string
		castlingState         *CastlingState
		moves                 []Move
		expectedCastlingState *CastlingState
	}{
		{
			name:          "White king moved.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 1, File: 4},
					MoveType:    NORMAL,
				},
			},
			expectedCastlingState: createCastlingState(false, false, true, true),
		},
		{
			name:          "White kingside castle.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 0, File: 6},
					MoveType:    KINGSIDE_CASTLE,
				},
			},
			expectedCastlingState: createCastlingState(false, false, true, true),
		},
		{
			name:          "White queenside castle.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    QUEENSIDE_CASTLE,
				},
			},
			expectedCastlingState: createCastlingState(false, false, true, true),
		},
		{
			name:          "Black king moved.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 6, File: 4},
					MoveType:    NORMAL,
				},
			},
			expectedCastlingState: createCastlingState(true, true, false, false),
		},
		{
			name:          "Black kingside castle.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 7, File: 6},
					MoveType:    KINGSIDE_CASTLE,
				},
			},
			expectedCastlingState: createCastlingState(true, true, false, false),
		},
		{
			name:          "Black queenside castle.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 7, File: 2},
					MoveType:    QUEENSIDE_CASTLE,
				},
			},
			expectedCastlingState: createCastlingState(true, true, false, false),
		},
		{
			name:          "White queenside rook moves.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 1, File: 0},
					MoveType:    NORMAL,
				},
			},
			expectedCastlingState: createCastlingState(true, false, true, true),
		},
		{
			name:          "Black queenside rook moves.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 7, File: 0},
					Destination: Position{Rank: 6, File: 0},
					MoveType:    NORMAL,
				},
			},
			expectedCastlingState: createCastlingState(true, true, true, false),
		},
		{
			name:          "White kingside rook moves.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 7},
					Destination: Position{Rank: 1, File: 7},
					MoveType:    NORMAL,
				},
			},
			expectedCastlingState: createCastlingState(false, true, true, true),
		},
		{
			name:          "Black kingside rook moves.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 7, File: 7},
					Destination: Position{Rank: 6, File: 7},
					MoveType:    NORMAL,
				},
			},
			expectedCastlingState: createCastlingState(true, true, false, true),
		},
		{
			name:          "Unrelated moves.",
			castlingState: NewDefaultCastlingState(),
			moves: []Move{
				{
					Source:      Position{Rank: 5, File: 5},
					Destination: Position{Rank: 5, File: 6},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 5, File: 5},
					Destination: Position{Rank: 5, File: 6},
					MoveType:    JUMP,
				},
				{
					Source:      Position{Rank: 5, File: 5},
					Destination: Position{Rank: 5, File: 6},
					MoveType:    JUMP_CAPTURE,
				},
			},
			expectedCastlingState: NewDefaultCastlingState(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, move := range tc.moves {
				tc.castlingState.UpdateCastleState(move, boardState)
			}
			assert.Equal(t, tc.expectedCastlingState, tc.castlingState)
		})
	}
}

func createCastlingState(whiteKingside, whiteQueenside, blackKingside, blackQueenside bool) *CastlingState {
	return &CastlingState{
		CastlingStateMap: map[MoveType]map[Color]bool{
			KINGSIDE_CASTLE: {
				WHITE: whiteKingside,
				BLACK: blackKingside,
			},
			QUEENSIDE_CASTLE: {
				WHITE: whiteQueenside,
				BLACK: blackQueenside,
			},
		},
	}
}

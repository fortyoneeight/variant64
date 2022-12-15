package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterOutOfBounds(t *testing.T) {
	testcases := []struct {
		name                string
		filter              FilterOutOfBounds
		moves               []Move
		expectedIsLegalMove bool
	}{
		{
			name: "Out of bounds destination",
			filter: FilterOutOfBounds{
				Bounds: Bounds{RankCount: 8, FileCount: 8},
			},
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: -1, File: 0},
					MoveType:    NORMAL,
				},
				Move{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 0, File: -1},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 0, File: 10},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 10, File: 0},
					MoveType:    NORMAL,
				},
			},
			expectedIsLegalMove: false,
		},
		{
			name: "Out of bounds source",
			filter: FilterOutOfBounds{
				Bounds: Bounds{RankCount: 8, FileCount: 8},
			},
			moves: []Move{
				{
					Source:      Position{Rank: -1, File: 0},
					Destination: Position{Rank: 0, File: 0},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 0, File: -1},
					Destination: Position{Rank: 0, File: 0},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 0, File: 10},
					Destination: Position{Rank: 0, File: 0},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 10, File: 0},
					Destination: Position{Rank: 0, File: 0},
					MoveType:    NORMAL,
				},
			},
			expectedIsLegalMove: false,
		},
		{
			name: "Inbounds corners",
			filter: FilterOutOfBounds{
				Bounds: Bounds{RankCount: 8, FileCount: 8},
			},
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 0, File: 1},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 0, File: 7},
					Destination: Position{Rank: 0, File: 6},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 7, File: 0},
					Destination: Position{Rank: 6, File: 0},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 7, File: 7},
					Destination: Position{Rank: 6, File: 6},
					MoveType:    NORMAL,
				},
			},
			expectedIsLegalMove: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, move := range tc.moves {
				isLegalMove := tc.filter.IsLegalMove(move, GameboardState{})
				assert.Equal(t, tc.expectedIsLegalMove, isLegalMove)
			}
		})
	}
}

func TestFilterPieceCollision(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	testcases := []struct {
		name                string
		filter              FilterPieceCollision
		state               GameboardState
		moves               []Move
		expectedIsLegalMove bool
	}{
		{
			name:   "Capture collide allowed at destination.",
			filter: FilterPieceCollision{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
					},
					7: {
						0: NewRook(BLACK, bounds),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    CAPTURE,
				},
			},
			expectedIsLegalMove: true,
		},
		{
			name:   "Capture collide not allowed outside of destination.",
			filter: FilterPieceCollision{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
					},
					6: {
						0: NewRook(BLACK, bounds),
					},
					7: {
						0: NewRook(BLACK, bounds),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    CAPTURE,
				},
			},
			expectedIsLegalMove: false,
		},
		{
			name:   "Normal collide not allowed.",
			filter: FilterPieceCollision{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
					},
					1: {
						0: NewRook(BLACK, bounds),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    NORMAL,
				},
			},
			expectedIsLegalMove: false,
		},
		{
			name:   "Double push collide not allowed.",
			filter: FilterPieceCollision{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					6: {
						0: NewPawn(BLACK),
					},
					5: {
						0: NewPawn(BLACK),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 6, File: 0},
					Destination: Position{Rank: 5, File: 0},
					MoveType:    PAWN_DOUBLE_PUSH,
				},
			},
			expectedIsLegalMove: false,
		},
		{
			name:   "Unsupported move types.",
			filter: FilterPieceCollision{},
			state:  GameboardState{},
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 0, File: 6},
					MoveType:    KINGSIDE_CASTLE,
				},
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 7, File: 2},
					MoveType:    QUEENSIDE_CASTLE,
				},
			},
			expectedIsLegalMove: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, move := range tc.moves {
				isLegalMove := tc.filter.IsLegalMove(move, tc.state)
				assert.Equal(t, tc.expectedIsLegalMove, isLegalMove)
			}
		})
	}
}

func TestFilterFriendlyCapture(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	testcases := []struct {
		name                string
		filter              FilterFriendlyCapture
		state               GameboardState
		moves               []Move
		expectedIsLegalMove bool
	}{
		{
			name:   "Capture empty position not allowed.",
			filter: FilterFriendlyCapture{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    CAPTURE,
				},
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    JUMP_CAPTURE,
				},
			},
			expectedIsLegalMove: false,
		},
		{
			name:   "Capture occupied position allowed.",
			filter: FilterFriendlyCapture{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
					},
					7: {
						0: NewRook(BLACK, bounds),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    CAPTURE,
				},
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    JUMP_CAPTURE,
				},
			},
			expectedIsLegalMove: true,
		},
		{
			name:   "Capture friendly position not allowed.",
			filter: FilterFriendlyCapture{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
					},
					7: {
						0: NewRook(WHITE, bounds),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    CAPTURE,
				},
				{
					Source:      Position{Rank: 0, File: 0},
					Destination: Position{Rank: 7, File: 0},
					MoveType:    JUMP_CAPTURE,
				},
			},
			expectedIsLegalMove: false,
		},
		{
			name:   "Unsupported move types.",
			filter: FilterFriendlyCapture{},
			state:  GameboardState{},
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 0, File: 6},
					MoveType:    KINGSIDE_CASTLE,
				},
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 7, File: 2},
					MoveType:    QUEENSIDE_CASTLE,
				},
			},
			expectedIsLegalMove: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, move := range tc.moves {
				isLegalMove := tc.filter.IsLegalMove(move, tc.state)
				assert.Equal(t, tc.expectedIsLegalMove, isLegalMove)
			}
		})
	}
}

func TestFilterInvalidPawnDoublePush(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	testcases := []struct {
		name                string
		filter              FilterInvalidPawnDoublePush
		state               GameboardState
		moves               []Move
		expectedIsLegalMove bool
	}{
		{
			name:   "Pawn double push at initial position allowed.",
			filter: FilterInvalidPawnDoublePush{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					1: {
						0: NewPawn(WHITE),
					},
					6: {
						0: NewPawn(BLACK),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 1, File: 0},
					Destination: Position{Rank: 2, File: 0},
					MoveType:    PAWN_DOUBLE_PUSH,
				},
				{
					Source:      Position{Rank: 6, File: 0},
					Destination: Position{Rank: 5, File: 0},
					MoveType:    PAWN_DOUBLE_PUSH,
				},
			},
			expectedIsLegalMove: true,
		},
		{
			name:   "Pawn double push not at initial position not allowed.",
			filter: FilterInvalidPawnDoublePush{},
			state: NewGameboardState(
				bounds,
				GameboardState{
					2: {
						0: NewPawn(WHITE),
					},
					4: {
						0: NewPawn(BLACK),
					},
				},
			),
			moves: []Move{
				{
					Source:      Position{Rank: 2, File: 0},
					Destination: Position{Rank: 4, File: 0},
					MoveType:    PAWN_DOUBLE_PUSH,
				},
				{
					Source:      Position{Rank: 4, File: 0},
					Destination: Position{Rank: 2, File: 0},
					MoveType:    PAWN_DOUBLE_PUSH,
				},
			},
			expectedIsLegalMove: false,
		},
		{
			name:   "Unsupported move types.",
			filter: FilterInvalidPawnDoublePush{},
			state:  GameboardState{},
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    CAPTURE,
				},
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    JUMP,
				},
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    JUMP_CAPTURE,
				},
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 0, File: 6},
					MoveType:    KINGSIDE_CASTLE,
				},
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 7, File: 2},
					MoveType:    QUEENSIDE_CASTLE,
				},
			},
			expectedIsLegalMove: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, move := range tc.moves {
				isLegalMove := tc.filter.IsLegalMove(move, tc.state)
				assert.Equal(t, tc.expectedIsLegalMove, isLegalMove)
			}
		})
	}
}

func TestFilterIllegalKingsideCastle(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	testcases := []struct {
		name                string
		filter              FilterIllegalKingsideCastle
		state               GameboardState
		moves               []Move
		expectedIsLegalMove bool
	}{
		{
			name: "Valid kingside castle move allowed.",
			filter: FilterIllegalKingsideCastle{
				CastlingState: createCastlingState(true, true, true, true),
			},
			state: NewGameboardState(
				bounds,
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
			expectedIsLegalMove: true,
		},
		{
			name: "Obstructed kingside castle move not allowed.",
			filter: FilterIllegalKingsideCastle{
				CastlingState: createCastlingState(true, true, true, true),
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: NewKing(WHITE),
						5: NewKnight(WHITE),
						7: NewRook(WHITE, bounds),
					},
					1: {
						4: NewKing(BLACK),
						5: NewKnight(BLACK),
						7: NewRook(BLACK, bounds),
					},
				},
			),
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
			expectedIsLegalMove: false,
		},
		{
			name: "Kingside castle move with queenside castle disabled allowed.",
			filter: FilterIllegalKingsideCastle{
				CastlingState: createCastlingState(true, false, true, false),
			},
			state: NewGameboardState(
				bounds,
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
			expectedIsLegalMove: true,
		},
		{
			name: "Kingside castle move with kingside castle disabled not allowed.",
			filter: FilterIllegalKingsideCastle{
				CastlingState: createCastlingState(false, true, false, true),
			},
			state: NewGameboardState(
				bounds,
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
			expectedIsLegalMove: false,
		},
		{
			name: "Non-kingside castle move allowed.",
			filter: FilterIllegalKingsideCastle{
				CastlingState: createCastlingState(true, true, true, true),
			},
			state: NewGameboardState(
				bounds,
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
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 1, File: 5},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 7, File: 4},
					Destination: Position{Rank: 6, File: 5},
					MoveType:    NORMAL,
				},
			},
			expectedIsLegalMove: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, move := range tc.moves {
				isLegalMove := tc.filter.IsLegalMove(move, tc.state)
				assert.Equal(t, tc.expectedIsLegalMove, isLegalMove)
			}
		})
	}
}

func TestFilterIllegalQueensideCastle(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	testcases := []struct {
		name                string
		filter              FilterIllegalQueensideCastle
		state               GameboardState
		moves               []Move
		expectedIsLegalMove bool
	}{
		{
			name: "Valid queenside castle move allowed.",
			filter: FilterIllegalQueensideCastle{
				CastlingState: createCastlingState(true, true, true, true),
			},
			state: NewGameboardState(
				bounds,
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
			expectedIsLegalMove: true,
		},
		{
			name: "Obstructed queenside castle move not allowed.",
			filter: FilterIllegalQueensideCastle{
				CastlingState: createCastlingState(true, true, true, true),
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewRook(WHITE, bounds),
						1: NewKnight(WHITE),
						4: NewRook(WHITE, bounds),
					},
					1: {
						0: NewRook(BLACK, bounds),
						1: NewKnight(BLACK),
						4: NewRook(BLACK, bounds),
					},
				},
			),
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
			expectedIsLegalMove: false,
		},
		{
			name: "Queenside castle move with kingside castle disabled allowed.",
			filter: FilterIllegalQueensideCastle{
				CastlingState: createCastlingState(false, true, false, true),
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: NewKing(WHITE),
						0: NewRook(WHITE, bounds),
					},
					7: {
						4: NewKing(BLACK),
						0: NewRook(BLACK, bounds),
					},
				},
			),
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
			expectedIsLegalMove: true,
		},
		{
			name: "Queenside castle move with queenside castle disabled not allowed.",
			filter: FilterIllegalQueensideCastle{
				CastlingState: createCastlingState(true, false, true, false),
			},
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						4: NewKing(WHITE),
						0: NewRook(WHITE, bounds),
					},
					7: {
						4: NewKing(BLACK),
						0: NewRook(BLACK, bounds),
					},
				},
			),
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
			expectedIsLegalMove: false,
		},
		{
			name:   "Unsupported move types.",
			filter: FilterIllegalQueensideCastle{},
			state:  GameboardState{},
			moves: []Move{
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    NORMAL,
				},
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    CAPTURE,
				},
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    JUMP,
				},
				{
					Source:      Position{Rank: 0, File: 1},
					Destination: Position{Rank: 0, File: 2},
					MoveType:    JUMP_CAPTURE,
				},
				{
					Source:      Position{Rank: 0, File: 4},
					Destination: Position{Rank: 0, File: 6},
					MoveType:    KINGSIDE_CASTLE,
				},
			},
			expectedIsLegalMove: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, move := range tc.moves {
				isLegalMove := tc.filter.IsLegalMove(move, tc.state)
				assert.Equal(t, tc.expectedIsLegalMove, isLegalMove)
			}
		})
	}
}

func createCastlingState(whiteKingside, whiteQueenside, blackKingside, blackQueenside bool) *CastlingState {
	return &CastlingState{
		map[MoveType]map[Color]bool{
			KINGSIDE_CASTLE: map[Color]bool{
				WHITE: whiteKingside,
				BLACK: blackKingside,
			},
			QUEENSIDE_CASTLE: map[Color]bool{
				WHITE: whiteQueenside,
				BLACK: blackQueenside,
			},
		},
	}
}

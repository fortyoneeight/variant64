package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIllegalCheckStateFilter(t *testing.T) {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	testcases := []struct {
		name                 string
		color                Color
		state                GameboardState
		availableMoveMap     AvailableMoveMap
		expectedIsLegalState bool
	}{
		{
			name:                 "Empty board is legal state.",
			color:                WHITE,
			state:                NewGameboardState(bounds, GameboardState{}),
			availableMoveMap:     NewAvailableMoveMap(bounds),
			expectedIsLegalState: true,
		},
		{
			name:  "Only king on board is legal state.",
			color: BLACK,
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewKing(BLACK),
					},
				}),
			availableMoveMap:     NewAvailableMoveMap(bounds),
			expectedIsLegalState: true,
		},
		{
			name:  "King in check via CAPTURE is illegal state.",
			color: BLACK,
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewKing(BLACK),
						7: NewRook(WHITE, bounds),
					},
				}),
			availableMoveMap: AvailableMoveMap{
				0: {
					7: MoveMap{
						CAPTURE: []Position{
							{Rank: 0, File: 0},
						},
					},
				},
			},
			expectedIsLegalState: false,
		},
		{
			name:  "King in check via JUMP_CAPTURE is illegal state.",
			color: BLACK,
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewKing(BLACK),
						1: NewQueen(BLACK, bounds),
					},
					2: {
						1: NewKnight(WHITE),
					},
				}),
			availableMoveMap: AvailableMoveMap{
				0: {
					7: MoveMap{
						JUMP_CAPTURE: []Position{
							{Rank: 0, File: 0},
						},
					},
				},
			},
			expectedIsLegalState: false,
		},
		{
			name:  "King in check but other color is legal state.",
			color: WHITE,
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewKing(BLACK),
						7: NewRook(WHITE, bounds),
					},
				}),
			availableMoveMap: AvailableMoveMap{
				0: {
					7: MoveMap{
						CAPTURE: []Position{
							{Rank: 0, File: 0},
						},
					},
				},
			},
			expectedIsLegalState: true,
		},
		{
			name:  "Capturing non-king piece is legal state.",
			color: BLACK,
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewKing(BLACK),
						1: NewQueen(BLACK, bounds),
						7: NewRook(WHITE, bounds),
					},
				}),
			availableMoveMap: AvailableMoveMap{
				0: {
					7: MoveMap{
						CAPTURE: []Position{
							{Rank: 0, File: 1},
						},
					},
				},
			},
			expectedIsLegalState: true,
		},
		{
			name:  "Jump capturing non-king piece is legal state.",
			color: BLACK,
			state: NewGameboardState(
				bounds,
				GameboardState{
					0: {
						0: NewKing(BLACK),
						1: NewQueen(BLACK, bounds),
					},
					2: {
						2: NewKnight(WHITE),
					},
				}),
			availableMoveMap: AvailableMoveMap{
				0: {
					7: MoveMap{
						JUMP_CAPTURE: []Position{
							{Rank: 0, File: 1},
						},
					},
				},
			},
			expectedIsLegalState: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			illegalCheckStateFilter := IllegalCheckStateFilter{
				TurnState: &TurnState{
					Active:    BLACK,
					TurnOrder: []Color{WHITE, BLACK},
				},
			}
			isLegalState := illegalCheckStateFilter.IsLegalState(tc.color, tc.state, tc.availableMoveMap)
			assert.Equal(t, tc.expectedIsLegalState, isLegalState)
		})
	}
}

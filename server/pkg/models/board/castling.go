package board

// CastlingState is used to represent which types of
// castling moves are still allowed for each player.
type CastlingState struct {
	CastlingStateMap map[MoveType]map[Color]bool
}

// NewCastlingState creates a new CastlingState
// all the values are initialized to true.
func NewCastlingState() *CastlingState {
	return &CastlingState{
		CastlingStateMap: map[MoveType]map[Color]bool{
			KINGSIDE_CASTLE: {
				WHITE: true,
				BLACK: true,
			},
			QUEENSIDE_CASTLE: {
				WHITE: true,
				BLACK: true,
			},
		},
	}
}

// IsAllowed returns true if the provided MoveType is allowed for the player Color.
func (c *CastlingState) IsAllowed(moveType MoveType, color Color) bool {
	if colorMap, ok := c.CastlingStateMap[moveType]; ok {
		if val, ok := colorMap[color]; ok {
			return val
		}
	}
	return false
}

// Disallow prevents the provided player Color from making the provided MoveType.
func (c *CastlingState) Disallow(moveType MoveType, color Color) {
	c.CastlingStateMap[moveType][color] = false
}

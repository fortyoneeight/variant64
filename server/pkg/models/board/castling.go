package board

// CastlingState is used to represent which types of
// castling moves are still allowed for each player.
type CastlingState struct {
	CastlingStateMap map[MoveType]map[Color]bool
}

// NewCastlingState creates a new CastlingState
// all the values are initialized to true.
func NewDefaultCastlingState() *CastlingState {
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
func (c *CastlingState) UpdateCastleState(move Move, state GameboardState) {
	sourcePiece := state[move.Source.Rank][move.Source.File]
	if sourcePiece == nil {
		return
	}
	if sourcePiece.GetType() == KING {
		switch sourcePiece.GetColor() {
		case BLACK:
			c.disallow(KINGSIDE_CASTLE, BLACK)
			c.disallow(QUEENSIDE_CASTLE, BLACK)
		case WHITE:
			c.disallow(KINGSIDE_CASTLE, WHITE)
			c.disallow(QUEENSIDE_CASTLE, WHITE)
		}
	} else if sourcePiece.GetType() == ROOK {
		switch sourcePiece.GetColor() {
		case WHITE:
			if move.Source.Rank == 0 && move.Source.File == 0 {
				c.disallow(QUEENSIDE_CASTLE, WHITE)
			} else if move.Source.Rank == 0 && move.Source.File == 7 {
				c.disallow(KINGSIDE_CASTLE, WHITE)
			}
		case BLACK:
			if move.Source.Rank == 7 && move.Source.File == 0 {
				c.disallow(QUEENSIDE_CASTLE, BLACK)
			} else if move.Source.Rank == 7 && move.Source.File == 7 {
				c.disallow(KINGSIDE_CASTLE, BLACK)
			}
		}
	}
}

func (c *CastlingState) disallow(moveType MoveType, color Color) {
	c.CastlingStateMap[moveType][color] = false
}

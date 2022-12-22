package board

// positionPredicate returns true if the position satisfies the predicate.
type positionPredicate = func(position Position, state GameboardState) bool

// predicateAttackingEnemyKing returns a positionPredicate that checks
// if the provided piece is attacking an enemy king.
func predicateAttackingEnemyKing(
	kingColor Color,
	state GameboardState,
	availableMoves AvailableMoveMap,
) positionPredicate {
	return func(position Position, state GameboardState) bool {
		moves := availableMoves[position.Rank][position.File]
		if moves == nil {
			return false
		}

		capturingMoveTypes := []MoveType{CAPTURE, JUMP_CAPTURE}
		for _, moveType := range capturingMoveTypes {
			if movesByType, ok := moves[moveType]; ok {
				for _, destination := range movesByType {
					destinationPiece := state[destination.Rank][destination.File]
					if destinationPiece != nil &&
						destinationPiece.PieceType == KING &&
						destinationPiece.Color == kingColor {
						return true
					}
				}
			}
		}

		return false
	}
}

// anyPosition applies the predicate to each position on the board,
// returns true if any position satisfies the predicate.
func anyPosition(
	kingColor Color,
	state GameboardState,
	pred positionPredicate,
) bool {
	for rank, files := range state {
		for file := range files {
			result := pred(Position{Rank: rank, File: file}, state)
			if result {
				return result
			}
		}
	}
	return false
}

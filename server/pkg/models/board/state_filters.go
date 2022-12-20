package board

type illegalStateFilter interface {
	IsLegalState(color Color, state GameboardState, availableMoveMap AvailableMoveMap) bool
}

// IllegalStateFilter bundles multiple IllegalStateFilter into a single struct.
type IllegalStateFilter struct {
	illegalStateFilters []illegalStateFilter
}

func (i *IllegalStateFilter) IsLegalState(
	color Color,
	state GameboardState,
	availableMoveMap AvailableMoveMap,
) bool {
	for _, filter := range i.illegalStateFilters {
		if !filter.IsLegalState(color, state, availableMoveMap) {
			return false
		}
	}
	return true
}

func NewIllegalStateFilter(illegalStateFilters ...illegalStateFilter) *IllegalStateFilter {
	return &IllegalStateFilter{
		illegalStateFilters: illegalStateFilters,
	}
}

// IllegalCheckStateFilter is used to verify if a king is in check.
type IllegalCheckStateFilter struct{}

// IsLegalState checks is the provided Color's king is in check.
func (s *IllegalCheckStateFilter) IsLegalState(
	color Color,
	state GameboardState,
	availableMoveMap AvailableMoveMap,
) bool {
	return !s.anyPosition(
		color,
		state,
		s.PredicateAttackingEnemyKing(color, state, availableMoveMap),
	)
}

// PredicateAttackingEnemyKing returns a positionPredicate that checks
// if the provided piece is attacking an enemy king.
func (s *IllegalCheckStateFilter) PredicateAttackingEnemyKing(
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

// positionPredicate returns true if the position satisfies the predicate.
type positionPredicate = func(position Position, state GameboardState) bool

// anyPosition applies the predicate to each position on the board,
// returns true if any position satisfies the predicate.
func (s *IllegalCheckStateFilter) anyPosition(
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

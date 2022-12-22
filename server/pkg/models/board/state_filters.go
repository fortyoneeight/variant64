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
type IllegalCheckStateFilter struct {
	*TurnState
}

// IsLegalState checks is the provided Color's king is in check.
func (s *IllegalCheckStateFilter) IsLegalState(
	color Color,
	state GameboardState,
	availableMoveMap AvailableMoveMap,
) bool {
	// Only check illegal states for active player.
	if color != s.GetActivePlayer() {
		return true
	}

	return !anyPosition(
		color,
		state,
		predicateAttackingEnemyKing(color, state, availableMoveMap),
	)
}

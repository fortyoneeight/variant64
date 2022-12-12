package board

// SingleNormalMoveGenerator generates NORMAL moves of one square
// in a single direction.
type SingleNormalMoveGenerator struct {
	direction Direction
}

func (g *SingleNormalMoveGenerator) GetMoves(source Position) MoveMap {
	return map[MoveType][]Position{
		NORMAL: {
			StepInDirection(source, g.direction),
		},
	}
}

// SingleDiagonalCaputureMoveGenerator generates CAPTURE moves of one square
// in each of the forward two diagonal directions.
type SingleDiagonalCaputureMoveGenerator struct {
	color Color
}

func (g *SingleDiagonalCaputureMoveGenerator) GetMoves(source Position) MoveMap {
	var rankDirection int
	if g.color == WHITE {
		rankDirection = 1
	} else {
		rankDirection = -1
	}

	return map[MoveType][]Position{
		CAPTURE: {
			{Rank: source.Rank + rankDirection, File: source.File - 1},
			{Rank: source.Rank + rankDirection, File: source.File + 1},
		},
	}
}

// DoublePushMoveGenerator generates PAWN_DOUBLE_PUSH moves of two squares
// in a single direction.
type DoublePushMoveGenerator struct {
	color Color
}

func (g *DoublePushMoveGenerator) GetMoves(source Position) MoveMap {
	var rankDirection int
	if g.color == WHITE {
		rankDirection = 1
	} else {
		rankDirection = -1
	}

	return map[MoveType][]Position{
		PAWN_DOUBLE_PUSH: {
			{Rank: source.Rank + rankDirection*2, File: source.File},
		},
	}
}

// KnightMoveGenerator generates NORMAL moves in an L-shape of two squares
// in one direction and one square in an orthogonal direction.
type KnightMoveGenerator struct{}

func (g *KnightMoveGenerator) GetMoves(source Position) MoveMap {
	moves := map[MoveType][]Position{
		NORMAL:  {},
		CAPTURE: {},
	}

	nextPositionList := []Position{
		{Rank: source.Rank + 2, File: source.File + 1},
		{Rank: source.Rank + 2, File: source.File - 1},
		{Rank: source.Rank + 1, File: source.File + 2},
		{Rank: source.Rank + 1, File: source.File - 2},
		{Rank: source.Rank - 1, File: source.File + 2},
		{Rank: source.Rank - 1, File: source.File - 2},
		{Rank: source.Rank - 2, File: source.File + 1},
		{Rank: source.Rank - 2, File: source.File - 1},
	}

	for _, nextPosition := range nextPositionList {
		moves[NORMAL] = append(moves[NORMAL], nextPosition)
		moves[CAPTURE] = append(moves[CAPTURE], nextPosition)
	}

	return moves
}

// RayMoveGenerator generates RAY moves in a single direction where
// the generated position is the edge of the provided bounds.
type RayMoveGenerator struct {
	direction Direction
	bounds    Bounds
}

func (g *RayMoveGenerator) GetMoves(source Position) MoveMap {
	return map[MoveType][]Position{
		RAY: {generateTerminalRayPosition(source, g.direction, g.bounds)},
	}
}

// CastleMoveGenerator generates KINGSIDE_CASTLE and QUEENSIDE_CASTLE moves.
type CastleMoveGenerator struct {
	color Color
}

func (g *CastleMoveGenerator) GetMoves(source Position) MoveMap {
	moves := map[MoveType][]Position{
		KINGSIDE_CASTLE:  {},
		QUEENSIDE_CASTLE: {},
	}

	castleOptions := []struct {
		requiredColor Color
		moveType      MoveType
		destination   Position
	}{
		{
			WHITE,
			KINGSIDE_CASTLE,
			Position{Rank: 0, File: 6},
		},
		{
			WHITE,
			QUEENSIDE_CASTLE,
			Position{Rank: 0, File: 2},
		},
		{
			BLACK,
			KINGSIDE_CASTLE,
			Position{Rank: 7, File: 6},
		},
		{
			BLACK,
			QUEENSIDE_CASTLE,
			Position{Rank: 7, File: 2},
		},
	}
	for _, option := range castleOptions {
		if option.requiredColor == g.color {
			moves[option.moveType] = append(moves[option.moveType], option.destination)
		}
	}

	return moves
}

// StepInDirection returns a Position one move in the direction from the source Position.
func StepInDirection(source Position, direction Direction) Position {
	nextPosition := source
	switch direction {
	case North:
		nextPosition.Rank += 1
	case NorthEast:
		nextPosition.Rank += 1
		nextPosition.File += 1
	case East:
		nextPosition.File += 1
	case SouthEast:
		nextPosition.Rank -= 1
		nextPosition.File += 1
	case South:
		nextPosition.Rank -= 1
	case SouthWest:
		nextPosition.Rank -= 1
		nextPosition.File -= 1
	case West:
		nextPosition.File -= 1
	case NorthWest:
		nextPosition.Rank += 1
		nextPosition.File -= 1
	case None:
		return source
	}
	return nextPosition
}

// GenerateRay generates a list of Positions by moving from the source in the provided direction,
// it continues until a board boundary is encountered.
func GenerateRay(source Position, direction Direction, bounds Bounds) []Position {
	positionList := []Position{}

	nextPosition := source
	for {
		nextPosition = StepInDirection(nextPosition, direction)

		if !bounds.IsInboundsPosition(nextPosition) {
			return positionList
		}

		positionList = append(positionList, nextPosition)
	}
}

// generateTerminalRayPosition determines the terminal position in a ray by moving in the
// provided direction until the bounds are breached.
func generateTerminalRayPosition(source Position, direction Direction, bounds Bounds) Position {
	previousTerminalPosition := source
	terminalPosition := source
	for {
		terminalPosition = StepInDirection(terminalPosition, direction)

		if bounds.IsInboundsPosition(terminalPosition) {
			previousTerminalPosition = terminalPosition
		} else {
			return previousTerminalPosition
		}
	}
}

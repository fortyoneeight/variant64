package board

type Pawn struct {
	Color Color
}

func (p *Pawn) GetType() PieceType {
	return PAWN
}

func (p *Pawn) GetColor() Color {
	return p.Color
}

func (p *Pawn) GetMoves(source Position) MoveMap {
	moves := NewMoveMap()

	var rankDirection int
	if p.Color == WHITE {
		rankDirection = 1
	} else {
		rankDirection = -1
	}

	doublePushPosition := Position{Rank: source.Rank + rankDirection*2, File: source.File}
	moves[PAWN_DOUBLE_PUSH] = append(moves[PAWN_DOUBLE_PUSH], doublePushPosition)

	forwardPosition := Position{Rank: source.Rank + rankDirection, File: source.File}
	moves[NORMAL] = append(moves[NORMAL], forwardPosition)

	leftDiagonalCapture := Position{Rank: source.Rank + rankDirection, File: source.File - 1}
	moves[CAPTURE] = append(moves[CAPTURE], leftDiagonalCapture)

	rightDiagonalCapture := Position{Rank: source.Rank + rankDirection, File: source.File + 1}
	moves[CAPTURE] = append(moves[CAPTURE], rightDiagonalCapture)

	return moves
}

type Knight struct {
	Color Color
}

func (k *Knight) GetType() PieceType {
	return KNIGHT
}

func (k *Knight) GetColor() Color {
	return k.Color
}

func (k *Knight) GetMoves(source Position) MoveMap {
	moves := NewMoveMap()

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

type Rook struct {
	Bounds
	Color Color
}

func (r *Rook) GetType() PieceType {
	return ROOK
}

func (r *Rook) GetColor() Color {
	return r.Color
}

func (r *Rook) GetMoves(source Position) MoveMap {
	moves := NewMoveMap()

	rays := []Position{
		GenerateTerminalRayPosition(source, North, r.Bounds),
		GenerateTerminalRayPosition(source, East, r.Bounds),
		GenerateTerminalRayPosition(source, South, r.Bounds),
		GenerateTerminalRayPosition(source, West, r.Bounds),
	}

	for _, position := range rays {
		moves[RAY] = append(moves[RAY], position)
	}

	return moves
}

type Bishop struct {
	Bounds
	Color Color
}

func (b *Bishop) GetType() PieceType {
	return BISHOP
}

func (b *Bishop) GetColor() Color {
	return b.Color
}

func (b *Bishop) GetMoves(source Position) MoveMap {
	moves := NewMoveMap()

	rays := []Position{
		GenerateTerminalRayPosition(source, NorthEast, b.Bounds),
		GenerateTerminalRayPosition(source, SouthEast, b.Bounds),
		GenerateTerminalRayPosition(source, SouthWest, b.Bounds),
		GenerateTerminalRayPosition(source, NorthWest, b.Bounds),
	}

	for _, position := range rays {
		moves[RAY] = append(moves[RAY], position)
	}

	return moves
}

type Queen struct {
	Bounds
	Color Color
}

func (q *Queen) GetType() PieceType {
	return QUEEN
}

func (q *Queen) GetColor() Color {
	return q.Color
}

func (q *Queen) GetMoves(source Position) MoveMap {
	moves := NewMoveMap()

	rays := []Position{
		GenerateTerminalRayPosition(source, North, q.Bounds),
		GenerateTerminalRayPosition(source, NorthEast, q.Bounds),
		GenerateTerminalRayPosition(source, East, q.Bounds),
		GenerateTerminalRayPosition(source, SouthEast, q.Bounds),
		GenerateTerminalRayPosition(source, South, q.Bounds),
		GenerateTerminalRayPosition(source, SouthWest, q.Bounds),
		GenerateTerminalRayPosition(source, West, q.Bounds),
		GenerateTerminalRayPosition(source, NorthWest, q.Bounds),
	}

	for _, position := range rays {
		moves[RAY] = append(moves[RAY], position)
	}

	return moves
}

type King struct {
	Bounds
	Color Color
}

func (k *King) GetType() PieceType {
	return KING
}

func (k *King) GetColor() Color {
	return k.Color
}

func (k *King) GetMoves(source Position) MoveMap {
	moves := NewMoveMap()

	nextPositionList := []Position{
		{Rank: source.Rank + 1, File: source.File},
		{Rank: source.Rank + 1, File: source.File + 1},
		{Rank: source.Rank, File: source.File + 1},
		{Rank: source.Rank - 1, File: source.File + 1},
		{Rank: source.Rank - 1, File: source.File},
		{Rank: source.Rank - 1, File: source.File - 1},
		{Rank: source.Rank, File: source.File - 1},
		{Rank: source.Rank + 1, File: source.File - 1},
	}

	for _, nextPosition := range nextPositionList {
		moves[NORMAL] = append(moves[NORMAL], nextPosition)
		moves[CAPTURE] = append(moves[CAPTURE], nextPosition)
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
		if option.requiredColor == k.Color {
			moves[option.moveType] = append(moves[option.moveType], option.destination)
		}
	}

	return moves
}

// GenerateTerminalRayPosition determines the terminal position in a ray by moving in the
// provided direction until the bounds are breached.
func GenerateTerminalRayPosition(source Position, direction Direction, bounds Bounds) Position {
	previousTerminalPostion := source
	terminalPosition := source
	for {
		switch direction {
		case North:
			terminalPosition.Rank += 1
		case NorthEast:
			terminalPosition.Rank += 1
			terminalPosition.File += 1
		case East:
			terminalPosition.File += 1
		case SouthEast:
			terminalPosition.Rank -= 1
			terminalPosition.File += 1
		case South:
			terminalPosition.Rank -= 1
		case SouthWest:
			terminalPosition.Rank -= 1
			terminalPosition.File -= 1
		case West:
			terminalPosition.File -= 1
		case NorthWest:
			terminalPosition.Rank += 1
			terminalPosition.File -= 1
		}
		if bounds.IsInboundsPosition(terminalPosition) {
			previousTerminalPostion = terminalPosition
		} else {
			return previousTerminalPostion
		}
	}
}

// GenerateRay generates a list of Positions by moving from the source in the provided direction,
// it continues until a board boundary is encountered.
func GenerateRay(source Position, direction Direction, bounds Bounds) []Position {
	positionList := []Position{}

	nextPosition := source
	for {
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
			return []Position{}
		}

		if !bounds.IsInboundsPosition(nextPosition) {
			return positionList
		}

		positionList = append(positionList, nextPosition)
	}
}

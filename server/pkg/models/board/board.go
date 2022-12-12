package board

type GameboardType string

const (
	GameboardTypeDefault GameboardType = ""
	GameboardTypeClassic GameboardType = "classic"
)

type PieceType int

const (
	NONE PieceType = iota
	PAWN
	ROOK
	BISHOP
	KNIGHT
	KING
	QUEEN
)

type Color int

const (
	BLACK = iota
	WHITE
)

func (c Color) String() string {
	switch c {
	case WHITE:
		return "white"
	case BLACK:
		return "black"
	}
	return "invalid"
}

type MoveType int

const (
	NORMAL MoveType = iota
	RAY
	CAPTURE
	PAWN_DOUBLE_PUSH
	QUEENSIDE_CASTLE
	KINGSIDE_CASTLE
	PROMOTION
	EN_PASSANT
)

func (m MoveType) String() string {
	switch m {
	case NORMAL:
		return "normal"
	case RAY:
		return "ray"
	case CAPTURE:
		return "capture"
	case PAWN_DOUBLE_PUSH:
		return "pawn_double_push"
	case QUEENSIDE_CASTLE:
		return "queenside_castle"
	case KINGSIDE_CASTLE:
		return "kingside_castle"
	case PROMOTION:
		return "promotion"
	case EN_PASSANT:
		return "en_passant"
	}
	return "invalid"
}

type Move struct {
	Source      Position
	Destination Position
	MoveType    MoveType
}

type MoveMap = map[MoveType][]Position

func NewMoveMap() MoveMap {
	return map[MoveType][]Position{
		NORMAL:           make([]Position, 0),
		CAPTURE:          make([]Position, 0),
		PAWN_DOUBLE_PUSH: make([]Position, 0),
		RAY:              make([]Position, 0),
		KINGSIDE_CASTLE:  make([]Position, 0),
		QUEENSIDE_CASTLE: make([]Position, 0),
	}
}

func JoinMoveMaps(left, right MoveMap) {
	for key := range left {
		left[key] = append(left[key], right[key]...)
	}
}

type Direction = int

func GetDirection(source, destination Position) Direction {
	if source == destination {
		return None
	}

	horizontalDiff := destination.File - source.File
	verticalDiff := destination.Rank - source.Rank

	switch {
	case verticalDiff == 0:
		if horizontalDiff > 0 {
			return East
		}
		return West
	case verticalDiff > 0:
		if horizontalDiff == 0 {
			return North
		} else if horizontalDiff > 0 {
			return NorthEast
		}
		return NorthWest
	default:
		if horizontalDiff == 0 {
			return South
		} else if horizontalDiff > 0 {
			return SouthEast
		}
		return SouthWest
	}
}

const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
	None
)

type Position struct {
	Rank int
	File int
}

type Bounds struct {
	Rank int
	File int
}

func (b Bounds) IsInboundsPosition(position Position) bool {
	return position.Rank >= 0 &&
		position.File >= 0 &&
		position.Rank < b.Rank &&
		position.File < b.File
}

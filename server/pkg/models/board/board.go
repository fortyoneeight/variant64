package board

import "encoding/json"

type GameboardType string

const (
	GameboardTypeDefault GameboardType = ""
	GameboardTypeClassic GameboardType = "classic"
)

type GameboardState = map[int]map[int]Square

type Square struct {
	Piece          *Piece   `json:"piece,omitempty"`
	AvailableMoves *MoveMap `json:"available_moves,omitempty"`
}

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

func (p PieceType) String() string {
	switch p {
	case NONE:
		return "none"
	case PAWN:
		return "pawn"
	case ROOK:
		return "rook"
	case BISHOP:
		return "bishop"
	case KNIGHT:
		return "knight"
	case KING:
		return "king"
	case QUEEN:
		return "queen"
	}
	return "invalid"
}

func (p PieceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

type Color int

const (
	NO_COLOR = iota
	BLACK
	WHITE
)

func (c Color) String() string {
	switch c {
	case NO_COLOR:
		return "none"
	case WHITE:
		return "white"
	case BLACK:
		return "black"
	}
	return "invalid"
}

func (c Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
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

func (m MoveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m MoveType) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

type Move struct {
	Source      Position `json:"source"`
	Destination Position `json:"destination"`
	MoveType    MoveType `json:"type"`
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
	Rank int `json:"rank"`
	File int `json:"file"`
}

type Bounds struct {
	Rank int `json:"rank"`
	File int `json:"file"`
}

func (b Bounds) IsInboundsPosition(position Position) bool {
	return position.Rank >= 0 &&
		position.File >= 0 &&
		position.Rank < b.Rank &&
		position.File < b.File
}

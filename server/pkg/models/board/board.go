package board

import (
	"encoding/json"
	"errors"
	"strings"
)

type GameboardType string

const (
	GameboardTypeDefault GameboardType = ""
	GameboardTypeClassic GameboardType = "classic"
)

type GameboardState = map[int]map[int]*Piece

func NewGameboardState(bounds Bounds, pieceLocations GameboardState) GameboardState {
	state := GameboardState{}
	for rank := 0; rank < bounds.RankCount; rank += 1 {
		state[rank] = map[int]*Piece{}
		for file := 0; file < bounds.FileCount; file += 1 {
			state[rank][file] = nil
		}
	}

	for rank := range pieceLocations {
		for file := range pieceLocations[rank] {
			state[rank][file] = pieceLocations[rank][file]
		}
	}

	return state
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
	JUMP
	JUMP_CAPTURE
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
	case JUMP:
		return "jump"
	case JUMP_CAPTURE:
		return "jump_capture"
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

func (t *MoveType) UnmarshalJSON(data []byte) error {
	switch strings.Trim(string(data), "\"") {
	case NORMAL.String():
		*t = NORMAL
		return nil
	case CAPTURE.String():
		*t = CAPTURE
		return nil
	case JUMP.String():
		*t = JUMP
		return nil
	case JUMP_CAPTURE.String():
		*t = JUMP
		return nil
	case PAWN_DOUBLE_PUSH.String():
		*t = PAWN_DOUBLE_PUSH
		return nil
	case QUEENSIDE_CASTLE.String():
		*t = QUEENSIDE_CASTLE
		return nil
	case KINGSIDE_CASTLE.String():
		*t = KINGSIDE_CASTLE
		return nil
	case PROMOTION.String():
		*t = PROMOTION
		return nil
	case EN_PASSANT.String():
		*t = EN_PASSANT
		return nil
	}
	return errors.New("invalid string value for MoveType")
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
	MoveType    MoveType `json:"move_type"`
}

type MoveMap = map[MoveType][]Position

func NewMoveMap() MoveMap {
	return map[MoveType][]Position{
		NORMAL:           make([]Position, 0),
		CAPTURE:          make([]Position, 0),
		JUMP:             make([]Position, 0),
		JUMP_CAPTURE:     make([]Position, 0),
		PAWN_DOUBLE_PUSH: make([]Position, 0),
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
	RankCount int `json:"rank"`
	FileCount int `json:"file"`
}

func (b Bounds) IsInboundsPosition(position Position) bool {
	return position.Rank >= 0 &&
		position.File >= 0 &&
		position.Rank < b.RankCount &&
		position.File < b.FileCount
}

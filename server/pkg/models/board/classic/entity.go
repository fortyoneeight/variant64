package classic

import (
	"github.com/variant64/server/pkg/models/board"
)

type piece interface {
	GetType() board.PieceType
	GetColor() board.Color
	GetMoves(source board.Position) board.MoveMap
}

type ClassicBoard struct {
	board.Bounds
	whiteAllowedKingsideCastle  bool
	whiteAllowedQueensideCastle bool
	blackAllowedKingsideCastle  bool
	blackAllowedQueensideCastle bool
	locations                   map[int]map[int]piece
}

// New creates a new ClassicBoard and returns it.
func New() *ClassicBoard {
	locations := map[int]map[int]piece{}
	for rank := 0; rank < 8; rank += 1 {
		locations[rank] = map[int]piece{}

		for file := 0; file < 8; file += 1 {
			locations[rank][file] = nil
		}
	}

	bounds := board.Bounds{Rank: 8, File: 8}

	pieces := []struct {
		Rank  int
		File  int
		piece piece
	}{
		// white rooks
		{Rank: 0, File: 0, piece: &board.Rook{Bounds: bounds, Color: board.WHITE}},
		{Rank: 0, File: 7, piece: &board.Rook{Bounds: bounds, Color: board.WHITE}},
		// white bishops
		{Rank: 0, File: 1, piece: &board.Bishop{Bounds: bounds, Color: board.WHITE}},
		{Rank: 0, File: 6, piece: &board.Bishop{Bounds: bounds, Color: board.WHITE}},
		// white kights
		{Rank: 0, File: 2, piece: &board.Knight{Color: board.WHITE}},
		{Rank: 0, File: 5, piece: &board.Knight{Color: board.WHITE}},
		// white queen
		{Rank: 0, File: 3, piece: &board.Queen{Bounds: bounds, Color: board.WHITE}},
		// white king
		{Rank: 0, File: 4, piece: &board.King{Bounds: bounds, Color: board.WHITE}},
		// white pawns
		{Rank: 1, File: 0, piece: &board.Pawn{Color: board.WHITE}},
		{Rank: 1, File: 1, piece: &board.Pawn{Color: board.WHITE}},
		{Rank: 1, File: 2, piece: &board.Pawn{Color: board.WHITE}},
		{Rank: 1, File: 3, piece: &board.Pawn{Color: board.WHITE}},
		{Rank: 1, File: 4, piece: &board.Pawn{Color: board.WHITE}},
		{Rank: 1, File: 5, piece: &board.Pawn{Color: board.WHITE}},
		{Rank: 1, File: 6, piece: &board.Pawn{Color: board.WHITE}},
		{Rank: 1, File: 7, piece: &board.Pawn{Color: board.WHITE}},
		// black pawns
		{Rank: 6, File: 0, piece: &board.Pawn{Color: board.BLACK}},
		{Rank: 6, File: 1, piece: &board.Pawn{Color: board.BLACK}},
		{Rank: 6, File: 2, piece: &board.Pawn{Color: board.BLACK}},
		{Rank: 6, File: 3, piece: &board.Pawn{Color: board.BLACK}},
		{Rank: 6, File: 4, piece: &board.Pawn{Color: board.BLACK}},
		{Rank: 6, File: 5, piece: &board.Pawn{Color: board.BLACK}},
		{Rank: 6, File: 6, piece: &board.Pawn{Color: board.BLACK}},
		{Rank: 6, File: 7, piece: &board.Pawn{Color: board.BLACK}},
		// black rooks
		{Rank: 7, File: 0, piece: &board.Rook{Bounds: bounds, Color: board.BLACK}},
		{Rank: 7, File: 7, piece: &board.Rook{Bounds: bounds, Color: board.BLACK}},
		// black bishops
		{Rank: 7, File: 1, piece: &board.Bishop{Bounds: bounds, Color: board.BLACK}},
		{Rank: 7, File: 6, piece: &board.Bishop{Bounds: bounds, Color: board.BLACK}},
		// black kights
		{Rank: 7, File: 2, piece: &board.Knight{Color: board.BLACK}},
		{Rank: 7, File: 5, piece: &board.Knight{Color: board.BLACK}},
		// black queen
		{Rank: 7, File: 3, piece: &board.Queen{Bounds: bounds, Color: board.BLACK}},
		// black king
		{Rank: 7, File: 4, piece: &board.King{Bounds: bounds, Color: board.BLACK}},
	}
	for _, p := range pieces {
		locations[p.Rank][p.File] = p.piece
	}

	board := &ClassicBoard{
		Bounds:                      board.Bounds{Rank: 8, File: 8},
		whiteAllowedKingsideCastle:  true,
		whiteAllowedQueensideCastle: true,
		blackAllowedKingsideCastle:  true,
		blackAllowedQueensideCastle: true,
		locations:                   locations,
	}

	return board
}

// GetAllMoves returns a map, by rank and file of all legal moves
// a piece at that position can take deliniated by MoveType.
func (b *ClassicBoard) GetAllMoves() map[int]map[int]board.MoveMap {
	allMoves := map[int]map[int]board.MoveMap{}

	for rank, files := range b.locations {
		allMoves[rank] = map[int]board.MoveMap{}
		for file, piece := range files {
			if piece != nil {
				allMoves[rank][file] = b.getLegalMoves(board.Position{Rank: rank, File: file})
			}
		}
	}

	return allMoves
}

// getPiece returns the piece at the provided Position.
func (b *ClassicBoard) getPiece(position board.Position) piece {
	return b.locations[position.Rank][position.File]
}

// getLegalMoves returns a LegalMovesMap for the piece at the provided position.
func (b *ClassicBoard) getLegalMoves(source board.Position) board.MoveMap {
	piece := b.getPiece(source)
	if piece == nil {
		return board.NewMoveMap()
	}

	moves := piece.GetMoves(source)

	b.filterNormalMoves(moves, source, piece)
	b.filterCaptureMoves(moves, source, piece)
	b.filterPawnDoublePushMoves(moves, source, piece)
	b.filterRayMoves(moves, source, piece)
	b.filterKingsideCastleMoves(moves, source, piece)
	b.filterQueensideCastleMoves(moves, source, piece)

	return moves
}

// filterNormalMoves filters board.NORMAL moves for a provided position.
func (b *ClassicBoard) filterNormalMoves(moveMap board.MoveMap, source board.Position, piece piece) {
	allowedMoves := []board.Position{}

	for _, position := range moveMap[board.NORMAL] {
		if b.IsInboundsPosition(position) {
			allowedMoves = append(allowedMoves, position)
		}
	}

	moveMap[board.NORMAL] = allowedMoves
}

// filterCaptureMoves filters board.CAPTURE moves for a provided position.
func (b *ClassicBoard) filterCaptureMoves(moveMap board.MoveMap, source board.Position, piece piece) {
	allowedMoves := []board.Position{}

	for _, position := range moveMap[board.CAPTURE] {
		capturedPiece := b.getPiece(position)
		if b.IsInboundsPosition(position) && capturedPiece != nil && capturedPiece.GetColor() != piece.GetColor() {
			allowedMoves = append(allowedMoves, position)
		}
	}

	moveMap[board.CAPTURE] = allowedMoves
}

// filterPawnDoublePushMoves filters board.PAWN_DOUBLE_PUSH moves for a provided position.
func (b *ClassicBoard) filterPawnDoublePushMoves(moveMap board.MoveMap, source board.Position, piece piece) {
	allowedMoves := []board.Position{}

	for _, position := range moveMap[board.PAWN_DOUBLE_PUSH] {
		if b.IsInboundsPosition(position) {
			if piece.GetColor() == board.WHITE && source.Rank == 1 && b.getPiece(board.Position{Rank: 2, File: source.File}) == nil {
				allowedMoves = append(allowedMoves, position)
			} else if piece.GetColor() == board.BLACK && source.Rank == 6 && b.getPiece(board.Position{Rank: 5, File: source.File}) == nil {
				allowedMoves = append(allowedMoves, position)
			}
		}
	}

	moveMap[board.PAWN_DOUBLE_PUSH] = allowedMoves
}

// filterKingsideCastleMoves filters board.KINGSIDE_CASTLE moves for a provided position.
func (b *ClassicBoard) filterKingsideCastleMoves(moveMap board.MoveMap, source board.Position, piece piece) {
	allowedMoves := []board.Position{}

	allowed := false
	requiredEmpty := []board.Position{}
	if piece.GetColor() == board.WHITE && b.whiteAllowedKingsideCastle {
		allowed = true
		requiredEmpty = []board.Position{
			{Rank: 0, File: 5},
			{Rank: 0, File: 6},
		}
	} else if piece.GetColor() == board.BLACK && b.blackAllowedKingsideCastle {
		allowed = true
		requiredEmpty = []board.Position{
			{Rank: 7, File: 5},
			{Rank: 7, File: 6},
		}
	}
	if allowed {
		for _, position := range moveMap[board.KINGSIDE_CASTLE] {
			allEmpty := true
			for _, r := range requiredEmpty {
				if b.getPiece(r) != nil {
					allEmpty = false
				}
			}
			if allEmpty {
				allowedMoves = append(allowedMoves, position)
			}
		}
	}

	moveMap[board.KINGSIDE_CASTLE] = allowedMoves
}

// filterQueensideCastleMoves filters board.QUEENSIDE_CASTLE moves for a provided position.
func (b *ClassicBoard) filterQueensideCastleMoves(moveMap board.MoveMap, source board.Position, piece piece) {
	allowedMoves := []board.Position{}

	allowed := false
	requiredEmpty := []board.Position{}
	if piece.GetColor() == board.WHITE && b.whiteAllowedQueensideCastle {
		allowed = true
		requiredEmpty = []board.Position{
			{Rank: 0, File: 1},
			{Rank: 0, File: 2},
			{Rank: 0, File: 3},
		}
	} else if piece.GetColor() == board.BLACK && b.blackAllowedQueensideCastle {
		allowed = true
		requiredEmpty = []board.Position{
			{Rank: 7, File: 1},
			{Rank: 7, File: 2},
			{Rank: 7, File: 3},
		}
	}
	if allowed {
		for _, position := range moveMap[board.QUEENSIDE_CASTLE] {
			allEmpty := true
			for _, r := range requiredEmpty {
				if b.getPiece(r) != nil {
					allEmpty = false
				}
			}
			if allEmpty {
				allowedMoves = append(allowedMoves, position)
			}
		}
	}

	moveMap[board.QUEENSIDE_CASTLE] = allowedMoves
}

// filterRayMoves filter board.RAY moves for a provided position.
func (b *ClassicBoard) filterRayMoves(moveMap board.MoveMap, source board.Position, piece piece) {
	moves := board.NewMoveMap()

	for _, position := range moveMap[board.RAY] {
		direction := board.GetDirection(source, position)
		ray := board.GenerateRay(source, direction, b.Bounds)
		for _, nextPosition := range ray {
			nextPositionPiece := b.getPiece(nextPosition)
			switch {
			case nextPositionPiece == nil:
				moves[board.NORMAL] = append(moves[board.NORMAL], nextPosition)
			case nextPositionPiece.GetColor() != piece.GetColor():
				moves[board.CAPTURE] = append(moves[board.CAPTURE], nextPosition)
			default:
				break
			}
		}
	}

	board.JoinMoveMaps(moveMap, moves)
}

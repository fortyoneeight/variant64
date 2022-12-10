package classic

import (
	"errors"
	"fmt"

	"github.com/variant64/server/pkg/models/board"
)

type piece interface {
	GetType() board.PieceType
	GetColor() board.Color
	GetMoves(source board.Position) board.MoveMap
}

type locationMap = map[int]map[int]piece

func NewPieceLocations(bounds board.Bounds, pieceLocations locationMap) locationMap {
	boardPieceLocations := locationMap{}
	for rank := 0; rank < bounds.Rank; rank += 1 {
		boardPieceLocations[rank] = map[int]piece{}
		for file := 0; file < bounds.File; file += 1 {
			boardPieceLocations[rank][file] = nil
		}
	}

	for rank := range pieceLocations {
		for file := range pieceLocations[rank] {
			boardPieceLocations[rank][file] = pieceLocations[rank][file]
		}
	}

	return boardPieceLocations
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
	bounds := board.Bounds{Rank: 8, File: 8}

	locations := map[int]map[int]piece{}
	for rank := 0; rank < bounds.Rank; rank += 1 {
		locations[rank] = map[int]piece{}

		for file := 0; file < bounds.File; file += 1 {
			locations[rank][file] = nil
		}
	}

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
		Bounds:                      bounds,
		whiteAllowedKingsideCastle:  true,
		whiteAllowedQueensideCastle: true,
		blackAllowedKingsideCastle:  true,
		blackAllowedQueensideCastle: true,
		locations:                   locations,
	}

	return board
}

// GetAllMoves returns a map, by rank and file of all legal moves
// a piece at that position can take delineated by MoveType.
func (b *ClassicBoard) GetAllMoves() map[int]map[int]board.MoveMap {
	allMoves := map[int]map[int]board.MoveMap{}

	for rank, files := range b.locations {
		allMoves[rank] = map[int]board.MoveMap{}
		for file, piece := range files {
			if piece != nil {
				allMoves[rank][file] = b.getMovesAtPosition(board.Position{Rank: rank, File: file})
			}
		}
	}

	return allMoves
}

// HandleMove handles a Move submitted by the client.
func (b *ClassicBoard) HandleMove(move board.Move) error {
	// Check if there is a piece at the source position.
	sourcePiece := b.getPiece(move.Source)
	if sourcePiece == nil {
		return errPieceNotFound
	}

	// Verify move is legal.
	err := b.isMoveAllowed(move, sourcePiece)
	if err != nil {
		return errMoveNotAllowed
	}

	// Update the board state.
	var moveErr error
	switch move.MoveType {
	case board.NORMAL:
		moveErr = b.applySingleMoveToLocations(move, sourcePiece)
	case board.CAPTURE:
		moveErr = b.applySingleMoveToLocations(move, sourcePiece)
	case board.PAWN_DOUBLE_PUSH:
		moveErr = b.applySingleMoveToLocations(move, sourcePiece)
	case board.KINGSIDE_CASTLE:
		moveErr = b.applyKingsideCastleToLocations(move, sourcePiece)
	case board.QUEENSIDE_CASTLE:
		moveErr = b.applyQueensideCastleToLocations(move, sourcePiece)
	default:
		moveErr = errMoveNotAllowed
	}
	if moveErr != nil {
		return moveErr
	}

	// Update the castle flags if necessary.
	return b.updateCastlingFlag(move, sourcePiece)
}

// isMoveAllowed checks if a move is legal.
func (b *ClassicBoard) isMoveAllowed(move board.Move, sourcePiece piece) error {
	// Check if the source and destination positions are within the board's boundaries.
	if !b.IsInboundsPosition(move.Source) || !b.IsInboundsPosition(move.Destination) {
		return errors.New("invalid move: position out of bounds")
	}

	// Check if the piece at the source position can make the move specified by the move type.
	legalMoves := b.getMovesAtPosition(move.Source)
	legalMovesByType, ok := legalMoves[move.MoveType]
	if !ok {
		return fmt.Errorf("invalid move: %s cannot be made with %d at source position", move.MoveType, sourcePiece.GetType())
	}

	// Check if the destination position is included in the list of legal destination positions for the move type.
	var legalMove bool
	for _, destination := range legalMovesByType {
		if destination == move.Destination {
			legalMove = true
			break
		}
	}
	if !legalMove {
		return errors.New("invalid move: destination position is not a valid destination for the specified move type")
	}

	return nil
}

// applySingleMoveToLocations handles moving a single piece from one square to another.
func (b *ClassicBoard) applySingleMoveToLocations(move board.Move, sourcePiece piece) error {
	b.locations[move.Destination.Rank][move.Destination.File] = sourcePiece
	b.locations[move.Source.Rank][move.Source.File] = nil
	return nil
}

// applyKingsideCastleToLocations handles a KINGSIDE_CASTLE move.
func (b *ClassicBoard) applyKingsideCastleToLocations(move board.Move, sourcePiece piece) error {
	// Check that castling is allowed.
	if sourcePiece.GetColor() == board.WHITE && !b.whiteAllowedKingsideCastle {
		return errNotAllowedToCastle(sourcePiece.GetColor())
	} else if sourcePiece.GetColor() == board.BLACK && !b.blackAllowedKingsideCastle {
		return errNotAllowedToCastle(sourcePiece.GetColor())
	}

	// Handle the king movement.
	b.locations[move.Destination.Rank][move.Destination.File] = sourcePiece
	b.locations[move.Source.Rank][move.Source.File] = nil

	// Handle the rook movement.
	switch sourcePiece.GetColor() {
	case board.WHITE:
		b.locations[0][5] = b.locations[0][7]
		b.locations[0][7] = nil
	case board.BLACK:
		b.locations[7][5] = b.locations[7][7]
		b.locations[7][7] = nil
	default:
		return errNotAllowedToCastle(sourcePiece.GetColor())
	}

	return nil
}

// applyQueensideCastleToLocations handles a QUEENSIDE_CASTLE move.
func (b *ClassicBoard) applyQueensideCastleToLocations(move board.Move, sourcePiece piece) error {
	// Check that castling is allowed.
	if sourcePiece.GetColor() == board.WHITE && !b.whiteAllowedQueensideCastle {
		return errNotAllowedToCastle(sourcePiece.GetColor())
	} else if sourcePiece.GetColor() == board.BLACK && !b.blackAllowedQueensideCastle {
		return errNotAllowedToCastle(sourcePiece.GetColor())
	}

	// Handle the king movement.
	b.locations[move.Destination.Rank][move.Destination.File] = sourcePiece
	b.locations[move.Source.Rank][move.Source.File] = nil

	// Handle the rook movement.
	switch sourcePiece.GetColor() {
	case board.WHITE:
		b.locations[0][3] = b.locations[0][0]
		b.locations[0][0] = nil
	case board.BLACK:
		b.locations[7][3] = b.locations[7][0]
		b.locations[7][0] = nil
	default:
		return errInvalidColor(sourcePiece.GetColor())
	}

	return nil
}

// updateCastlingFlag updates the castling flag for a player if neccessary.
func (b *ClassicBoard) updateCastlingFlag(move board.Move, sourcePiece piece) error {
	if sourcePiece.GetType() == board.KING {
		switch sourcePiece.GetColor() {
		case board.BLACK:
			b.blackAllowedKingsideCastle = false
			b.blackAllowedQueensideCastle = false
		case board.WHITE:
			b.whiteAllowedKingsideCastle = false
			b.whiteAllowedQueensideCastle = false
		default:
			return errInvalidColor(sourcePiece.GetColor())
		}
	} else if sourcePiece.GetType() == board.ROOK {
		switch sourcePiece.GetColor() {
		case board.WHITE:
			if move.Source.Rank == 0 && move.Source.File == 0 {
				b.whiteAllowedQueensideCastle = false
			} else if move.Source.Rank == 0 && move.Source.File == 7 {
				b.whiteAllowedKingsideCastle = false
			}
		case board.BLACK:
			if move.Source.Rank == 7 && move.Source.File == 0 {
				b.blackAllowedQueensideCastle = false
			} else if move.Source.Rank == 7 && move.Source.File == 7 {
				b.blackAllowedKingsideCastle = false
			}
		default:
			return errInvalidColor(sourcePiece.GetColor())
		}
	}

	return nil
}

// getPiece returns the piece at the provided Position.
func (b *ClassicBoard) getPiece(position board.Position) piece {
	return b.locations[position.Rank][position.File]
}

// getMovesAtPosition returns a board.MoveMap for the piece at the provided position.
func (b *ClassicBoard) getMovesAtPosition(source board.Position) board.MoveMap {
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

package classic

import (
	"errors"
	"fmt"

	"github.com/variant64/server/pkg/models/board"
)

type ClassicBoard struct {
	board.Bounds
	board.MoveApplicator
	*board.MoveFilter
	*board.CastlingState
	gameboardState board.GameboardState
}

// New creates a new ClassicBoard and returns it.
func New() *ClassicBoard {
	bounds := board.Bounds{RankCount: 8, FileCount: 8}
	castlingState := board.NewCastlingState()
	moveHandler := board.NewMoveApplicator(
		&board.SinglePieceMoveApplicator{},
		&board.KingsideCastleMoveApplicator{},
		&board.QueensideCastleMoveApplicator{},
	)
	moveFilter := board.NewMoveFilter(
		&board.FilterOutOfBounds{Bounds: bounds},
		&board.FilterPieceCollision{},
		&board.FilterFriendlyCapture{},
		&board.FilterInvalidPawnDoublePush{},
		&board.FilterIllegalKingsideCastle{
			CastlingState: castlingState,
		},
		&board.FilterIllegalQueensideCastle{
			CastlingState: castlingState,
		},
	)

	gameboardState := map[int]map[int]*board.Piece{}
	for rank := 0; rank < bounds.RankCount; rank += 1 {
		gameboardState[rank] = map[int]*board.Piece{}
		for file := 0; file < bounds.FileCount; file += 1 {
			gameboardState[rank][file] = nil
		}
	}

	pieces := []struct {
		Rank  int
		File  int
		piece *board.Piece
	}{
		// white rooks
		{Rank: 0, File: 0, piece: board.NewRook(board.WHITE, bounds)},
		{Rank: 0, File: 7, piece: board.NewRook(board.WHITE, bounds)},
		// white kights
		{Rank: 0, File: 1, piece: board.NewKnight(board.WHITE)},
		{Rank: 0, File: 6, piece: board.NewKnight(board.WHITE)},
		// white bishops
		{Rank: 0, File: 2, piece: board.NewBishop(board.WHITE, bounds)},
		{Rank: 0, File: 5, piece: board.NewBishop(board.WHITE, bounds)},
		// white queen
		{Rank: 0, File: 3, piece: board.NewQueen(board.WHITE, bounds)},
		// white king
		{Rank: 0, File: 4, piece: board.NewKing(board.WHITE)},
		// white pawns
		{Rank: 1, File: 0, piece: board.NewPawn(board.WHITE)},
		{Rank: 1, File: 1, piece: board.NewPawn(board.WHITE)},
		{Rank: 1, File: 2, piece: board.NewPawn(board.WHITE)},
		{Rank: 1, File: 3, piece: board.NewPawn(board.WHITE)},
		{Rank: 1, File: 4, piece: board.NewPawn(board.WHITE)},
		{Rank: 1, File: 5, piece: board.NewPawn(board.WHITE)},
		{Rank: 1, File: 6, piece: board.NewPawn(board.WHITE)},
		{Rank: 1, File: 7, piece: board.NewPawn(board.WHITE)},
		// black pawns
		{Rank: 6, File: 0, piece: board.NewPawn(board.BLACK)},
		{Rank: 6, File: 1, piece: board.NewPawn(board.BLACK)},
		{Rank: 6, File: 2, piece: board.NewPawn(board.BLACK)},
		{Rank: 6, File: 3, piece: board.NewPawn(board.BLACK)},
		{Rank: 6, File: 4, piece: board.NewPawn(board.BLACK)},
		{Rank: 6, File: 5, piece: board.NewPawn(board.BLACK)},
		{Rank: 6, File: 6, piece: board.NewPawn(board.BLACK)},
		{Rank: 6, File: 7, piece: board.NewPawn(board.BLACK)},
		// black rooks
		{Rank: 7, File: 0, piece: board.NewRook(board.BLACK, bounds)},
		{Rank: 7, File: 7, piece: board.NewRook(board.BLACK, bounds)},
		// black kights
		{Rank: 7, File: 1, piece: board.NewKnight(board.BLACK)},
		{Rank: 7, File: 6, piece: board.NewKnight(board.BLACK)},
		// black bishops
		{Rank: 7, File: 2, piece: board.NewBishop(board.BLACK, bounds)},
		{Rank: 7, File: 5, piece: board.NewBishop(board.BLACK, bounds)},
		// black queen
		{Rank: 7, File: 3, piece: board.NewQueen(board.BLACK, bounds)},
		// black king
		{Rank: 7, File: 4, piece: board.NewKing(board.BLACK)},
	}
	for _, p := range pieces {
		gameboardState[p.Rank][p.File] = p.piece
	}

	board := &ClassicBoard{
		Bounds:         bounds,
		MoveApplicator: moveHandler,
		MoveFilter:     moveFilter,
		CastlingState:  castlingState,
		gameboardState: gameboardState,
	}
	board.updateAvailableMoves()

	return board
}

// GetState returns a board.GameboardState for the ClassicBoard.
func (b *ClassicBoard) GetState() board.GameboardState {
	return b.gameboardState
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
	moveErr := b.ApplyMove(move, b.gameboardState)
	if moveErr != nil {
		return moveErr
	}

	// Update the castle flags if necessary.
	err = b.updateCastlingFlag(move, sourcePiece)
	if err != nil {
		return err
	}

	// Update the available moves for each piece.
	b.updateAvailableMoves()

	return nil
}

// isMoveAllowed checks if a move is legal.
func (b *ClassicBoard) isMoveAllowed(move board.Move, sourcePiece *board.Piece) error {
	// Check if the source and destination positions are within the board's boundaries.
	if !b.IsInboundsPosition(move.Source) || !b.IsInboundsPosition(move.Destination) {
		return errors.New("invalid move: position out of bounds")
	}

	// Check if the piece at the source position can make the move specified by the move type.
	legalMovesByType, ok := sourcePiece.AvailableMoves[move.MoveType]
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

// updateAvailableMoves sets the available moves for each piece in the game.
func (b *ClassicBoard) updateAvailableMoves() {
	for rank, files := range b.gameboardState {
		for file, piece := range files {
			if piece != nil {
				moves := b.getMovesAtPosition(board.Position{Rank: rank, File: file})
				piece.AvailableMoves = moves
			}
		}
	}
}

// updateCastlingFlag updates the castling flag for a player if neccessary.
func (b *ClassicBoard) updateCastlingFlag(move board.Move, sourcePiece *board.Piece) error {
	if sourcePiece.GetType() == board.KING {
		switch sourcePiece.GetColor() {
		case board.BLACK:
			b.CastlingState.Disallow(board.KINGSIDE_CASTLE, board.BLACK)
			b.CastlingState.Disallow(board.QUEENSIDE_CASTLE, board.BLACK)
		case board.WHITE:
			b.CastlingState.Disallow(board.KINGSIDE_CASTLE, board.WHITE)
			b.CastlingState.Disallow(board.QUEENSIDE_CASTLE, board.WHITE)
		default:
			return errInvalidColor(sourcePiece.GetColor())
		}
	} else if sourcePiece.GetType() == board.ROOK {
		switch sourcePiece.GetColor() {
		case board.WHITE:
			if move.Source.Rank == 0 && move.Source.File == 0 {
				b.CastlingState.Disallow(board.KINGSIDE_CASTLE, board.BLACK)
			} else if move.Source.Rank == 0 && move.Source.File == 7 {
				b.CastlingState.Disallow(board.KINGSIDE_CASTLE, board.WHITE)
			}
		case board.BLACK:
			if move.Source.Rank == 7 && move.Source.File == 0 {
				b.CastlingState.Disallow(board.QUEENSIDE_CASTLE, board.BLACK)
			} else if move.Source.Rank == 7 && move.Source.File == 7 {
				b.CastlingState.Disallow(board.KINGSIDE_CASTLE, board.BLACK)
			}
		default:
			return errInvalidColor(sourcePiece.GetColor())
		}
	}

	return nil
}

// getPiece returns the piece at the provided Position.
func (b *ClassicBoard) getPiece(position board.Position) *board.Piece {
	return b.gameboardState[position.Rank][position.File]
}

// getMovesAtPosition returns a board.MoveMap for the piece at the provided position.
func (b *ClassicBoard) getMovesAtPosition(source board.Position) board.MoveMap {
	piece := b.getPiece(source)
	if piece == nil {
		return board.NewMoveMap()
	}

	availableMoves := board.NewMoveMap()
	moves := piece.GetMoves(source)
	for moveType, moveListByType := range moves {
		availableMovesByType := []board.Position{}
		for _, destination := range moveListByType {
			move := board.Move{Source: source, Destination: destination, MoveType: moveType}
			if b.IsLegalMove(move, b.gameboardState) {
				availableMovesByType = append(availableMovesByType, destination)
			}
		}
		availableMoves[moveType] = availableMovesByType
	}

	return availableMoves
}

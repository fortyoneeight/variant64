package board

type moveFilter interface {
	IsLegalMove(move Move, state GameboardState) bool
}

// MoveFilter bundles multiple moveFilters into a single struct.
type MoveFilter struct {
	moveFilters []moveFilter
}

func NewMoveFilter(moveFilters ...moveFilter) *MoveFilter {
	return &MoveFilter{
		moveFilters: moveFilters,
	}
}

func (m *MoveFilter) IsLegalMove(move Move, state GameboardState) bool {
	for _, filter := range m.moveFilters {
		if !filter.IsLegalMove(move, state) {
			return false
		}
	}
	return true
}

// FilterOutOfBounds disallows piece to move out of bounds.
type FilterOutOfBounds struct {
	Bounds
}

func (f *FilterOutOfBounds) IsLegalMove(move Move, state GameboardState) bool {
	return f.IsInboundsPosition(move.Source) && f.IsInboundsPosition(move.Destination)
}

// FilterPieceCollision disallows pieces to move through other pieces.
type FilterPieceCollision struct{}

func (f *FilterPieceCollision) IsLegalMove(move Move, state GameboardState) bool {
	switch move.MoveType {
	case NORMAL, CAPTURE, PAWN_DOUBLE_PUSH:
		return f.isEmptyRay(move.MoveType, move.Source, move.Destination, state)
	default:
		return true
	}
}

func (f *FilterPieceCollision) isEmptyRay(moveType MoveType, source, destination Position, state GameboardState) bool {
	direction := GetDirection(source, destination)
	next := source
	for {
		next = StepInDirection(next, direction)
		if next == destination {
			return state[next.Rank][next.File] == nil || moveType == CAPTURE
		} else if state[next.Rank][next.File] != nil {
			return false
		}
	}
}

// FilterFriendlyCapture disallows pieces to capture a piece with the same COLOR.
type FilterFriendlyCapture struct{}

func (f *FilterFriendlyCapture) IsLegalMove(move Move, state GameboardState) bool {
	switch move.MoveType {
	case CAPTURE, JUMP_CAPTURE:
		capturingPiece := state[move.Source.Rank][move.Source.File]
		capturedPiece := state[move.Destination.Rank][move.Destination.File]
		if capturingPiece == nil || capturedPiece == nil || capturingPiece.Color == capturedPiece.Color {
			return false
		}
		return true
	default:
		return true
	}
}

// FilterInvalidPawnDoublePush disallows pawns to double outside of their initial position.
type FilterInvalidPawnDoublePush struct{}

func (f *FilterInvalidPawnDoublePush) IsLegalMove(move Move, state GameboardState) bool {
	switch move.MoveType {
	case PAWN_DOUBLE_PUSH:
		piece := state[move.Source.Rank][move.Source.File]
		if piece == nil {
			return false
		}
		switch piece.Color {
		case WHITE:
			return move.Source.Rank == 1
		case BLACK:
			return move.Source.Rank == 6
		default:
			return false
		}
	default:
		return true
	}
}

// FilterIllegalKingsideCastle disallows illegal kingside castles.
type FilterIllegalKingsideCastle struct {
	*CastlingState
}

func (f *FilterIllegalKingsideCastle) IsLegalMove(move Move, state GameboardState) bool {
	switch move.MoveType {
	case KINGSIDE_CASTLE:
		piece := state[move.Source.Rank][move.Source.File]
		if piece == nil {
			return false
		}
		switch piece.Color {
		case WHITE:
			if !f.IsAllowed(KINGSIDE_CASTLE, WHITE) {
				return false
			}
			return state[0][5] == nil && state[0][6] == nil
		case BLACK:
			if !f.IsAllowed(KINGSIDE_CASTLE, BLACK) {
				return false
			}
			return state[7][5] == nil && state[7][6] == nil
		default:
			return false
		}
	default:
		return true
	}
}

// FilterIllegalQueensideCastle disallows illegal queenside castles.
type FilterIllegalQueensideCastle struct {
	*CastlingState
}

func (f *FilterIllegalQueensideCastle) IsLegalMove(move Move, state GameboardState) bool {
	switch move.MoveType {
	case QUEENSIDE_CASTLE:
		piece := state[move.Source.Rank][move.Source.File]
		if piece == nil {
			return false
		}
		switch piece.Color {
		case WHITE:
			if !f.IsAllowed(QUEENSIDE_CASTLE, WHITE) {
				return false
			}
			return state[0][1] == nil && state[0][2] == nil && state[0][3] == nil
		case BLACK:
			if !f.IsAllowed(QUEENSIDE_CASTLE, BLACK) {
				return false
			}
			return state[7][1] == nil && state[7][2] == nil && state[7][3] == nil
		default:
			return false
		}
	default:
		return true
	}
}

// FilterIllegalPromotion disallows illegal promotions.
type FilterIllegalPromotion struct {
	bounds Bounds
}

func (f *FilterIllegalPromotion) IsLegalMove(move Move, state GameboardState) bool {
	switch move.MoveType {
	case PROMOTION:
		piece := state[move.Source.Rank][move.Source.File]
		if piece == nil {
			return false
		}
		if piece.PieceType != PAWN {
			return false
		}

		switch piece.Color {
		case WHITE:
			return move.Source.Rank == 6 &&
				move.Destination.Rank == 7 &&
				move.Source.File == move.Destination.File
		case BLACK:
			return move.Source.Rank == 1 &&
				move.Destination.Rank == 0 &&
				move.Source.File == move.Destination.File
		default:
			return false
		}
	default:
		return true
	}
}

// FilterIllegalPromotionCapture disallows illegal promotion captures.
type FilterIllegalPromotionCapture struct {
	bounds Bounds
}

func (f *FilterIllegalPromotionCapture) IsLegalMove(move Move, state GameboardState) bool {
	switch move.MoveType {
	case PROMOTION_CAPTURE:
		piece := state[move.Source.Rank][move.Source.File]
		if piece == nil {
			return false
		}
		if piece.PieceType != PAWN {
			return false
		}

		capturedPiece := state[move.Destination.Rank][move.Destination.File]
		if capturedPiece == nil {
			return false
		}
		if capturedPiece.Color == piece.Color {
			return false
		}

		// check for diagonal movement
		if move.Source.File != move.Destination.File-1 &&
			move.Source.File != move.Destination.File+1 {
			return false
		}

		switch piece.Color {
		case WHITE:
			return move.Source.Rank == f.bounds.RankCount-2 &&
				move.Destination.Rank == f.bounds.RankCount-1
		case BLACK:
			return move.Source.Rank == 1 && move.Destination.Rank == 0
		default:
			return false
		}
	default:
		return true
	}
}

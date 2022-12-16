package board

type moveApplicator interface {
	GetTypesToHandle() map[MoveType]bool
	ApplyMove(move Move, state GameboardState) error
}

// MoveApplicator bundles multiple moveApplicators into a single struct.
type MoveApplicator struct {
	moveApplicatorMap map[MoveType]moveApplicator
}

// NewMoveApplicator returns a MoveApplicator composed of all provided moveApplicator.
func NewMoveApplicator(moveHandlers ...moveApplicator) *MoveApplicator {
	newMoveApplicator := &MoveApplicator{
		moveApplicatorMap: make(map[MoveType]moveApplicator),
	}

	for _, h := range moveHandlers {
		for mt := range h.GetTypesToHandle() {
			newMoveApplicator.moveApplicatorMap[mt] = h
		}
	}

	return newMoveApplicator
}

func (h *MoveApplicator) ApplyMove(move Move, state GameboardState) error {
	handler, ok := h.moveApplicatorMap[move.MoveType]
	if !ok {
		return errCannotHandleMoveType(move.MoveType)
	}

	return handler.ApplyMove(move, state)
}

// SinglePieceMoveApplicator applies a single piece movement to a StateMap.
type SinglePieceMoveApplicator struct{}

func (h *SinglePieceMoveApplicator) GetTypesToHandle() map[MoveType]bool {
	return map[MoveType]bool{
		NORMAL:           true,
		CAPTURE:          true,
		JUMP:             true,
		JUMP_CAPTURE:     true,
		PAWN_DOUBLE_PUSH: true,
	}
}

func (h *SinglePieceMoveApplicator) ApplyMove(move Move, state GameboardState) error {
	if _, ok := h.GetTypesToHandle()[move.MoveType]; !ok {
		return errCannotHandleMoveType(move.MoveType)
	}

	sourcePiece := state[move.Source.Rank][move.Source.File]
	if sourcePiece == nil {
		return errSourcePieceNotFound
	}

	state[move.Destination.Rank][move.Destination.File] = sourcePiece
	state[move.Source.Rank][move.Source.File] = nil

	return nil
}

// KingsideCastleMoveApplicator applies a kingside castle to a StateMap.
type KingsideCastleMoveApplicator struct{}

func (h *KingsideCastleMoveApplicator) GetTypesToHandle() map[MoveType]bool {
	return map[MoveType]bool{
		KINGSIDE_CASTLE: true,
	}
}

func (h *KingsideCastleMoveApplicator) ApplyMove(move Move, state GameboardState) error {
	if _, ok := h.GetTypesToHandle()[move.MoveType]; !ok {
		return errCannotHandleMoveType(move.MoveType)
	}

	// Check king present.
	kingPiece := state[move.Source.Rank][move.Source.File]
	if kingPiece == nil {
		return errSourcePieceNotFound
	}

	// Check rook present.
	switch kingPiece.GetColor() {
	case WHITE:
		rookPiece := state[0][7]
		if rookPiece == nil {
			return errSourcePieceNotFound
		}
	case BLACK:
		rookPiece := state[7][7]
		if rookPiece == nil {
			return errSourcePieceNotFound
		}
	}

	// Handle the king movement.
	state[move.Destination.Rank][move.Destination.File] = kingPiece
	state[move.Source.Rank][move.Source.File] = nil

	// Handle the rook movement.
	switch kingPiece.GetColor() {
	case WHITE:
		state[0][5] = state[0][7]
		state[0][7] = nil
	case BLACK:
		state[7][5] = state[7][7]
		state[7][7] = nil
	}

	return nil
}

// QueensideCastleMoveApplicator applies a queenside castle to a StateMap.
type QueensideCastleMoveApplicator struct{}

func (h *QueensideCastleMoveApplicator) GetTypesToHandle() map[MoveType]bool {
	return map[MoveType]bool{
		QUEENSIDE_CASTLE: true,
	}
}

func (h *QueensideCastleMoveApplicator) ApplyMove(move Move, state GameboardState) error {
	if _, ok := h.GetTypesToHandle()[move.MoveType]; !ok {
		return errCannotHandleMoveType(move.MoveType)
	}
	// Check king present.
	kingPiece := state[move.Source.Rank][move.Source.File]
	if kingPiece == nil {
		return errSourcePieceNotFound
	}

	// Check rook present.
	switch kingPiece.GetColor() {
	case WHITE:
		rookPiece := state[0][0]
		if rookPiece == nil {
			return errSourcePieceNotFound
		}
	case BLACK:
		rookPiece := state[7][0]
		if rookPiece == nil {
			return errSourcePieceNotFound
		}
	}

	// Handle the king movement.
	state[move.Destination.Rank][move.Destination.File] = kingPiece
	state[move.Source.Rank][move.Source.File] = nil

	// Handle the rook movement.
	switch kingPiece.GetColor() {
	case WHITE:
		state[0][3] = state[0][0]
		state[0][0] = nil
	case BLACK:
		state[7][3] = state[7][0]
		state[7][0] = nil
	}

	return nil
}

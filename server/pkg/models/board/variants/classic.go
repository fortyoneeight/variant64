package variants

import (
	"github.com/variant64/server/pkg/models/board"
)

// NewClassicBoard creates a new Board with classic rules and returns it.
func NewClassicBoard() *board.Board {
	bounds := board.Bounds{RankCount: 8, FileCount: 8}
	castlingState := board.NewDefaultCastlingState()
	turnState := &board.TurnState{
		Active:    board.WHITE,
		TurnOrder: []board.Color{board.BLACK, board.WHITE},
	}

	return board.Build(
		board.WithBounds(bounds),
		board.WithCastlingState(castlingState),
		board.WithMoveApplicator(
			board.NewMoveApplicator(
				&board.SinglePieceMoveApplicator{},
				&board.KingsideCastleMoveApplicator{},
				&board.QueensideCastleMoveApplicator{},
				&board.PromotionMoveApplicator{Bounds: bounds},
			),
		),
		board.WithMoveFilter(
			board.NewMoveFilter(
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
				&board.FilterIllegalPromotion{},
				&board.FilterIllegalPromotionCapture{},
			),
		),
		board.WithGameboardState(
			board.GameboardState{
				7: {
					0: board.NewRook(board.BLACK, bounds),
					1: board.NewKnight(board.BLACK),
					2: board.NewBishop(board.BLACK, bounds),
					3: board.NewQueen(board.BLACK, bounds),
					4: board.NewKing(board.BLACK),
					5: board.NewBishop(board.BLACK, bounds),
					6: board.NewKnight(board.BLACK),
					7: board.NewRook(board.BLACK, bounds),
				},
				6: {
					0: board.NewPawn(board.BLACK),
					1: board.NewPawn(board.BLACK),
					2: board.NewPawn(board.BLACK),
					3: board.NewPawn(board.BLACK),
					4: board.NewPawn(board.BLACK),
					5: board.NewPawn(board.BLACK),
					6: board.NewPawn(board.BLACK),
					7: board.NewPawn(board.BLACK),
				},
				1: {
					0: board.NewPawn(board.WHITE),
					1: board.NewPawn(board.WHITE),
					2: board.NewPawn(board.WHITE),
					3: board.NewPawn(board.WHITE),
					4: board.NewPawn(board.WHITE),
					5: board.NewPawn(board.WHITE),
					6: board.NewPawn(board.WHITE),
					7: board.NewPawn(board.WHITE),
				},
				0: {
					0: board.NewRook(board.WHITE, bounds),
					1: board.NewKnight(board.WHITE),
					2: board.NewBishop(board.WHITE, bounds),
					3: board.NewQueen(board.WHITE, bounds),
					4: board.NewKing(board.WHITE),
					5: board.NewBishop(board.WHITE, bounds),
					6: board.NewKnight(board.WHITE),
					7: board.NewRook(board.WHITE, bounds),
				},
			},
		),
		board.WithIllegalStateFilter(
			board.NewIllegalStateFilter(
				&board.IllegalCheckStateFilter{
					TurnState: turnState,
				},
			),
		),
		board.WithTurnState(turnState),
	)
}

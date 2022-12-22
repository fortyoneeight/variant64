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

func NewGameboardState(bounds Bounds, state GameboardState) GameboardState {
	gameboardState := GameboardState{}
	for rank := 0; rank < bounds.RankCount; rank += 1 {
		gameboardState[rank] = map[int]*Piece{}
		for file := 0; file < bounds.FileCount; file += 1 {
			gameboardState[rank][file] = nil
		}
	}

	for rank := range state {
		for file := range state[rank] {
			gameboardState[rank][file] = state[rank][file]
		}
	}

	return gameboardState
}

// CopyGameboardState returns a copy of the provided GameboardState.
func CopyGameboardState(state GameboardState) GameboardState {
	copiedState := GameboardState{}

	for rank, files := range state {
		copiedState[rank] = map[int]*Piece{}
		for file, piece := range files {
			copiedState[rank][file] = piece
		}
	}

	return copiedState
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

type AvailableMoveMap map[int]map[int]MoveMap

func NewAvailableMoveMap(bounds Bounds) AvailableMoveMap {
	availableMoveMap := AvailableMoveMap{}
	for rank := 0; rank < bounds.RankCount; rank++ {
		availableMoveMap[rank] = map[int]MoveMap{}
		for file := 0; file < bounds.FileCount; file++ {
			availableMoveMap[rank][file] = NewMoveMap()
		}
	}
	return availableMoveMap
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

type TurnState struct {
	Active    Color
	TurnOrder []Color
}

func (t *TurnState) GetActivePlayer() Color {
	return t.Active
}

func (t *TurnState) PassTurn() {
	t.Active = t.TurnOrder[0]
	t.TurnOrder = append(t.TurnOrder[1:], t.TurnOrder[0])
}

type EndStateType string

const (
	EndStateNone      EndStateType = "none"
	EndStateCheckmate EndStateType = "checkmate"
	EndStateStalemate EndStateType = "stalemate"
)

type GameEndState struct {
	EndStateType
	Winner Color
	Loser  Color
}

type builderOption = func(c *Builder)

type Builder struct {
	bounds             Bounds
	castlingState      *CastlingState
	moveApplicator     *MoveApplicator
	moveFilter         *MoveFilter
	illegalStateFilter *IllegalStateFilter
	gameboardState     GameboardState
	turnState          *TurnState
	gameEndState       GameEndState
}

func NewBuilder() *Builder {
	bounds := Bounds{RankCount: 8, FileCount: 8}
	castlingState := NewDefaultCastlingState()
	turnState := &TurnState{
		Active:    WHITE,
		TurnOrder: []Color{BLACK, WHITE},
	}
	return &Builder{
		bounds:        bounds,
		castlingState: castlingState,
		moveApplicator: NewMoveApplicator(
			&SinglePieceMoveApplicator{},
			&KingsideCastleMoveApplicator{},
			&QueensideCastleMoveApplicator{},
		),
		moveFilter: NewMoveFilter(
			&FilterOutOfBounds{Bounds: bounds},
			&FilterPieceCollision{},
			&FilterFriendlyCapture{},
			&FilterInvalidPawnDoublePush{},
			&FilterIllegalKingsideCastle{
				CastlingState: castlingState,
			},
			&FilterIllegalQueensideCastle{
				CastlingState: castlingState,
			},
		),
		illegalStateFilter: NewIllegalStateFilter(
			&IllegalCheckStateFilter{
				TurnState: turnState,
			},
		),
		gameboardState: NewGameboardState(bounds, GameboardState{}),
		turnState:      turnState,
		gameEndState: GameEndState{
			EndStateType: EndStateNone,
			Winner:       NO_COLOR,
			Loser:        NO_COLOR,
		},
	}
}

func WithBounds(bounds Bounds) builderOption {
	return func(c *Builder) {
		c.bounds = bounds
	}
}

func WithCastlingState(castlingState *CastlingState) builderOption {
	return func(c *Builder) {
		c.castlingState = castlingState
	}
}

func WithMoveApplicator(moveApplicator *MoveApplicator) builderOption {
	return func(c *Builder) {
		c.moveApplicator = moveApplicator
	}
}

func WithMoveFilter(moveFilter *MoveFilter) builderOption {
	return func(c *Builder) {
		c.moveFilter = moveFilter
	}
}

func WithIllegalStateFilter(illegalStateFilter *IllegalStateFilter) builderOption {
	return func(c *Builder) {
		c.illegalStateFilter = illegalStateFilter
	}
}

func WithGameboardState(state GameboardState) builderOption {
	return func(c *Builder) {
		c.gameboardState = NewGameboardState(c.bounds, state)
	}
}

func WithTurnState(state *TurnState) builderOption {
	return func(c *Builder) {
		c.turnState = state
	}
}

func Build(options ...builderOption) *Board {
	builder := NewBuilder()
	for _, option := range options {
		option(builder)
	}
	board := &Board{
		Bounds:             builder.bounds,
		MoveApplicator:     builder.moveApplicator,
		MoveFilter:         builder.moveFilter,
		IllegalStateFilter: builder.illegalStateFilter,
		CastlingState:      builder.castlingState,
		GameboardState:     builder.gameboardState,
		TurnState:          builder.turnState,
		GameEndState:       builder.gameEndState,
	}
	board.updateMoves()
	return board
}

type Board struct {
	Bounds
	*MoveApplicator
	*MoveFilter
	*IllegalStateFilter
	*CastlingState
	GameboardState
	*TurnState
	GameEndState
}

// GetState returns a GameboardState for the Board.
func (b *Board) GetState() GameboardState {
	return b.GameboardState
}

// HandleMove handles a Move submitted by the client.
func (b *Board) HandleMove(move Move) error {
	// Check if there is a piece at the source position.
	sourcePiece := b.GameboardState[move.Source.Rank][move.Source.File]
	if sourcePiece == nil {
		return errPieceNotFound
	}

	// Verify move is an available move.
	isAvailableMove := sourcePiece.IsAvailableMove(move)
	if !isAvailableMove {
		return errMoveNotAllowed
	}

	// Update the castle flags if necessary.
	if b.CastlingState != nil {
		b.UpdateCastleState(move, b.GameboardState)
	}

	// Update the board state.
	updatedState, moveErr := b.ApplyMove(move, b.GameboardState)
	if moveErr != nil {
		return moveErr
	}
	b.GameboardState = updatedState

	// Pass the turn.
	b.PassTurn()

	// Update the available moves for each piece.
	b.updateMoves()

	// Check for game ending.
	switch b.checkGameEnd() {
	case EndStateCheckmate:
		b.GameEndState.EndStateType = EndStateCheckmate
		b.Winner = b.TurnState.TurnOrder[0]
		b.Loser = b.TurnState.Active
	case EndStateStalemate:
		b.GameEndState.EndStateType = EndStateStalemate
	}

	return nil
}

// updateMoves updates all the pieces moves by applying filters and state checks
func (b *Board) updateMoves() {
	// Update the available moves for each piece.
	possibleMoves := b.generatePossibleMoves(b.GameboardState)
	filtered := b.filterAvailableMoveMap(b.GameboardState, possibleMoves, b.legalMoveFilterPredicate)
	filtered = b.filterAvailableMoveMap(b.GameboardState, filtered, b.legalGamebaordStatePredicate)

	b.setAvailableMoves(possibleMoves, b.GameboardState)
}

// filterAvailableMoveMap filters moves in the provided AvailableMoveMap with the provided predicate.
func (b *Board) filterAvailableMoveMap(
	state GameboardState,
	availableMoveMap AvailableMoveMap,
	predicate legalMovePredicate,
) AvailableMoveMap {
	filteredAvailableMoveMap := NewAvailableMoveMap(b.Bounds)
	for rank, files := range availableMoveMap {
		for file, moveMap := range files {
			filteredAvailableMoveMap[rank][file] = b.filterMoveMap(
				Position{Rank: rank, File: file},
				state,
				moveMap,
				predicate,
			)
		}
	}
	return filteredAvailableMoveMap
}

// filterMoveMap filters moves in the provided MoveMap with the provided predicate.
func (b *Board) filterMoveMap(
	position Position,
	state GameboardState,
	moveMap MoveMap,
	predicate legalMovePredicate,
) MoveMap {
	filteredMoveMap := NewMoveMap()
	for moveType := range moveMap {
		for _, destination := range moveMap[moveType] {
			piece := state[position.Rank][position.File]
			move := Move{
				Source:      position,
				Destination: destination,
				MoveType:    moveType,
			}
			if predicate(piece, move, state) {
				filteredMoveMap[moveType] = append(filteredMoveMap[moveType], destination)
			}
		}
	}
	return filteredMoveMap
}

// generatePossibleMoves generates the possible moves for each piece in the game.
func (b *Board) generatePossibleMoves(state GameboardState) AvailableMoveMap {
	availableMoveMap := NewAvailableMoveMap(b.Bounds)
	b.forEachPiece(
		state,
		func(source Position, piece *Piece) {
			if piece != nil {
				availableMoveMap[source.Rank][source.File] = piece.GenerateMoves(source)
			}
		},
	)
	return availableMoveMap
}

// getAvailableMoves gets the available moves for each piece in the game.
func (b *Board) getAvailableMoves() AvailableMoveMap {
	availableMoveMap := NewAvailableMoveMap(b.Bounds)
	b.forEachPiece(
		b.GameboardState,
		func(source Position, piece *Piece) {
			if piece != nil {
				availableMoveMap[source.Rank][source.File] = piece.AvailableMoves
			}
		},
	)
	return availableMoveMap
}

// setAvailableMoves sets the available moves for each piece in the game.
func (b *Board) setAvailableMoves(availableMoveMap AvailableMoveMap, state GameboardState) {
	b.forEachPiece(
		state,
		func(source Position, piece *Piece) {
			if piece != nil {
				piece.AvailableMoves = availableMoveMap[source.Rank][source.File]
			}
		},
	)
}

// legalMovePredicate is used to filter moves that are not allowed in the game.
type legalMovePredicate = func(piece *Piece, move Move, state GameboardState) bool

// legalMoveFilterPredicate returns true if the provided move passes the MoveFilter.
func (b *Board) legalMoveFilterPredicate(piece *Piece, move Move, state GameboardState) bool {
	return b.IsLegalMove(move, state)
}

// legalGamebaordStatePredicate returns true if the provided move results in a legal board state.
func (b *Board) legalGamebaordStatePredicate(piece *Piece, move Move, state GameboardState) bool {
	copiedState := CopyGameboardState(state)
	copiedState, err := b.ApplyMove(move, copiedState)
	if err != nil {
		return false
	}

	potentialNextTurnMoves := b.filterAvailableMoveMap(
		copiedState,
		b.generatePossibleMoves(copiedState),
		b.legalMoveFilterPredicate,
	)

	return b.IsLegalState(piece.Color, copiedState, potentialNextTurnMoves)
}

// forEachPiece applies the function to each piece on the board.
func (b *Board) forEachPiece(state GameboardState, fn func(position Position, piece *Piece)) {
	for rank, files := range state {
		for file, piece := range files {
			fn(Position{Rank: rank, File: file}, piece)
		}
	}
}

// checkGameEnd checks if the game has ended.
// If the active player has no moves they have lost.
func (b *Board) checkGameEnd() EndStateType {
	activePlayerHasMove := false
	b.forEachPiece(
		b.GameboardState,
		func(position Position, piece *Piece) {
			if !activePlayerHasMove &&
				piece != nil &&
				piece.Color == b.GetActivePlayer() {
				for _, moveList := range piece.AvailableMoves {
					if len(moveList) != 0 {
						activePlayerHasMove = true
					}
				}
			}
		},
	)

	activePlayerIsInCheck := !anyPosition(
		b.GetActivePlayer(),
		b.GameboardState,
		predicateAttackingEnemyKing(
			b.GetActivePlayer(),
			b.GameboardState,
			b.getAvailableMoves()),
	)

	switch {
	case !activePlayerHasMove && activePlayerIsInCheck:
		return EndStateCheckmate
	case !activePlayerHasMove && !activePlayerIsInCheck:
		return EndStateStalemate
	default:
		return EndStateNone
	}
}

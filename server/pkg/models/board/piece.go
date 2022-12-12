package board

type moveGenerator interface {
	GetMoves(source Position) MoveMap
}

// Piece is used to represent a movable entity on a board.
type Piece struct {
	Color          Color     `json:"color"`
	PieceType      PieceType `json:"piece_type"`
	moveGenerators []moveGenerator
}

// GetColor returns the Piece's Color.
func (p *Piece) GetColor() Color {
	return p.Color
}

// GetType returns the Piece's PieceType.
func (p *Piece) GetType() PieceType {
	return p.PieceType
}

// GetMoves returns a list of possible moves for the Piece at the given position.
func (p *Piece) GetMoves(source Position) MoveMap {
	moveMap := NewMoveMap()
	for _, generator := range p.moveGenerators {
		JoinMoveMaps(moveMap, generator.GetMoves(source))
	}
	return moveMap
}

// NewPiece creates a new Piece based on the input parameters.
func NewPiece(color Color, pieceType PieceType, moveGenerators ...moveGenerator) *Piece {
	return &Piece{
		Color:          color,
		PieceType:      pieceType,
		moveGenerators: moveGenerators,
	}
}

// NewPawn creates a Piece with type PAWN.
func NewPawn(color Color) *Piece {
	var direction Direction
	switch color {
	case WHITE:
		direction = North
	case BLACK:
		direction = South
	}

	return NewPiece(
		color,
		PAWN,
		&SingleNormalMoveGenerator{direction: direction},
		&DoublePushMoveGenerator{color: color},
		&SingleDiagonalCaputureMoveGenerator{color: color},
	)
}

// NewKnight creates a Piece with type KNIGHT.
func NewKnight(color Color) *Piece {
	return NewPiece(
		color,
		KNIGHT,
		&KnightMoveGenerator{},
	)
}

// NewRook creates a Piece with type ROOK.
func NewRook(color Color, bounds Bounds) *Piece {
	return NewPiece(
		color,
		ROOK,
		&RayMoveGenerator{direction: North, bounds: bounds},
		&RayMoveGenerator{direction: East, bounds: bounds},
		&RayMoveGenerator{direction: South, bounds: bounds},
		&RayMoveGenerator{direction: West, bounds: bounds},
	)
}

// NewBishop creates a Piece with type BISHOP.
func NewBishop(color Color, bounds Bounds) *Piece {
	return NewPiece(
		color,
		BISHOP,
		&RayMoveGenerator{direction: NorthEast, bounds: bounds},
		&RayMoveGenerator{direction: SouthEast, bounds: bounds},
		&RayMoveGenerator{direction: SouthWest, bounds: bounds},
		&RayMoveGenerator{direction: NorthWest, bounds: bounds},
	)
}

// NewQueen creates a Piece with type QUEEN.
func NewQueen(color Color, bounds Bounds) *Piece {
	return NewPiece(
		color,
		QUEEN,
		NewBishop(color, bounds),
		NewRook(color, bounds),
	)
}

// NewKing creates a Piece with type KING.
func NewKing(color Color) *Piece {
	return NewPiece(
		color,
		KING,
		&SingleNormalMoveGenerator{direction: North},
		&SingleNormalMoveGenerator{direction: NorthEast},
		&SingleNormalMoveGenerator{direction: East},
		&SingleNormalMoveGenerator{direction: SouthEast},
		&SingleNormalMoveGenerator{direction: South},
		&SingleNormalMoveGenerator{direction: SouthWest},
		&SingleNormalMoveGenerator{direction: West},
		&SingleNormalMoveGenerator{direction: NorthWest},
		&CastleMoveGenerator{color: color},
	)
}

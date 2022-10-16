export interface Board {
  size: {
    length: number;
    width: number;
  };
  cells: BoardCells;
}

export enum BoardCellItemTypes {
  PIECE = 'piece',
}

export type BoardCells = Array<BoardCellItem>;
export interface BoardCellItem {
  type: BoardCellItem;
}

export enum BoardPieceTypes {
  PAWN = 'pawn',
  BISHOP = 'bishop',
  KNIGHT = 'knight',
  ROOK = 'rook',
  QUEEN = 'queen',
  KING = 'king',
}

export interface BoardPiece extends BoardCellItem {
  name: BoardPieceTypes;
  player: String;
  moves: BoardMoves;
}

export type BoardMoves = Array<BoardMove>;
export interface BoardMove {
  x: number;
  y: number;
}

export interface BoardClocks {
  [key: string]: string;
}

export interface BoardPlayer {
  id: string;
  display_name: string;
}

import { type } from 'os';

export interface Board {
  size: {
    length: number;
    width: number;
  };
  cells: BoardCells;
  clocks: BoardClocks;
}

export interface Coordinate {
  x: number;
  y: number;
}

export enum BoardCellItemTypes {
  PIECE = 'piece',
}

export type BoardCells = Array<BoardPiece>;
export interface BoardCellItem extends Coordinate {
  type: BoardCellItemTypes;
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
export type BoardMove = Coordinate;

export interface BoardClocks {
  [key: string]: string;
}

export interface BoardPlayer {
  player_id: string;
  display_name: string;
}

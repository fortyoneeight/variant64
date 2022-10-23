import {
  Board,
  BoardCellItemTypes,
  BoardPiece,
  BoardPieceTypes,
} from '../types';

const boardWidth = 8;
const boardLength = 8;

function genPawn(n: number, startingRow: number, player: string): BoardPiece {
  return {
    x: startingRow,
    y: n,
    type: BoardCellItemTypes.PIECE,
    moves: [
      {
        x: n + 1,
        y: n + 1,
      },
    ],
    name: BoardPieceTypes.PAWN,
    player,
  };
}

function genPlayerPawns(
  startingRow: number,
  player: string
): Array<BoardPiece> {
  return Array(boardWidth)
    .fill('')
    .map((val, index) => genPawn(index, startingRow, player));
}

const pawns = [
  ...genPlayerPawns(1, 'player1'),
  ...genPlayerPawns(boardLength - 2, 'player2'),
];

function genRook(
  startingRow: number,
  startingColumn: number,
  player: string
): BoardPiece {
  return {
    x: startingRow,
    y: startingColumn,
    type: BoardCellItemTypes.PIECE,
    moves: [
      {
        x: 0,
        y: 0,
      },
    ],
    name: BoardPieceTypes.ROOK,
    player,
  };
}

function genPlayerRooks(
  startingRow: number,
  player: string
): Array<BoardPiece> {
  const columnOffset = 0;
  return [
    genRook(startingRow, columnOffset, player),
    genRook(startingRow, boardWidth - 1 - columnOffset, player),
  ];
}

const rooks = [
  ...genPlayerRooks(0, 'player1'),
  ...genPlayerRooks(boardWidth - 1, 'player2'),
];

function genKnight(
  startingRow: number,
  startingColumn: number,
  player: string
): BoardPiece {
  return {
    x: startingRow,
    y: startingColumn,
    type: BoardCellItemTypes.PIECE,
    moves: [
      {
        x: 0,
        y: 0,
      },
    ],
    name: BoardPieceTypes.KNIGHT,
    player,
  };
}

function genPlayerKnights(
  startingRow: number,
  player: string
): Array<BoardPiece> {
  const columnOffset = 1;
  return [
    genKnight(startingRow, columnOffset, player),
    genKnight(startingRow, boardWidth - 1 - columnOffset, player),
  ];
}

const knights = [
  ...genPlayerKnights(0, 'player1'),
  ...genPlayerKnights(boardWidth - 1, 'player2'),
];

function genBishop(
  startingRow: number,
  startingColumn: number,
  player: string
): BoardPiece {
  return {
    x: startingRow,
    y: startingColumn,
    type: BoardCellItemTypes.PIECE,
    moves: [
      {
        x: 0,
        y: 0,
      },
    ],
    name: BoardPieceTypes.BISHOP,
    player,
  };
}

function genPlayerBishops(
  startingRow: number,
  player: string
): Array<BoardPiece> {
  const columnOffset = 2;
  return [
    genBishop(startingRow, columnOffset, player),
    genBishop(startingRow, boardWidth - 1 - columnOffset, player),
  ];
}

const bishops = [
  ...genPlayerBishops(0, 'player1'),
  ...genPlayerBishops(boardWidth - 1, 'player2'),
];

function genQueen(
  startingRow: number,
  startingColumn: number,
  player: string
): BoardPiece {
  return {
    x: startingRow,
    y: startingColumn,
    type: BoardCellItemTypes.PIECE,
    moves: [
      {
        x: 0,
        y: 0,
      },
    ],
    name: BoardPieceTypes.QUEEN,
    player,
  };
}

function genPlayerQueen(
  startingRow: number,
  player: string
): Array<BoardPiece> {
  const columnOffset = 3;
  return [genQueen(startingRow, columnOffset, player)];
}

const queens = [
  ...genPlayerQueen(0, 'player1'),
  ...genPlayerQueen(boardWidth - 1, 'player2'),
];

function genKing(
  startingRow: number,
  startingColumn: number,
  player: string
): BoardPiece {
  return {
    x: startingRow,
    y: startingColumn,
    type: BoardCellItemTypes.PIECE,
    moves: [
      {
        x: 0,
        y: 0,
      },
    ],
    name: BoardPieceTypes.KING,
    player,
  };
}

function genPlayerKing(startingRow: number, player: string): Array<BoardPiece> {
  const columnOffset = 4;
  return [genKing(startingRow, columnOffset, player)];
}

const kings = [
  ...genPlayerKing(0, 'player1'),
  ...genPlayerKing(boardWidth - 1, 'player2'),
];

export const mockdataBoard: Board = {
  size: {
    length: 8,
    width: boardWidth,
  },
  cells: [...pawns, ...rooks, ...knights, ...bishops, ...queens, ...kings],
  clocks: {
    player1: '100',
    player2: '100',
  },
};

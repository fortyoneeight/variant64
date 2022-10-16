import { Board, BoardCellItemTypes, BoardPieceTypes } from '../types';

export const mockdataBoard: Board = {
  size: {
    length: 8,
    width: 8,
  },
  cells: [
    {
      type: BoardCellItemTypes.PIECE,
      moves: [
        {
          x: 0,
          y: 0,
        },
      ],
      name: BoardPieceTypes.PAWN,
      player: 'player1',
    },
  ],
  clocks: {
    player1: '100',
    player2: '100',
  },
};

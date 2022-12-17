import React from 'react';
import { Board } from '../../../models';
import { pieceImages } from '../../../assets/pieceImages';
import { GameboardGrid, GridCell, PieceImage } from './BoardDisplay.styled';

export default function BoardDisplay({ board }: { board: Board }) {
  return (
    <GameboardGrid rows={board.size.length} columns={board.size.width}>
      {[...Array(board.size.length)].map((row: any, i: any) => {
        return [...Array(board.size.width)].map((cell: any, j: any) => {
          const color = (i % 2) + (j % 2) === 1 ? 'white' : '#549e4c';
          return (
            <GridCell key={i + '' + j} style={{ backgroundColor: color }}>
              {board.cells.map((piece: any, k: any) => {
                if (piece.x === i && piece.y === j) {
                  return (
                    <PieceImage
                      key={k}
                      src={
                        pieceImages[piece.name as keyof typeof pieceImages][
                          (piece.player == 'player1' ? 'white' : 'black')!
                        ]
                      }
                    />
                  );
                }
              })}
            </GridCell>
          );
        });
      })}
    </GameboardGrid>
  );
}

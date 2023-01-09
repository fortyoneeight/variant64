import React from 'react';
import { Board } from '../../../models';
import { GameboardGrid } from './BoardDisplay.styled';
import Piece from './Piece';
import Square from './Square';

export default function BoardDisplay({ board }: { board: Board }) {
  return (
    <GameboardGrid rows={board.size.length} columns={board.size.width}>
      {[...Array(board.size.length)].map((row: any, i: any) => {
        return [...Array(board.size.width)].map((cell: any, j: any) => {
          const color = (i % 2) + (j % 2) === 1 ? 'white' : '#549e4c';
          return (
            <Square key={i + j} i={i} j={j} color={color}>
              {board.cells.map((piece: any, k: any) => {
                if (piece.x === i && piece.y === j) {
                  return <Piece key={k} piece={piece}></Piece>;
                }
              })}
            </Square>
          );
        });
      })}
    </GameboardGrid>
  );
}

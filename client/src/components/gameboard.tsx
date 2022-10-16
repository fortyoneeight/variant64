import React from 'react';
import { Board } from '../types';

export default function Gameboard({ board }: { board: Board }) {
  console.log(board);

  const numCells = board.size.length * board.size.width;

  return (
    <div
      className='gameboard-grid'
      style={{
        gridTemplateRows: `repeat(${board.size.length}, 10vh)`,
        gridTemplateColumns: `repeat(${board.size.width}, 10vh)`,
      }}
    >
      {[...Array(board.size.length)].map((row: any, i: any) => {
        return [...Array(board.size.width)].map((cell: any, j: any) => {
          const color = (i % 2) + (j % 2) == 1 ? 'white' : 'blue';
          return (
            <div
              key={i}
              className='grid-cell'
              style={{ backgroundColor: color }}
            ></div>
          );
        });
      })}
    </div>
  );
}

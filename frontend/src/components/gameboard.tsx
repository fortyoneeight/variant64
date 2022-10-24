import React from 'react';
import { Board, BoardPiece } from '../types';

export default function Gameboard({ board }: { board: Board }) {
  console.log(board);

  return (
    <div
      className="gameboard-grid"
      style={{
        gridTemplateRows: `repeat(${board.size.length}, 7vh)`,
        gridTemplateColumns: `repeat(${board.size.width}, 7vh)`,
      }}
    >
      {[...Array(board.size.length)].map((row: any, i: any) => {
        return [...Array(board.size.width)].map((cell: any, j: any) => {
          const color = (i % 2) + (j % 2) === 1 ? 'white' : '#549e4c';
          return (
            <div key={i + '' + j} className="grid-cell" style={{ backgroundColor: color }}>
              {board.cells.map((piece: any, k: any) => {
                if (piece.x === i && piece.y === j) {
                  return <p>{piece.name}</p>;
                }
              })}
            </div>
          );
        });
      })}
    </div>
  );
}

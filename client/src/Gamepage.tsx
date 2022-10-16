import React from 'react';
import Gameboard from './components/gameboard';
import { mockdataBoard } from './mockdata/board';

export default function Gamepage() {
  return (
    <div className='column'>
      <Gameboard board={mockdataBoard} />
    </div>
  );
}

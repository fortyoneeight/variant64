import React from "react";
import Gameboard from "./components/gameboard";
import { mockdataBoard } from "./mockdata/board";

export default function Gamepage() {
  return (
    <div className="column">
      <div className="row">
        <p className="name">Name</p>
        <p className="clock">{mockdataBoard.clocks.player1}</p>
      </div>
      <Gameboard board={mockdataBoard} />
      <div className="row">
        <p className="name">Name</p>
        <p className="clock">{mockdataBoard.clocks.player2}</p>
      </div>
      <div className='gameplayButtonContainer'>
        <button className='drawButton'>Offer Draw</button>
        <button className='concedeButton'>Concede</button>
      </div>
    </div>
  );
}

import React from "react";
import Gameboard from "./components/gameboard";
import { mockdataBoard } from "./mockdata/board";

export default function Gamepage() {
  return (
    <div className="column">
      <div className="row">
        <p className="name">Name</p>
        <p className="clock">Clock</p>
      </div>
      <Gameboard board={mockdataBoard} />
      <div className="row">
        <p className="name">Name</p>
        <p className="clock">Clock</p>
      </div>
    </div>
  );
}

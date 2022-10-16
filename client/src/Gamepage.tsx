import React from "react";
import Gameboard from "./components/gameboard";

const snapshot = {
    "board": {
        "size": {
            "length": 8,
            "width" : 8
        },
        "cells": [[0,0,0,0,0,0,0,0],
                  [0,0,0,0,0,0,0,0],
                  [0,0,0,0,0,0,0,0],
                  [0,0,0,0,0,0,0,0],
                  [0,0,0,0,0,0,0,0],
                  [0,0,0,0,0,0,0,0],
                  [0,0,0,0,0,0,0,0],
                  [0,0,0,0,0,0,0,0]],
        "active_player":"ben",
        "clocks":{
            "player1": 100,
            "player2": 100,
        }
    }
}

export default function Gamepage() {

    return <div className='column'>
        <Gameboard board={snapshot.board}/>
    </div>
    
}

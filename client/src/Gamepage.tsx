import React from "react";
import Gameboard from "./components/gameboard";

const snapshot = {
    "board": {
        "size": {
            "length": 8,
            "width" : 8
        },
        "cells": [{x: 1, y: 1, cellItem: {type: 'piece', data: { name: 'pawn', player: {id:'uuid', display_name: 'player1'}, moves: ['a2']}}}],
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

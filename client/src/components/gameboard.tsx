import React from "react";

export default function Gameboard(snapshot: any) {

    console.log(snapshot);

    const numCells = snapshot.board.size.length * snapshot.board.size.width;

    return <div className='gameboard-grid' style={{
            gridTemplateRows: `repeat(${snapshot.board.size.length}, 10vh)`,
            gridTemplateColumns: `repeat(${snapshot.board.size.width}, 10vh)`
        }}>
            {[...Array(snapshot.board.size.length)].map((row: any, i: any) => {
                return [...Array(snapshot.board.size.width)].map((cell: any, j: any) =>{
                    const color = i % 2 + j % 2 == 1 ? 'white' : 'blue'  
                    return <div key={i} className='grid-cell' style={{backgroundColor: color}}>
                        
                    </div>
                }) 
            })}
    </div>
}
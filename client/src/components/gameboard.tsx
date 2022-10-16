import React from "react";

export default function Gameboard(snapshot: any) {


    console.log("potato rotato :) we have le mcgriddle")

    console.log(snapshot);



    return <div className='gameboard-grid'
        style={{
            gridTemplateRows: `repeat(${snapshot.board.size.length}, 10vh)`,
            gridTemplateColumns: `repeat(${snapshot.board.size.width}, 10vh)`
        }}>

        {snapshot.board.cells.map((row: any, i: any) => {
            {return row.map((cell: any, j: any) => {   
                return <div key={j} className='grid-cell' style={{backgroundColor: (j + i) % 2 == 1 ? 'white' : '#21823b',
                }}>
                    <p style={{fontSize:'4rem'}}>{cell}</p>
                </div>
            })}
        })}
    </div>
}
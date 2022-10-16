import React from "react";

export default function Gameboard(snapshot: any) {


    console.log("potato rotato :) we have le mcgriddle")

    console.log(snapshot);

    const cells = snapshot.board.size.length * snapshot.board.size.width;

    return <div className='gameboard-grid' style={{
            gridTemplateRows: `repeat(${snapshot.board.size.length}, 10vh)`,
            gridTemplateColumns: `repeat(${snapshot.board.size.width}, 10vh)`
        }}>
            {[...Array(cells)].map((cell: any, i: any) => {  
                return <div key={i} className='grid-cell' style={{backgroundColor: i % 2 == 1 ? 'white' : '#21823b'}}>
                    <p style={{fontSize:'4rem'}}></p>
                </div>
            })}
    </div>
}
import React from "react";

export default function Gameboard(snapshot: any) {


    console.log("potato rotato :) we have le mcgriddle")

    console.log(snapshot);



    return <div className='gameboard-grid'
        style={{
            gridTemplateRows: `repeat(${snapshot.board.size.length}, 100px)`,
            gridTemplateColumns: `repeat(${snapshot.board.size.width}, 100px)`
        }}>

        {snapshot.board.cells.map((row: any) => {
            {
                return row.map((cell: any, j: any) => {

                    //test indexing by y
                    // if(j == 2){
                    //     return <p>secret elf</p>
                    // }
                    
                    
                    return <div key={j} style={{
                        width: '100px',
                        height: '100px',
                        backgroundColor: 'blue'
                    }}></div>
                })
            }
        })}
    </div>
}
import React from "react";
import { Link } from "react-router-dom";
import './index.css'

const sampleNames = 'Henry Izak Jarron Ben'.split(' ')

function renderRoom(data: { name: string }) {
    return (<div>
        <Link
            to={'/chess'} //`/join/{data.name}`
            className='grid-2-horizontal-leftbias'>


            <span>{data.name}'s room</span> <span className="outline">0/2</span>

        </Link>
    </div>)
}

function RoomList() {
    return (<div className="roomList">

        <aside>
            <h1>Existing Rooms</h1>
        </aside>

        <main>
            {
                sampleNames.map((name) => {
                    return renderRoom({ name: name })
                })
            }
        </main>


    </div>)
}


export default function Homepage() {

    return (<div className="grid-2-horizontal">

        <RoomList />

        <button
            className="createRoom"
            onClick={
                () => {
                    console.log('wowie!!! someone PLEASE impelemnt createThatRoom()!!!!!! T_T')
                    return "potato";
                }
            }>
            Create a room
        </button>
    </div>)
}
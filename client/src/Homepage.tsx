import { Link } from "react-router-dom";
import './index.css'
import {
    selector,
    useRecoilState,
    useRecoilValue,
} from 'recoil';
import { Room } from "./types";

import { roomsState } from "./store/atoms/rooms";

function renderRoom(room: Room) {
    return (<div>
        <Link
            to={'/chess'} //`/join/{data.name}`
            className='grid-2-horizontal-leftbias'>


            <span>{room.room_name}</span> <span className="outline">{room.players.length}/{room.players_total}</span>

        </Link>
    </div>)
}

function RoomList() {

    const [rooms, setRooms] = useRecoilState(roomsState);

    return (


        <div className="roomList">

            <aside>
                <h1>Existing Rooms</h1>
            </aside>

            <main>
                {
                    rooms.map((room) => {
                        return renderRoom({ ...room })
                    })
                }
            </main>


        </div>)
}

export default function Homepage() {

    const [rooms, setRooms] = useRecoilState(roomsState);

    return (<div className="grid-2-horizontal">

        <RoomList />

        {/* <CharacterCounter /> */}

        <button
            className="createRoom"
            onClick={
                () => {
                    console.log('TODO Implement createThatRoom()')

                    var newRooms: Array<Room> = [...rooms, {
                        players: [],
                        players_total: 0,
                        room_name: "Mystery man's room"
                    }]
                    setRooms(newRooms)
                }
            }>
            Create a room
        </button>
    </div>)
}

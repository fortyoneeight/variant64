import React from "react";
import { Link } from "react-router-dom";
import './index.css'
import {
    RecoilRoot,
    atom,
    selector,
    useRecoilState,
    useRecoilValue,
} from 'recoil';

const roomsState = atom({
    key: 'roomsState',
    default: [
        'Henry',
        'Izak',
        'Jarron',
        'Ben',
    ]
})

const textState = atom({
    key: 'textState', // unique ID (with respect to other atoms/selectors)
    default: '', // default value (aka initial value)
});

const charCountState = selector({
    key: 'charCountState', // unique ID (with respect to other atoms/selectors)
    get: ({ get }) => {
        const text = get(textState);

        return text.length;
    },
});

function CharacterCount() {
    const count = useRecoilValue(charCountState);

    return <>Character Count: {count}</>;
}



function CharacterCounter() {
    return (
        <div>
            <TextInput />
            <CharacterCount />
        </div>
    );
}

function TextInput() {

    const [text, setText] = useRecoilState(textState);

    const onChange = (event: any) => {
        setText(event.target.value);
    };

    return (
        <div>
            <input type="text" value={text} onChange={onChange} />
            <br />
            Echo: {text}
        </div>
    );
}


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

    const [rooms, setRooms] = useRecoilState(roomsState);

    return (


        <div className="roomList">

            <aside>
                <h1>Existing Rooms</h1>
            </aside>

            <main>
                {
                    rooms.map((name) => {
                        return renderRoom({ name: name })
                    })
                }
            </main>


        </div>)
}


export default function Homepage() {

    const [text, setText] = useRecoilState(textState);
    const [rooms, setRooms] = useRecoilState(roomsState);

    return (<div className="grid-2-horizontal">

        <RoomList />

        {/* <CharacterCounter /> */}

        <button
            className="createRoom"
            onClick={
                () => {
                    console.log('TODO Implement createThatRoom()')

                    setText(text + " create room clicked")
                    var newRooms = [...rooms, 'mystery man']
                    setRooms(newRooms)
                }
            }>
            Create a room
        </button>
    </div>)
}

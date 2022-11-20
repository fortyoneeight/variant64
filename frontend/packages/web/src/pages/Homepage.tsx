import React, { useMemo } from 'react';
import { Link } from 'react-router-dom';
import '../index.css';
import { useRecoilState } from 'recoil';
import { Room } from '../types';
import { roomsState } from '../store/atoms/rooms';
import { HttpContext } from '../store/context';
import { HomepageService } from './hompage-service';

function renderRoom(room: Room) {
  return (
    <div>
      <Link
        to={'/chess'} //`/join/{data.name}`
        className="grid-2-horizontal-leftbias"
      >
        <span>{room.room_name}</span>{' '}
        <span className="outline">
          {room.players.length}/{room.players_total}
        </span>
      </Link>
    </div>
  );
}

function RoomList() {
  const [rooms, setRooms] = useRecoilState(roomsState);

  return (
    <div className="roomList">
      <aside>
        <h1>Existing Rooms</h1>
      </aside>

      <main>
        {rooms.map((room) => {
          return renderRoom({ ...room });
        })}
      </main>
    </div>
  );
}

export default function Homepage() {
  const [rooms, setRooms] = useRecoilState(roomsState);
  const context = React.useContext(HttpContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomService),
    [context.roomService]
  );

  // Test api calls.
  homepageService.getRooms();

  homepageService
    .createRoom(crypto.randomUUID())
    .then((room) => homepageService.getRoom(room.id))
    .then((room) => {
      homepageService
        .createPlayer('player_' + Math.floor(Math.random() * 10))
        .then((player) => homepageService.getPlayer(player.id))
        .then((player) => {
          homepageService
            .joinRoom(room.id, player.id)
            .then(() => homepageService.leaveRoom(room.id, player.id));
        });
    });

  return (
    <div className="grid-2-horizontal">
      <RoomList />

      {/* <CharacterCounter /> */}

      <button
        className="createRoom"
        onClick={() => {
          console.log('TODO Implement createThatRoom()');

          var newRooms: Array<Room> = [
            ...rooms,
            {
              id: '1',
              players: [],
              players_total: 0,
              room_name: "Mystery man's room",
            },
          ];
          setRooms(newRooms);
        }}
      >
        Create a room
      </button>
    </div>
  );
}

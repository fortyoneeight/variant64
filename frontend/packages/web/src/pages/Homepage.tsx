import React, { useEffect, useMemo } from 'react';
import { Link } from 'react-router-dom';
import '../index.css';
import { useRecoilState, useRecoilValue } from 'recoil';
import { roomsState, playerState, roomState } from '../store/atoms';
import { ServicesContext } from '../store/context';
import { HomepageService } from './hompage-service';
import { PlayerForm, RoomForm } from '../components';

function RoomList() {
  const rooms = useRecoilValue(roomsState);
  const [_, setRoom] = useRecoilState(roomState);

  return (
    <div className="hompage_room-list">
      <aside>
        <h1>Existing Rooms</h1>
      </aside>

      <main>
        {rooms.map((room) => {
          return (
            <Link
              key={room.id}
              className="hompage_room-link outline"
              to={`/room/${room.id}`}
              onClick={() => setRoom(room)}
            >
              <p>Room: {room.name}</p>
              <span>Players: {room.players.length}</span>
            </Link>
          );
        })}
      </main>
    </div>
  );
}

function PlayerDetails() {
  const [player, _] = useRecoilState(playerState);
  return (
    <div>
      <p>Player ID: {player.id}</p>
      <p>Player Name: {player.display_name}</p>
    </div>
  );
}
export default function Homepage() {
  const [rooms, setRooms] = useRecoilState(roomsState);
  const [player, setPlayer] = useRecoilState(playerState);

  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService),
    [context.roomHttpService]
  );

  // Test api calls.
  useEffect(() => {
    homepageService.getRooms().then((rooms) => setRooms(rooms));
  }, []);

  const handleRoomSubmit = (homepageService: HomepageService, roomName: string) => {
    homepageService.createRoom(roomName).then((room) => {
      setRooms([...rooms, room]);
    });
  };

  const handlePlayerSubmit = (homepageService: HomepageService, displayName: string) => {
    homepageService.createPlayer(displayName).then((playerResponse) => {
      setPlayer({ ...player, ...playerResponse });
    });
  };

  const playerSection = () =>
    player.id ? (
      <PlayerDetails />
    ) : (
      <PlayerForm
        submitForm={(displayName: string) => handlePlayerSubmit(homepageService, displayName)}
      />
    );
  return (
    <>
      <div>
        {playerSection()}

        <RoomForm submitForm={(roomName: string) => handleRoomSubmit(homepageService, roomName)} />
      </div>
      <div>
        <RoomList />

        {/* <CharacterCounter /> */}
      </div>
    </>
  );
}

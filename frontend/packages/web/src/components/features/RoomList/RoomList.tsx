import React, { useEffect, useMemo } from 'react';
import { Link } from 'react-router-dom';
import { useRecoilState, useRecoilValue } from 'recoil';
import { roomsState, roomState } from '../../../store/atoms';
import { HomepageService } from '../../pages/hompage-service';
import { ServicesContext } from '../../../store/context';
import './RoomList.css';

export default function RoomList() {
  const [_, setRoom] = useRecoilState(roomState);
  const [rooms, setRooms] = useRecoilState(roomsState);

  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  useEffect(() => {
    homepageService.getRooms().then((rooms) => setRooms(rooms));
  }, []);

  return (
    <div className="roomListContainer gradientBackground">
      <h1 style={{ color: '#ffffff' }}>Join a Room</h1>

      <main className="listContainer">
        {rooms.map((room) => {
          return (
            <Link
              key={room.id}
              className="roomLink"
              to={`/room/${room.id}`}
              onClick={() => setRoom(room)}
            >
              <div className="roomListRow">
                <p className="roomListTextField">{room.name}</p>
                <span className="roomListTextField">Players: {room.players.length} / 2</span>
              </div>
            </Link>
          );
        })}
      </main>
    </div>
  );
}

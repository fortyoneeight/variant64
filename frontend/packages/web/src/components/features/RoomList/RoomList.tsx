import React, { useEffect, useMemo } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { playerState, roomState, roomsState } from '../../../store/atoms';
import { HomepageService } from '../../pages/hompage-service';
import { ServicesContext } from '../../../store/context';
import { WhiteButton } from '../../atoms';
import { ListContainer, RoomListContainer } from './RoomList.Styled';
import { Room } from '../../../models';

export default function RoomList() {
  const [player, _] = useRecoilState(playerState);
  const [room, setRoom] = useRecoilState(roomState);
  const [rooms, setRooms] = useRecoilState(roomsState);

  const navigate = useNavigate();

  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  const handleJoinClick = (r: Room) => {
    homepageService.joinRoom(r.id, player.id).then((roomResponse) => {
      setRoom({ ...room, ...roomResponse });
      navigate('/room/' + r.id);
    });
  };

  useEffect(() => {
    homepageService.getRooms().then((rooms) => setRooms(rooms));
  }, []);

  return (
    <RoomListContainer>
      <h1 style={{ color: '#ffffff' }}>Join a Room</h1>
      <ListContainer>
        {rooms.map((r) => {
          return (
            <WhiteButton key={r.id} onClick={() => handleJoinClick(r)}>
              {r.name} --- Players: {r.players.length} / 2
            </WhiteButton>
          );
        })}
      </ListContainer>
    </RoomListContainer>
  );
}

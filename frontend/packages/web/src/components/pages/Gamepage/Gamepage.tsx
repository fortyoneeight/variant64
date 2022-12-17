import React, { useEffect, useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { playerState, roomState, gameUpdateState } from '../../../store/atoms';
import { ServicesContext } from '../../../store/context';
import { mockdataBoard } from '../../../store/mockdata';
import { BoardDisplay, GameplayButtons, PlayerInfo } from '../../features';
import { HomepageService } from '../hompage-service';
import { GameContainer } from './Gamepage.styled';

export default function Gamepage() {
  const { id } = useParams();
  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  const [player, _] = useRecoilState(playerState);
  const [room, setRoom] = useRecoilState(roomState);
  const [gameUpdate, setGameUpdate] = useRecoilState(gameUpdateState);

  const cb = (data: any) => {
    console.log('[COMPONENT_DATA]', data);
    if (
      data.channel === 'room' &&
      Object.values(data.data.players).find((id) => {
        id == player.id;
      })
    ) {
      setRoom({ ...room, ...data.data });
    }

    if (data.channel === 'game') {
      setGameUpdate({ ...gameUpdate, ...data.data });
    }
  };

  useEffect(() => {
    homepageService.registerCallback(cb);
  }, []);

  useEffect(() => {
    homepageService.subscribeToRoomUpdates(id as string);
  }, [id]);

  return (
    <GameContainer>
      <PlayerInfo />
      <BoardDisplay board={mockdataBoard} />
      <GameplayButtons />
    </GameContainer>
  );
}

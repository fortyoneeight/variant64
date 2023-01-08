import React, { useEffect, useMemo } from 'react';
import { useRecoilState } from 'recoil';
import { Game } from '../../../models';
import { gameState, roomState, roomSubState, gameSubState } from '../../../store/atoms';
import { ServicesContext } from '../../../store/context';
import { HomepageService } from '../../pages/hompage-service';

export default function Subscription() {
  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  const [room, setRoom] = useRecoilState(roomState);
  const [roomSub, setRoomSub] = useRecoilState(roomSubState);

  const [game, setGame] = useRecoilState(gameState);
  const [gameSub, setGameSub] = useRecoilState(gameSubState);

  const handleRoomMessage = (data: any) => {
    setRoom({ ...room, ...data.data });
    if (data.data.game_id) {
      setGame({ id: data.data.game_id } as Game);
    }
  };

  const handleGameMessage = (data: any) => {
    setGame({ ...game, ...data.data });
  };

  const callback = (data: any) => {
    console.log('[COMPONENT_DATA]', data);
    switch (data.channel) {
      case 'room':
        handleRoomMessage(data);
        break;
      case 'game':
        handleGameMessage(data);
        break;
    }
  };

  useEffect(() => {
    if (room.id && roomSub === '') {
      setRoomSub(room.id);
      homepageService.subscribeToRoomUpdates(room.id);
    }
  }, [room, roomSub]);

  useEffect(() => {
    if (game.id && gameSub === '') {
      setGameSub(game.id);
      homepageService.subscribeToGameUpdates(game.id);
    }
  }, [game, gameSub]);

  useEffect(() => {
    homepageService.registerCallback(callback);
  }, []);

  return <></>;
}

import React, { useEffect, useMemo, useState } from 'react';
import { useRecoilState } from 'recoil';
import { Game } from '../../../models';
import { roomSubState, gameSubState, roomState, gameState } from '../../../store/atoms';
import { ServicesContext } from '../../../store/context';
import { HomepageService } from '../../pages/hompage-service';

export default function Subscription() {
  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  const [room, setRoom] = useRecoilState(roomState);
  const [game, setGame] = useRecoilState(gameState);
  const [roomSub, setRoomSub] = useRecoilState(roomSubState);
  const [gameSub, setGameSub] = useRecoilState(gameSubState);
  const [messageState, setMessageState]: any = useState('');

  useEffect(() => {
    if (roomSub && roomSub !== '') {
      homepageService.subscribeToRoomUpdates(roomSub);
    }
  }, [roomSub]);

  useEffect(() => {
    if (gameSub && gameSub !== '') {
      homepageService.subscribeToGameUpdates(gameSub);
    }
  }, [gameSub]);

  useEffect(() => {
    homepageService.registerCallback((data: any) => {
      setMessageState(data);
    });
  }, []);

  useEffect(() => {
    if (messageState !== '') {
      console.log('[COMPONENT_DATA]', messageState);
      switch (messageState.channel) {
        case 'room':
          setRoom({ ...room, ...messageState.data });
          if (messageState.data.game_id) {
            setGameSub(messageState.data.game_id);
          }
          break;
        case 'game':
          setGame({ ...game, ...(messageState.data as Game) });
          break;
      }
    }
  }, [messageState]);

  return <></>;
}

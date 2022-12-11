import React, { useEffect, useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { roomState, playerState, gameUpdateState } from '../../store/atoms';
import { ServicesContext } from '../../store/context';
import { mockdataBoard } from '../../store/mockdata';
import { Gameboard } from '../features';
import { HomepageService } from './hompage-service';

export default function Gamepage() {
  const { id } = useParams();
  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  const [room, setRoom] = useRecoilState(roomState);
  const [player, setPlayer] = useRecoilState(playerState);
  const [gameUpdate, setGameUpdate] = useRecoilState(gameUpdateState);
  const defaultClockMillis = 600000;

  if (!room.id && id) {
    homepageService.getRoom(id).then((roomResponse) => {
      setRoom({
        ...room,
        ...roomResponse,
      });
    });
  }
  const isPlayerPlaying = room.players?.find((roomPlayerID) => roomPlayerID == player.id);

  const handleJoinClick = () => {
    homepageService.joinRoom(room.id, player.id).then((roomResponse) => {
      setRoom({
        ...room,
        ...roomResponse,
      });
    });
  };

  const handleLeaveClick = () => {
    homepageService.leaveRoom(room.id, player.id).then((roomResponse) => {
      setRoom({
        ...room,
        ...roomResponse,
      });
    });
  };

  const handleStartClick = () => {
    homepageService.startRoom(room.id, defaultClockMillis).then((game) => {
      homepageService.subscribeToGameUpdates(game.id);
    });
  };

  const handleConcedeClick = () => {
    if (!gameUpdate.game_id) {
      return;
    }
    homepageService.concedeGame(gameUpdate.game_id, player.id).then((gameResponse) => {
      setGameUpdate({
        ...gameUpdate,
        ...gameResponse,
      });
    });
  };

  const cb = (data: any) => {
    console.log('[COMPONENT_DATA]', data);
    if (data.channel === 'room') {
      setRoom({
        ...room,
        ...data,
      });
    }

    if (data.channel === 'game') {
      setGameUpdate({
        ...gameUpdate,
        ...data,
      });
    }
  };

  useEffect(() => {
    homepageService.registerCallback(cb);
  }, []);

  useEffect(() => {
    homepageService.subscribeToRoomUpdates(id as string);
  }, [id]);

  const joinLeaveButton = isPlayerPlaying ? (
    <button className="drawButton" onClick={() => handleLeaveClick()}>
      Quit Game
    </button>
  ) : (
    <button className="drawButton" onClick={() => handleJoinClick()}>
      Join Game
    </button>
  );

  return (
    <div className="column">
      <div className="row">
        <p className="name">Name</p>
        <p className="clock">{mockdataBoard.clocks.player1}</p>
      </div>
      <Gameboard board={mockdataBoard} />
      <div className="row">
        <p className="name">Name</p>
        <p className="clock">{mockdataBoard.clocks.player2}</p>
      </div>
      <div className="gameplayButtonContainer">
        <button className="startButton" onClick={() => handleStartClick()}>
          Start Game
        </button>
        {joinLeaveButton}
        <button className="drawButton">Offer Draw</button>
        <button className="concedeButton" onClick={() => handleConcedeClick()}>
          Concede
        </button>
      </div>
    </div>
  );
}

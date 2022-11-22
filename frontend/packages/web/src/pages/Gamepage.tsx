import React, { useMemo } from 'react';
import { useParams } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import Gameboard from '../components/gameboard';
import { playerState, roomState } from '../store/atoms';
import { ServicesContext } from '../store/context';
import { mockdataBoard } from '../store/mockdata/board';
import { HomepageService } from './hompage-service';

export default function Gamepage() {
  const { id } = useParams();
  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService),
    [context.roomHttpService]
  );

  const [room, setRoom] = useRecoilState(roomState);
  const [player, setPlayer] = useRecoilState(playerState);

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
        {joinLeaveButton}
        <button className="drawButton">Offer Draw</button>
        <button className="concedeButton">Concede</button>
      </div>
    </div>
  );
}

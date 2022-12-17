import React, { useMemo } from 'react';
import { Link } from 'react-router-dom';
import { GradientButton } from '../../atoms';
import { useRecoilState } from 'recoil';
import { roomState, playerState, gameState } from '../../../store/atoms';
import { ServicesContext } from '../../../store/context';
import { HomepageService } from '../../pages/hompage-service';
import { Room, Game } from '../../../models';
import { ButtonContainer, TopButtonContainer } from './GameplayButtons.styled';

export default function GameplayButtons() {
  const [room, setRoom] = useRecoilState(roomState);
  const [player, _] = useRecoilState(playerState);
  const [game, setGame] = useRecoilState(gameState);
  const defaultClockMillis = 600000;

  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  const handleLeaveClick = () => {
    homepageService.leaveRoom(room.id, player.id).then(() => {
      setGame({} as Game);
      setRoom({} as Room);
    });
  };

  const handleStartClick = () => {
    homepageService.startRoom(room.id, defaultClockMillis).then((gameResponse) => {
      homepageService.subscribeToGameUpdates(gameResponse.id);
      setGame({ ...game, ...gameResponse });
    });
  };

  const handleConcedeClick = () => {
    if (!game.id) return;
    homepageService.concedeGame(game.id, player.id).then((gameResponse) => {
      setGame({ ...game, ...gameResponse });
    });
  };

  const TopButtons = () => {
    if (game.id) {
      return (
        <TopButtonContainer>
          <GradientButton className="drawButton">Offer Draw</GradientButton>
          <GradientButton className="concedeButton" onClick={() => handleConcedeClick()}>
            Concede
          </GradientButton>
        </TopButtonContainer>
      );
    } else {
      return (
        <GradientButton className="startButton" onClick={() => handleStartClick()}>
          Start
        </GradientButton>
      );
    }
  };

  return (
    <ButtonContainer>
      <TopButtons></TopButtons>
      <Link to={`/home`}>
        <GradientButton style={{ width: '100%' }} onClick={() => handleLeaveClick()}>
          Leave
        </GradientButton>
      </Link>
    </ButtonContainer>
  );
}

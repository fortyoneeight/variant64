import React, { useMemo, useEffect } from 'react';
import { useRecoilState } from 'recoil';
import { playerState, gameState } from '../../../store/atoms';
import { ServicesContext } from '../../../store/context';
import { HomepageService } from '../../pages/hompage-service';
import './DrawButton.css';

export default function DrawButton() {
  const [player, _] = useRecoilState(playerState);
  const [game, setGame] = useRecoilState(gameState);

  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  const handleDrawClick = (accepted: boolean) => {
    homepageService.drawGame(accepted, game.id, player.id).then((gameResponse) => {
      setGame({
        ...game,
        ...gameResponse,
      });
    });
  };

  function getDrawState() {
    const drawStates = game.approved_draw_players;

    if (!drawStates) return 'game_not_started';
    if (drawStates[player.id]) return 'outgoing_draw';

    for (const id in drawStates) {
      if (drawStates[id]) {
        return 'incoming_draw';
      }
    }
    return 'no_pending_draw';
  }

  function getDrawProgress() {
    let num_approved = 0;
    let num_total = 0;
    for (const id in game.approved_draw_players) {
      num_total += 1;
      if (game.approved_draw_players[id]) {
        num_approved += 1;
      }
    }
    return num_approved + '/' + num_total;
  }

  const StatefulDrawButton = () => {
    switch (getDrawState()) {
      case 'game_not_started': {
        return <></>;
      }
      case 'outgoing_draw': {
        return (
          <button className="drawButton gradientBackground" disabled>
            Draw Pending <br />
            {getDrawProgress()} Accepted
          </button>
        );
      }
      case 'no_pending_draw': {
        return (
          <button
            className="drawButton gradientBackground"
            onClick={() => {
              handleDrawClick(true);
            }}
          >
            Offer Draw
          </button>
        );
      }
      case 'incoming_draw': {
        return (
          <div className="drawButton gradientBackground">
            Draw? <br />
            {getDrawProgress()} Accepted
            <br />
            <button
              className="whiteButton"
              style={{ margin: '0.5rem' }}
              onClick={() => {
                handleDrawClick(true);
              }}
            >
              Yes
            </button>
            <button
              className="whiteButton"
              onClick={() => {
                handleDrawClick(false);
              }}
            >
              No
            </button>
          </div>
        );
      }
    }
  };

  return <StatefulDrawButton />;
}

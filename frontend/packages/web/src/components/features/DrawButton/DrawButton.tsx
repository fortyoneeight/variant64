import React, { useMemo } from 'react';
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

  const handleDrawOnClick = (accepted: boolean) => {
    homepageService.drawGame(accepted, game.id, player.id).then((gameResponse) => {
      setGame({ ...game, ...gameResponse });
    });
  };

  enum drawState {
    NONE = 'none',
    INCOMING_DRAW = 'incoming_draw',
    OUTGOING_DRAW = 'outgoing_draw',
  }

  function getDrawState() {
    if (!game.approved_draw_players) return drawState.NONE;
    if (game.approved_draw_players[player.id]) return drawState.OUTGOING_DRAW;
    if (Object.values(game.approved_draw_players).some((value) => value)) {
      return drawState.INCOMING_DRAW;
    }
    return drawState.NONE;
  }

  function getDrawProgress() {
    let approvals = Object.values(game.approved_draw_players);
    return approvals.filter((val) => val).length + '/' + approvals.length;
  }

  const StatefulDrawButton = () => {
    switch (getDrawState()) {
      case drawState.NONE: {
        return (
          <button
            className="drawButton gradientBackground"
            onClick={() => {
              handleDrawOnClick(true);
            }}
          >
            Offer Draw
          </button>
        );
      }
      case drawState.OUTGOING_DRAW: {
        return (
          <button className="drawButton gradientBackground" disabled>
            Draw Pending <br />
            {getDrawProgress()} Accepted
          </button>
        );
      }
      case drawState.INCOMING_DRAW: {
        return (
          <div className="drawButton gradientBackground">
            Draw? <br />
            {getDrawProgress()} Accepted
            <br />
            <button
              className="whiteButton"
              style={{ margin: '0.5rem' }}
              onClick={() => {
                handleDrawOnClick(true);
              }}
            >
              Yes
            </button>
            <button
              className="whiteButton"
              onClick={() => {
                handleDrawOnClick(false);
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

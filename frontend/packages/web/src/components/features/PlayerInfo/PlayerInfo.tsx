import React from 'react';
import { useNavigate } from 'react-router-dom';
import { BoardCellItemTypes, Player } from '../../../models';
import { useRecoilState } from 'recoil';
import { playerState, gameState, roomState } from '../../../store/atoms';
import { PlayerInfoCard, PlayerInfoContainer } from './PlayerInfo.styled';
import { pieceImages } from '../../../assets/pieceImages';

export default function PlayerInfo() {
  const [game, setGame] = useRecoilState(gameState);
  const [room, setRoom] = useRecoilState(roomState);

  const PlayerCard = (props: any) => {
    return (
      <PlayerInfoCard>
        <p>
          Temp
          <br />
          10:00{' '}
        </p>
        <img src={pieceImages['king'][props.id == 0 ? 'white' : 'black']} />
      </PlayerInfoCard>
    );
  };
  return (
    <PlayerInfoContainer>
      {room.players.map((player, i) => {
        return <PlayerCard key={i} id={i} />;
      })}
    </PlayerInfoContainer>
  );
}

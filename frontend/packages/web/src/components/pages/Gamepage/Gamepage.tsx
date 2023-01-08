import React from 'react';
import { mockdataBoard } from '../../../store/mockdata';
import { BoardDisplay, GameplayButtons, PlayerInfo } from '../../features';
import { GameContainer } from './Gamepage.styled';

export default function Gamepage() {
  return (
    <GameContainer>
      <PlayerInfo />
      <BoardDisplay board={mockdataBoard} />
      <GameplayButtons />
    </GameContainer>
  );
}

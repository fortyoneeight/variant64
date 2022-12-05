import React from 'react';
import { useRecoilState } from 'recoil';
import { playerState } from '../../../store/atoms';
import { RoomForm, RoomList } from '../../features';
import './Homepage.css';

export default function Homepage() {
  const [player, _] = useRecoilState(playerState);

  return (
    <div className="column">
      <h1 className="welcomeHeader">Welcome: {player.display_name}</h1>
      <div className="row">
        <RoomList />
        <RoomForm />
      </div>
    </div>
  );
}

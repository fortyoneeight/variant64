import React from 'react';
import { useRecoilState } from 'recoil';
import { playerState } from '../../../store/atoms';
import { CreateRoomForm, RoomList } from '../../features';
import { FlexRow, FlexColumn } from './Homepage.styled';

export default function Homepage() {
  const [player, _] = useRecoilState(playerState);

  return (
    <FlexColumn>
      <h1 className="welcomeHeader">Welcome: {player.display_name}</h1>
      <FlexRow>
        <RoomList />
        <CreateRoomForm />
      </FlexRow>
    </FlexColumn>
  );
}

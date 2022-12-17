import React, { useState, useMemo, useEffect } from 'react';
import { Room } from '../../../models';
import { useRecoilState } from 'recoil';
import { roomsState, roomState } from '../../../store/atoms';
import { ServicesContext } from '../../../store/context';
import { HomepageService } from '../../pages/hompage-service';
import { WhiteButton, TextInput, WhiteBackgroundText, Dropdown } from '../../atoms/';
import { CreateRoomFormContainer, SettingsContainer } from './CreateRoomForm.styled';

export default function CreateRoomForm() {
  const [roomForm, setRoomForm] = useState<Room>({} as Room);
  const [rooms, setRooms] = useRecoilState(roomsState);

  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  useEffect(() => {
    homepageService.getRooms().then((rooms) => setRooms(rooms));
  }, []);

  const handleRoomSubmit = (homepageService: HomepageService, roomName: string) => {
    homepageService.createRoom(roomName).then((roomResponse) => {
      setRooms([...rooms, roomResponse]);
    });
  };

  const handleOnchange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const [_, name] = e.target.name.split('roomForm_');
    setRoomForm({
      ...roomForm,
      [name]: e.target.value,
    });
  };

  const handleOnSubmit = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();
    if (roomForm.name === '') {
      return;
    }
    let roomName = document.querySelector('#roomForm_name') as any;
    if (roomName) {
      roomName.value = '';
    }
    handleRoomSubmit(homepageService, roomForm.name);
  };

  return (
    <CreateRoomFormContainer>
      <h1 style={{ color: '#ffffff' }}>Create a Room</h1>
      <SettingsContainer>
        <WhiteBackgroundText>Room Name:</WhiteBackgroundText>
        <TextInput
          id="roomForm_name"
          title="roomForm_name"
          name="roomForm_name"
          placeholder="Name"
          onChange={(e) => handleOnchange(e)}
        ></TextInput>

        <WhiteBackgroundText style={{ background: '#FFFF00' }}>Password:</WhiteBackgroundText>
        <TextInput style={{ background: '#FFFF00' }} placeholder="(Optional)"></TextInput>

        <WhiteBackgroundText style={{ background: '#FFFF00' }}>Gamemode:</WhiteBackgroundText>
        <Dropdown style={{ background: '#FFFF00' }}>
          <option value="classic">Classic</option>
          <option value="small">Small</option>
          <option value="fourPlayer">Four Player</option>
          <option value="guards">Guards</option>
          <option value="theMine">The Mine</option>
        </Dropdown>

        <WhiteBackgroundText style={{ background: '#FFFF00' }}>Clock Time:</WhiteBackgroundText>
        <Dropdown style={{ background: '#FFFF00' }}>
          <option value="600000">10 Minute</option>
          <option value="300000">5 Minute</option>
          <option value="120000">2 Minute</option>
          <option value="1800000">30 Minute</option>
          <option value="60000">60 Seconds</option>
        </Dropdown>
      </SettingsContainer>
      <WhiteButton onClick={(e) => handleOnSubmit(e)}>Submit</WhiteButton>{' '}
    </CreateRoomFormContainer>
  );
}

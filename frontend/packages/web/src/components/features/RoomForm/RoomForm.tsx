import React, { useState, useMemo, useEffect } from 'react';
import { Room } from '../../../models';
import { useRecoilState } from 'recoil';
import { roomsState, roomState } from '../../../store/atoms';
import { ServicesContext } from '../../../store/context';
import { HomepageService } from '../../pages/hompage-service';
import './RoomForm.css';

export default function RoomForm() {
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
    <form className="roomFormContainer gradientBackground" action="">
      <h1 style={{ color: '#ffffff' }}>Create a Room</h1>
      <div className="formContainer">
        <div className="formRow">
          <p className="roomFormTextField">Room Name:</p>
          <input
            id="roomForm_name"
            title="roomForm_name"
            name="roomForm_name"
            className="roomFormTextField"
            placeholder="Room Name"
            onChange={(e) => handleOnchange(e)}
          ></input>
        </div>
        <div className="formRow">
          <p className="roomFormTextField" style={{ background: '#FFFF00' }}>
            Password:
          </p>
          <input
            className="roomFormTextField"
            style={{ background: '#FFFF00' }}
            placeholder="Password"
          ></input>
        </div>
        <div className="formRow">
          <p className="roomFormTextField" style={{ background: '#FFFF00' }}>
            Gamemode:
          </p>
          <p className="roomFormTextField" style={{ background: '#FFFF00' }}>
            Dropdown
          </p>
        </div>
        <div className="formRow">
          <p className="roomFormTextField" style={{ background: '#FFFF00' }}>
            Clock Time:
          </p>
          <p className="roomFormTextField" style={{ background: '#FFFF00' }}>
            Dropdown
          </p>
        </div>
      </div>
      <button className="submitButton" onClick={(e) => handleOnSubmit(e)}>
        Submit
      </button>{' '}
    </form>
  );
}

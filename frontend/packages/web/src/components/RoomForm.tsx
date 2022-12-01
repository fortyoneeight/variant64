import React, { useState } from 'react';
import { Room } from '../models';

export default function RoomForm({ submitForm }: { submitForm: any }) {
  const [form, setForm] = useState<Room>({} as Room);

  const handleOnchange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const [_, name] = e.target.name.split('roomForm_');
    setForm({
      ...form,
      [name]: e.target.value,
    });
  };

  const handleOnSubmit = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();
    submitForm(form.name);
    let roomName = document.querySelector('#roomForm_name') as any;
    if (roomName) {
      roomName.value = '';
    }
  };

  return (
    <form action="">
      <label htmlFor="roomForm_name">
        Room name <br />
        <input
          id="roomForm_name"
          title="roomForm_name"
          name="roomForm_name"
          placeholder="Room Name"
          onChange={(e) => handleOnchange(e)}
        ></input>
        <button onClick={(e) => handleOnSubmit(e)}>Create Room</button>
      </label>
    </form>
  );
}

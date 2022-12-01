import React, { useState } from 'react';
import { Player } from '../models';

export default function PlayerForm({ submitForm }: { submitForm: any }) {
  const [form, setForm] = useState<Player>({} as Player);

  const handleOnchange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const [_, name] = e.target.name.split('playerForm_');
    setForm({
      ...form,
      [name]: e.target.value,
    });
  };

  const handleOnSubmit = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();
    submitForm(form.display_name);
    setForm({} as Player);
  };

  return (
    <form>
      <label htmlFor="playerForm_display_name">
        Player display name <br />
        <input
          id="playerForm_display_name"
          title="playerForm_display_name"
          name="playerForm_display_name"
          placeholder="Display Name"
          onChange={(e) => handleOnchange(e)}
        ></input>
        <button onClick={(e) => handleOnSubmit(e)}>Create Player</button>
      </label>
    </form>
  );
}

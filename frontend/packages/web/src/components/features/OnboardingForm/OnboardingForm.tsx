import React, { useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Player } from '../../../models';
import { useRecoilState } from 'recoil';
import { playerState } from '../../../store/atoms';
import { HomepageService } from '../../pages/hompage-service';
import { ServicesContext } from '../../../store/context';
import { WhiteButton, TextInput } from '../../atoms';
import { OnboardingFormContainer } from './OnboardingForm.styled';

export default function OnboardingForm() {
  const [form, setForm] = useState<Player>({
    display_name: localStorage.getItem('playerName')!,
  } as Player);
  const [player, setPlayer] = useRecoilState(playerState);

  const navigate = useNavigate();
  const context = React.useContext(ServicesContext);
  const homepageService = useMemo(
    () => new HomepageService(context.roomHttpService, context.roomWebSocketService),
    [context.roomHttpService, context.roomWebSocketService]
  );

  const handleOnchange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const [_, name] = e.target.name.split('playerForm_');
    setForm({
      ...form,
      [name]: e.target.value,
    });
  };

  const handleOnSubmit = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();
    homepageService.createPlayer(form.display_name).then((playerResponse) => {
      setPlayer({ ...player, ...playerResponse });
      navigate('/home');
      localStorage.setItem('playerName', playerResponse.display_name);
    });
    setForm({} as Player);
  };

  return (
    <OnboardingFormContainer>
      <TextInput
        id="playerForm_display_name"
        title="playerForm_display_name"
        name="playerForm_display_name"
        placeholder="Enter Name"
        defaultValue={localStorage.getItem('playerName')!}
        onChange={(e) => handleOnchange(e)}
      ></TextInput>
      <WhiteButton onClick={(e) => handleOnSubmit(e)}>Submit</WhiteButton>
    </OnboardingFormContainer>
  );
}

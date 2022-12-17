import React from 'react';
import styled from 'styled-components';

export const CreateRoomFormContainer = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  gap: 1vh;
  padding: 4vh;
  border-radius: 2em;
  min-height: 60vh;
  min-width: 35vw;
  background-image: linear-gradient(to bottom right, #097500, #2dff74);
  box-shadow: inset 0.5em 0.5em 0.5em 0em #57b45e, inset -0.5em -0.5em 0.5em 0em #1ba23f;
`;

export const SettingsContainer = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-gap: 1vh;
  width: 100%;
`;

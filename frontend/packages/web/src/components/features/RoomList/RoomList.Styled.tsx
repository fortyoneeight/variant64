import React from 'react';
import styled from 'styled-components';

export const RoomListContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 4vh;
  gap: 3vh;
  border-radius: 2em;
  min-height: 60vh;
  min-width: 35vw;
  background-image: linear-gradient(to bottom right, #097500, #2dff74);
  box-shadow: inset 0.5em 0.5em 0.5em 0em #57b45e, inset -0.5em -0.5em 0.5em 0em #1ba23f;
`;

export const ListContainer = styled.div`
  display: flex;
  flex-direction: column;
  gap: 1vh;
  min-width: 35vw;
`;

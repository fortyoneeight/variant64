import React from 'react';
import styled from 'styled-components';

export const PlayerInfoCard = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.2em 0.8em;
  border: none;
  border-radius: 1em;
  line-height: 1em;
  font-size: 1.5em;
  color: white;
  height: 2.5em;
  background-image: linear-gradient(to bottom right, #097500, #2dff74);
  box-shadow: inset 0.5em 0.5em 0.5em 0em #57b45e, inset -0.5em -0.5em 0.5em 0em #1ba23f;
`;

export const PlayerInfoContainer = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  width: 12rem;
`;

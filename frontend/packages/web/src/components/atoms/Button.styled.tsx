import React from 'react';
import styled from 'styled-components';

const BasicButton = styled.button`
  font-size: 1.5em;
  padding: 0.3em 2em;
  border: none;
  border-radius: 0.5em;
  width: auto;
`;

export const WhiteButton = styled(BasicButton)`
  background-color: #fff;
  &:hover {
    background-color: rgb(202, 199, 199);
  }
  &:active {
    background-color: #2b8024;
  }
`;

export const GradientButton = styled(BasicButton)`
  color: #fff;
  background-image: linear-gradient(to bottom right, #097500, #2dff74);
  box-shadow: inset 0.5em 0.5em 0.5em 0em #57b45e, inset -0.5em -0.5em 0.5em 0em #1ba23f;
  &:hover {
    filter: brightness(1.2);
  }
  &:active {
    filter: brightness(1.4);
  }
`;

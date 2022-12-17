import React from 'react';
import styled from 'styled-components';

const BasicText = styled.p`
  font-size: 1.5em;
  padding: 0.2em;
  border: none;
  border-radius: 0.5em;
  width: auto;
`;

export const WhiteBackgroundText = styled(BasicText)`
  color: black;
  background-color: #fff;
`;

export const GradientBackgroundText = styled(BasicText)`
  color: #fff;
  text-align: left;
  background-image: linear-gradient(to bottom right, #097500, #2dff74);
  box-shadow: inset 0.5em 0.5em 0.5em 0em #57b45e, inset -0.5em -0.5em 0.5em 0em #1ba23f;
`;

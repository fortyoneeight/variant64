import React from 'react';
import styled from 'styled-components';

interface Boundaries {
  rows: number;
  columns: number;
}

export const GameboardGrid = styled.div<Boundaries>`
  display: grid;
  grid-gap: 0.3rem;
  background-color: black;
  border: 0.3rem solid black;
  grid-template-rows: repeat(${(props) => props.rows}, ${(props) => 75 / props.columns + 'vh'});
  grid-template-columns: repeat(
    ${(props) => props.columns},
    ${(props) => 75 / props.columns + 'vh'}
  );
`;

export const GridCell = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
`;

export const PieceImage = styled.img`
  width: 100%;
  height: 100%;
`;

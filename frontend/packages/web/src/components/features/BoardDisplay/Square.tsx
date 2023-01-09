import React from 'react';
import { useDrop } from 'react-dnd';
import { GridCell } from './BoardDisplay.styled';

export default function Square({ children, i, j, color }: any) {
  const [{ isOver, canDrop }, drop] = useDrop(
    () => ({
      accept: 'piece',
      canDrop: () => true,
      collect: (monitor) => ({
        isOver: !!monitor.isOver(),
        canDrop: !!monitor.canDrop(),
      }),
    }),
    []
  );

  return (
    <GridCell ref={drop} key={i + '' + j} style={{ backgroundColor: isOver ? 'red' : color }}>
      {children}
    </GridCell>
  );
}

import React from 'react';
import { useDrag } from 'react-dnd';
import { pieceImages } from '../../../assets/pieceImages';
import { PieceImage } from './BoardDisplay.styled';

export default function Piece({ piece }: any) {
  const [collected, drag, dragPreview] = useDrag(() => ({
    previewOptions: {
      offsetX: 30,
      offsetY: 30,
    },
    type: 'piece',
  }));

  const pieceImageSource =
    pieceImages[piece.name as keyof typeof pieceImages][
      (piece.player == 'player1' ? 'white' : 'black')!
    ];

  const image = new Image();
  image.src = pieceImageSource;
  dragPreview(image);

  return (
    <div ref={drag}>
      <PieceImage src={pieceImageSource} />
    </div>
  );
}

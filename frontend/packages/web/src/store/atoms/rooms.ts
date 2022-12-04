import { atom } from 'recoil';
import { Player, Room, Game } from '../../models';

export const roomsState = atom({
  key: 'roomsState',
  default: [] as Array<Room>,
});

export const roomState = atom({
  key: 'roomState',
  default: {} as Room,
});

export const playerState = atom<Player>({
  key: 'playerState',
  default: {} as Player,
});

export const gameState = atom<Game>({
  key: 'gameState',
  default: {} as Game,
});

export const gameUpdateState = atom<{
  game_id: string;
  clocks: {
    [key: string]: string;
  };
}>({
  key: 'gameUpdateState',
  default: {} as {
    game_id: string;
    clocks: {
      [key: string]: string;
    };
  },
});

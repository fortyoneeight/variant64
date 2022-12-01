import { atom } from 'recoil';
import { Player, Room } from '../../models';

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

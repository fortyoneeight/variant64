import { atom } from 'recoil';

export const roomSubState = atom<string>({
  key: 'room-sub-state',
  default: '',
});

export const gameSubState = atom<string>({
  key: 'game-sub-state',
  default: '',
});

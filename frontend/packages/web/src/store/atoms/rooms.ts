import { atom } from 'recoil';
import { mockdataRooms } from '../../store/mockdata';

export const roomsState = atom({
  key: 'roomsState',
  default: mockdataRooms,
});

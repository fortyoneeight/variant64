import { atom } from "recoil";
import { mockdataRooms } from "../../mockdata/rooms";

export const roomsState = atom({
    key: 'roomsState',
    default: mockdataRooms
})

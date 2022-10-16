import { BoardPlayer } from "./board";

export type Room = {
  room_name: string;
  players_total: number;
  players: Array<BoardPlayer>;
};

export type CreateRoomRequest = {
  room_name: string;
};

export type CreateRoomResponse = {} & Room;

export type GetRoomsRequest = {};

export type GetRoomsResponse = {
  rooms: Array<Room>;
};

export type JoinRoomRequest = {
  room_name: string;
  player_name: string;
};

export type JoinRoomResponse = {} & Room;
export type JoinRoomErrorResponse = {
  error: string;
};

export type StartRoomRequest = {
  room_name: string;
};
export type StartRoomResponse = {};
export type StartRoomErrorResponse = {
  error: string;
};

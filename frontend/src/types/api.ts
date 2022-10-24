import { BoardPlayer } from './board';

// TODO: move this to another file
export type Room = {
  room_name: string;
  players_total: number;
  players: Array<BoardPlayer>;
};

export type CreateRoomRequest = {
  room_name: string;
};

export type CreateRoomResponse = {} & Room;

export type GetRoomsRequest = Array<Room>;

export type GetRoomsResponse = Array<Room>;

export type JoinRoomRequest = {
  player_id: string;
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

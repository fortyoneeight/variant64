import { Player, Room } from '../../models';
import { RoutesParams } from '../sharedTypes';

export interface HttpRequestEvent {
  action: HTTPActions;
  params?: any;
  body?: any;
}

export interface RoutesConfig {
  name: string;
  routes: Record<HTTPActions, RouteConfig>;
}

export interface HttpServiceParams {
  url: string;
  routesConfig: RoutesConfig;
}

export enum HTTPActions {
  // rooms
  CREATE_ROOM = 'CREATE_ROOM',
  GET_ROOM = 'GET_ROOM',
  JOIN_ROOM = 'JOIN_ROOM',
  LEAVE_ROOM = 'LEAVE_ROOM',
  START_ROOM = 'START_ROOM',
  GET_ROOMS = 'GET_ROOMS',

  // player
  CREATE_PLAYER = 'CREATE_PLAYER',
  GET_PLAYER = 'GET_PLAYER',
}

export type CreateRoomRequest = {
  [RoutesParams.ROOM_NAME]: string;
};

export type CreateRoomResponse = {} & Room;

export type GetRoomRequest = {
  [RoutesParams.ROOM_ID]: string;
};
export type GetRoomResponse = Room;

export type GetRoomsRequest = Array<Room>;

export type GetRoomsResponse = Array<Room>;

export type JoinRoomParams = {
  [RoutesParams.ROOM_NAME]: string;
};

export type JoinRoomRequest = {
  [RoutesParams.PLAYER_ID]: string;
};

export type JoinRoomResponse = {} & Room;
export type JoinRoomErrorResponse = {
  error: string;
};

export type LeaveRoomRequest = {
  [RoutesParams.PLAYER_ID]: string;
  [RoutesParams.ROOM_ID]: string;
};
export type LeaveRoomResponse = {} & Room;

export type StartRoomRequest = {
  [RoutesParams.ROOM_NAME]: string;
};
export type StartRoomResponse = {};
export type StartRoomErrorResponse = {
  error: string;
};

export type CreatePlayerRequest = {
  [RoutesParams.PLAYER_ID]: string;
};
export type CreatePlayerResponse = Player;

export type GetPlayerRequest = {
  [RoutesParams.PLAYER_ID]: string;
};

export type GetPlayerResponse = Player;
export interface RouteConfig {
  path: (id?: string) => string;
  method: string;
}

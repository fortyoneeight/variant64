import { Player } from './board';
import { AppActions } from './actions';

// TODO: move this to another file
export type Room = {
  id: string;
  name: string;
  players: Array<string>;
};

export enum RoutesParams {
  ROOM_NAME = 'room_name',
  ROOM_ID = 'room_id',
  PLAYER_ID = 'player_id',
  PLAYER_DISPLAY_NAME = 'display_name',
  PLAYER_TIME_MILLIS = 'player_time_ms',
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

export function NewRoute(path: (params?: any) => string, method: string) {
  return {
    path: path,
    method,
  };
}

export interface RoutesConfig {
  name: string;
  routes: Record<AppActions, RouteConfig>;
}

export interface HttpServiceParams {
  url: string;
  routesConfig: RoutesConfig;
}

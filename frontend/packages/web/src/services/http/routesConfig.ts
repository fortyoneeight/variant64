import { RoutesParams } from '../sharedTypes';
import { HTTPActions, RouteConfig } from './types';

export function NewRoute(path: (params?: any) => string, method: string) {
  return {
    path: path,
    method,
  };
}

export const RoomAPIRoutesConfig = {
  name: 'ROOM_API',
  routes: {
    // rooms
    [HTTPActions.CREATE_ROOM]: NewRoute(() => '/room', 'post'),
    [HTTPActions.GET_ROOM]: NewRoute((params) => `/room/${params[RoutesParams.ROOM_ID]}`, 'get'),
    [HTTPActions.JOIN_ROOM]: NewRoute(
      (params) => `/room/${params[RoutesParams.ROOM_ID]}/join`,
      'post'
    ),
    [HTTPActions.LEAVE_ROOM]: NewRoute(
      (params) => `/room/${params[RoutesParams.ROOM_ID]}/leave`,
      'post'
    ),
    [HTTPActions.GET_ROOMS]: NewRoute(() => '/rooms', 'get'),

    // player
    [HTTPActions.CREATE_PLAYER]: NewRoute(() => '/player', 'post'),
    [HTTPActions.GET_PLAYER]: NewRoute(
      (params) => `/player/${params[RoutesParams.PLAYER_ID]}`,
      'get'
    ),
    [HTTPActions.START_ROOM]: NewRoute(
      (params) => `/room/${params[RoutesParams.ROOM_ID]}/start`,
      'post'
    ),
  } as Record<HTTPActions, RouteConfig>,
};

export type RoomAPIRoutesConfigType = typeof RoomAPIRoutesConfig;

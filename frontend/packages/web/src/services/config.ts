import { AppActions, NewRoute, RoutesParams, RouteConfig } from '../types';

export const RoomAPIRoutesConfig = {
  name: 'ROOM_API',
  routes: {
    // rooms
    [AppActions.CREATE_ROOM]: NewRoute(() => '/room', 'post'),
    [AppActions.GET_ROOM]: NewRoute((params) => `/room/${params[RoutesParams.ROOM_ID]}`, 'get'),
    [AppActions.JOIN_ROOM]: NewRoute(
      (params) => `/room/${params[RoutesParams.ROOM_ID]}/join`,
      'post'
    ),
    [AppActions.LEAVE_ROOM]: NewRoute(
      (params) => `/room/${params[RoutesParams.ROOM_ID]}/leave`,
      'post'
    ),
    [AppActions.START_ROOM]: NewRoute(
      (params) => `/room/${params[RoutesParams.ROOM_ID]}/start`,
      'post'
    ),
    [AppActions.GET_ROOMS]: NewRoute(() => '/rooms', 'get'),

    // player
    [AppActions.CREATE_PLAYER]: NewRoute(() => '/player', 'post'),
    [AppActions.GET_PLAYER]: NewRoute(
      (params) => `/player/${params[RoutesParams.PLAYER_ID]}`,
      'get'
    ),
  } as Record<AppActions, RouteConfig>,
};

export type RoomAPIRoutesConfigType = typeof RoomAPIRoutesConfig;

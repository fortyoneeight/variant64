import { RoutesParams } from '../sharedTypes';
export interface WebSocketServiceParams {
  url: string;
}

export interface WebSocketRequestEvent {
  action: WebSocketActions;
  body?: any;
}

export enum WebSocketActions {
  SUBSCRIBE_GAME_UPDATES = 'game_subscribe',
  SUBSCRIBE_ROOM_UPDATES = 'room_subscribe',
}

export interface SubscribeGameUpdatesCommand {
  [RoutesParams.GAME_ID]: string;
}

export interface SubscribeRoomUpdatesCommand {
  [RoutesParams.ROOM_ID]: string;
}

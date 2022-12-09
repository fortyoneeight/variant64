import { RoutesParams } from '../sharedTypes';
export interface WebSocketServiceParams {
  url: string;
}

export interface WebSocketRequestEvent {
  action: WebSocketActions;
  channel: WebSocketChannels;
  body?: any;
}

export enum WebSocketChannels {
  GAME = 'game',
  ROOM = 'room',
}

export enum WebSocketActions {
  SUBSCRIBE_GAME_UPDATES = 'subscribe',
  SUBSCRIBE_ROOM_UPDATES = 'subscribe',
}

export interface SubscribeGameUpdatesCommand {
  [RoutesParams.GAME_ID]: string;
}

export interface SubscribeRoomUpdatesCommand {
  [RoutesParams.ROOM_ID]: string;
}

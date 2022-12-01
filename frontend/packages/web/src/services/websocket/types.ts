import { RoutesParams } from '../sharedTypes';
export interface WebSocketServiceParams {
  url: string;
}

export interface WebSocketRequestEvent {
  action: WebSocketActions;
  body?: any;
}

export enum WebSocketActions {
  SUBSCRIBE_GAME_UPDATES = 'subscribe',
}

export interface SubscribeGameUpdatesCommand {
  [RoutesParams.GAME_ID]: string;
}

import { HTTPActions } from './http';
import { WebSocketActions } from './websocket';

export * from './http';
export * from './websocket';
export * from './sharedTypes';

export const AppActions = { ...HTTPActions, ...WebSocketActions };
export type AppActions = typeof AppActions;

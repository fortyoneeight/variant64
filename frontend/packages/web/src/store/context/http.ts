import React from 'react';
import { HttpService } from '../../services';
import { WebSocketService } from '../../services/websocket/websocket';

export const ServicesContext = React.createContext({
  roomHttpService: {} as HttpService,
  roomWebSocketService: {} as WebSocketService,
});

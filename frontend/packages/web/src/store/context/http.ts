import React from 'react';
import { HttpService } from '../../services';
import { WebSocketService } from '../../services/websocket';

export const ServicesContext = React.createContext({
  roomHttpService: {} as HttpService,
  roomWebSocketService: {} as WebSocketService,
});

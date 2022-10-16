import React from 'react';
import { RoomHttpService } from '../../services';

export const HttpContext = React.createContext({
  roomService: {} as RoomHttpService,
});

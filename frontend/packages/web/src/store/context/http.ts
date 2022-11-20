import React from 'react';
import { HttpService } from '../../services';

export const HttpContext = React.createContext({
  roomService: {} as HttpService,
});

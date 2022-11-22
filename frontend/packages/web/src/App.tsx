import React from 'react';
import './App.css';
import Routing from './Routing';
import { RecoilRoot } from 'recoil';
import { ServicesContext } from './store/context';
import { RoomAPIRoutesConfig, HttpService, WebSocketService } from './services';

let roomHttpService = new HttpService({
  url: 'http://0.0.0.0:8001/0.0.0.0:8000/api',
  routesConfig: RoomAPIRoutesConfig,
});

let roomWebSocketService = new WebSocketService({
  url: 'ws://0.0.0.0:8000/ws',
  routesConfig: RoomAPIRoutesConfig,
});

function App() {
  return (
    <>
      <RecoilRoot>
        <ServicesContext.Provider
          value={{ roomHttpService: roomHttpService, roomWebSocketService: roomWebSocketService }}
        >
          <div>
            <Routing />
          </div>
        </ServicesContext.Provider>
      </RecoilRoot>
    </>
  );
}

export default App;

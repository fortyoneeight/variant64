import React from 'react';
import './App.css';
import Routing from './Routing';
import { RecoilRoot } from 'recoil';
import { ServicesContext } from './store/context';
import { RoomAPIRoutesConfig, HttpService, WebSocketService } from './services';

const { REACT_APP_BE_SEVER, REACT_APP_DEVELOPMENT_PROXY } = process.env;
const proxyUrl = REACT_APP_DEVELOPMENT_PROXY ? REACT_APP_DEVELOPMENT_PROXY : '';
console.log(REACT_APP_BE_SEVER, REACT_APP_DEVELOPMENT_PROXY, process.env.NODE_ENV);

let roomHttpService = new HttpService({
  url: `http://${proxyUrl}/${REACT_APP_BE_SEVER}/api`,
  routesConfig: RoomAPIRoutesConfig,
});

let roomWebSocketService = new WebSocketService({
  url: `ws://${REACT_APP_BE_SEVER}/ws`,
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

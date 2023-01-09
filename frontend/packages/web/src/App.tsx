import React from 'react';
import Routing from './Routing';
import { RecoilRoot } from 'recoil';
import { ServicesContext } from './store/context';
import { RoomAPIRoutesConfig, HttpService, WebSocketService } from './services';
import { GlobalStyles } from './GlobalStyles.styled';
import { DndProvider } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';

const { REACT_APP_BE_SERVER, REACT_APP_DEVELOPMENT_PROXY } = process.env;
const proxyUrl = REACT_APP_DEVELOPMENT_PROXY ? REACT_APP_DEVELOPMENT_PROXY : '';
console.log(REACT_APP_BE_SERVER, REACT_APP_DEVELOPMENT_PROXY, process.env.NODE_ENV);

let roomHttpService = new HttpService({
  url: `http://${proxyUrl}/${REACT_APP_BE_SERVER}/api`,
  routesConfig: RoomAPIRoutesConfig,
});

let roomWebSocketService = new WebSocketService({
  url: `ws://${REACT_APP_BE_SERVER}/ws`,
});

function App() {
  return (
    <>
      <GlobalStyles />
      <RecoilRoot>
        <DndProvider backend={HTML5Backend}>
          <ServicesContext.Provider
            value={{ roomHttpService: roomHttpService, roomWebSocketService: roomWebSocketService }}
          >
            <div className="appBody">
              <Routing />
            </div>
          </ServicesContext.Provider>
        </DndProvider>
      </RecoilRoot>
    </>
  );
}

export default App;

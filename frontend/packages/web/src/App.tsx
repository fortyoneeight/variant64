import React from 'react';
import './App.css';
import Routing from './Routing';
import { RecoilRoot } from 'recoil';
import { HttpContext } from './store/context';
import { RoomAPIRoutesConfig } from './services';
import { HttpService } from './services/http';

let roomHttpService = new HttpService({
  url: 'http://0.0.0.0:8001/0.0.0.0:8000/api',
  routesConfig: RoomAPIRoutesConfig,
});

function App() {
  return (
    <>
      <RecoilRoot>
        <HttpContext.Provider value={{ roomService: roomHttpService }}>
          <div className="grid-3-vertical-centerbias">
            <Routing />
          </div>
        </HttpContext.Provider>
      </RecoilRoot>
    </>
  );
}

export default App;

import React from 'react';
import logo from './logo.svg';
import './App.css';
import Routing from './Routing';
import { Link } from 'react-router-dom';
import { RecoilRoot } from 'recoil';
import { RoomHttpService } from './services';
import { HttpContext } from './store/context';

let roomHttpService = new RoomHttpService({
  // url: 'http://localhost:8000',
  url: 'http://0.0.0.0:8001/0.0.0.0:8000',
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

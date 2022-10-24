import React from 'react';
import logo from './logo.svg';
import './App.css';
import Routing from './Routing';
import { Link } from 'react-router-dom';
import { RecoilRoot } from 'recoil';
import { RoomHttpService } from './services';
import { HttpContext } from './store/context';

let roomHttpService = new RoomHttpService({
  // url: 'https://3fa5-24-142-141-179.ngrok.io',
  url: 'http://localhost:8000',
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

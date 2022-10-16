import React from 'react';
import logo from './logo.svg';
import './App.css';
import Routing from './Routing';
import { RoomHttpService } from './services';
import { HttpContext } from './store/context';

let roomHttpService = new RoomHttpService({
  // url: 'https://1dda-24-142-141-179.ngrok.io/',
  url: 'http://localhost:8000/',
});

function App() {
  return (
    <>
      <HttpContext.Provider value={{ roomService: roomHttpService }}>
        <div className='grid-3-vertical-centerbias'>
          <Routing />
        </div>
      </HttpContext.Provider>
    </>
  );
}

export default App;

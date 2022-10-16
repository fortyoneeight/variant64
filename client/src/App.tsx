import React from 'react';
import logo from './logo.svg';
import './App.css';
import Routing from './Routing';
import { Link } from "react-router-dom";
import { RecoilRoot } from 'recoil';
function App() {
  return <>
    <RecoilRoot>
      <div className='grid-3-vertical-centerbias'>
        <Routing />
      </div>
    </RecoilRoot>
  </>
}

export default App;

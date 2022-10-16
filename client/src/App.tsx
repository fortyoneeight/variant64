import React from 'react';
import logo from './logo.svg';
import './App.css';
import Routing from './Routing';
import { Link } from "react-router-dom";

function Header() {
  return (
    <a href='/'>
      <h1 style={{ textAlign: 'center' }}>
        Variant64 Chess
      </h1>
    </a>
  )
}

function Footer() {
  return <aside>
    <a href='https://github.com/izakfr/variant64'><small>Variant64 - Github Repo</small></a>
  </aside>
}


function App() {
  return <>

    <div className='grid-3-vertical-centerbias'>
      <Header />

      <Routing />

      <Footer />
    </div>
  </>
}

export default App;

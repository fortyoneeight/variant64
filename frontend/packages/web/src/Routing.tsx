import React from 'react';
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom';
import Gamepage from './pages/Gamepage';
import Homepage from './pages/Homepage';

function Header() {
  return (
    <nav>
      <h1 style={{ textAlign: 'center' }}>
        <Link to="/">Variant64 Chess</Link>
      </h1>
    </nav>
  );
}

function Footer() {
  return (
    <div>
      <aside>
        <p style={{ textAlign: 'center' }}>
          <a href="https://github.com/izakfr/variant64">Variant64 - Github Repo</a>
        </p>
      </aside>
    </div>
  );
}

export default function Routing() {
  return (
    <>
      <BrowserRouter>
        <Header />

        <Routes>
          <Route path="/" element={<Homepage />}></Route>
          <Route path="/chess" element={<Gamepage />}></Route>
          <Route
            path={'*'} //match all routes - for http 404
            element={<div>404!!! ahh!!!</div>}
          />
        </Routes>

        <Footer />
      </BrowserRouter>
    </>
  );
}

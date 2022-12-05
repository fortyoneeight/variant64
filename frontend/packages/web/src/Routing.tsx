import React from 'react';
import { BrowserRouter, Route, Routes, Navigate, Outlet } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { playerState } from './store/atoms';
import { Gamepage, Homepage, Onboarding } from './components/pages';

function Header() {
  return (
    <aside className="header">
      <h1 style={{ color: '#29EF67' }}>Variant 64</h1>
    </aside>
  );
}

function Footer() {
  return (
    <aside className="footer">
      <a style={{ color: '#29EF67' }} href="https://github.com/fortyoneeight/variant64">
        Github Repo
      </a>
    </aside>
  );
}

const PrivateRoute = () => {
  const [player, _] = useRecoilState(playerState);
  return player.id ? <Outlet /> : <Navigate to="/" />;
};

export default function Routing() {
  return (
    <>
      <BrowserRouter>
        <Header />
        <Routes>
          <Route path="/" element={<Onboarding />} />
          <Route path="/home" element={<PrivateRoute />}>
            <Route path="/home" element={<Homepage />} />
          </Route>
          <Route path="/room/:id" element={<PrivateRoute />}>
            <Route path="/room/:id" element={<Gamepage />} />
          </Route>
          <Route path={'*'} element={<div>404!!! ahh!!!</div>} />
        </Routes>
        <Footer />
      </BrowserRouter>
    </>
  );
}

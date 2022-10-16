import React from "react";
import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import Gamepage from "./Gamepage";
import Homepage from "./Homepage";

function Header() {
    return (
        <Link to='/'>
            <h1 style={{ textAlign: 'center' }}>
                Variant64 Chess
            </h1>
        </Link>
    )
}

function Footer() {
    return <aside>
        <a href='https://github.com/izakfr/variant64'><small>Variant64 - Github Repo</small></a>
    </aside>
}



export default function Routing() {
    return <>
        <BrowserRouter>

            <Header />

            <Routes>
                <Route
                    path="/"
                    element={<Homepage />}>
                </Route>
                <Route
                    path="/chess"
                    element={<Gamepage />}>
                </Route>
                <Route path={"*"} //match all routes - for http 404
                    element={<div>404!!! ahh!!!</div>} />
            </Routes>

            <Footer />

        </BrowserRouter>



    </>

}

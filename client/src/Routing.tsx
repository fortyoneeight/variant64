import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Gamepage from "./Gamepage";
import Homepage from "./Homepage";

export default function Routing() {
    return (
        <BrowserRouter>
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
        </BrowserRouter>
    );
}

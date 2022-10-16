import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

export default function Routing() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<div>wowie i am le homepage :) try /chess</div>}></Route>
                <Route path="/chess" element={<div>wowie i am le chess game :3c</div>}></Route>
                <Route path={"*"} //match all routes - for http 404
                    element={<div>404!!! ahh!!!</div>} />
            </Routes>
        </BrowserRouter>
    );
}

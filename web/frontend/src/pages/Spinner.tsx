import React from "react";
import Spinner1 from "react-bootstrap/Spinner"
export default function Spinner() {
    return (
        <div className="flex items-center justify-center h-[20vw] w-screen">
        <Spinner1 animation={"border"}></Spinner1>
        </div>
    )
}
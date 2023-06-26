import NavBar from "../navbar";
import {useParams} from "react-router";
import React, {useEffect, useState} from "react";

export default function GuildCommand() {
    let {id, command} = useParams()
    let [cmd, setCommand] = useState({} as Command)
    let [loaded, setLoaded] = useState(false)
    useEffect(() => {

        function fetchCommands()  {
            fetch(`/api/guilds/${id}/commands`, {credentials: "include"}).then(res => res.json()).then(data => {
                if (data.error) {
                    //setLoaded(true);
                    return
                }
                setCommand(data)
                setLoaded(true);
            })
        }
        fetchCommands()
    }, []);
    return (
        <div className={"text-white"}>
            <NavBar/>
            <div className={"h-20"}/>
            <h1 className={"text-center w-screen"}>
                Edit {command} command
                <br/>
                <a className={"no-underline"} target="_blank" href={`https://glitchii.github.io/embedbuilder/?data=JTdCJTdE`}>Response editor</a>
            </h1>

        </div>
    )
}
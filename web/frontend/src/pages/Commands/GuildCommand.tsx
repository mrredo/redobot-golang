import NavBar from "../navbar";
import {useParams} from "react-router";
import React, {useEffect, useState} from "react";
import InputGroup from "react-bootstrap/InputGroup";
import Form from "react-bootstrap/Form";

export default function GuildCommand() {
    let {id, command} = useParams()
    let [cmd, setCommand] = useState({} as Command)
    let [loaded, setLoaded] = useState(false)
    useEffect(() => {

        function fetchCommands()  {
            fetch(`/api/guilds/${id}/commands/${command}`, {credentials: "include"}).then(res => res.json()).then(data => {
                if (data.error) {
                    //setLoaded(true);
                    return
                }
                setCommand(data)
                let desc = document.getElementById("description") as HTMLInputElement,
                    res = document.getElementById("response") as HTMLInputElement
                desc.value = data.description
                res.value = data.response
                setLoaded(true);
            })
        }
        fetchCommands()
    }, []);
    function UpdateCommand() {

    }
    function DeleteCommand() {

    }
    return (
        <div className={"text-white"}>
            <NavBar/>
            <div className={"h-20"}/>
            <h1 className={"text-center w-screen"}>
                Edit {command} command
                <br/>
                <a className={"no-underline"} target="_blank" href={`https://glitchii.github.io/embedbuilder/?data=JTdCJTdE`}>Response editor</a>
            </h1>
            <div className="flex justify-center items-center">
                <button className={"mx-[3rem] p-3 border-2 bg-green-700 text-lg transition-all duration-300 border-gray-800 hover:border-white hover:border-3 rounded-md hover:rounded-xl m-[1rem] "}>
                    Update</button>
                <button className={"mx-[3rem] p-3 border-2 bg-red-800 hover:bg-red-700 text-lg transition-all duration-300 border-gray-800 hover:border-white hover:border-3 rounded-md hover:rounded-xl m-[1rem]"} >
                    Delete</button>
            </div>
            <div className=" flex justify-center items-center">
                <div className={"w-[50vw]"}>
                    <InputGroup className="mb-3 ">
                        <InputGroup.Text id="basic-addon1">Description</InputGroup.Text>
                        <Form.Control
                            max={100}
                            min={1}

                            id="description"
                            placeholder="Description"
                            aria-label="cmd-description"
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                </div>


            </div>
            <div className=" flex justify-center items-center">
                <div className={"w-[50vw]"}>
                    <InputGroup>
                        <InputGroup.Text>response</InputGroup.Text>
                        <Form.Control as="textarea" aria-label=""
                        id={"response"}
                        />
                    </InputGroup>
                </div>


            </div>

        </div>
    )
}
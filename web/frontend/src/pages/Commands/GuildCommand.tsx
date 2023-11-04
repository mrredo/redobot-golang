import NavBar from "../navbar";
import {useParams} from "react-router";
import React, {ChangeEvent, useEffect, useState} from "react";
import InputGroup from "react-bootstrap/InputGroup";
import Form from "react-bootstrap/Form";
import { redirect } from 'react-router-dom';
import Swal1 from 'sweetalert2'
import withReactContent from "sweetalert2-react-content";
let Swal = withReactContent(Swal1)
export default function GuildCommand() {
    let {id, command} = useParams()
    let [cmd, setCommand] = useState({} as Command)
    let [placeholder, setPlaceholder] = useState({})

    let [loaded, setLoaded] = useState(false)
    useEffect(() => {
        function fetchPlaceholders() {
            fetch('/api/placeholders/member,user,guild,command', {credentials: "include"}).then(res => res.json()).then((data: string[]) => {
                setPlaceholder(data)
            })

        }
        function fetchCommands()  {
            fetch(`/api/guilds/${id}/commands/${command}`, {credentials: "include"}).then(res => res.json()).then(data => {
                if (data.error) {
                    window.location.href = `/guilds/${id}/commands`
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
        fetchPlaceholders()
        fetchCommands()
    }, []);
    function UpdateCommand() {
        let up = document.getElementById("upd") as HTMLButtonElement
        up.disabled = true
        fetch(`/api/guilds/${id}/commands?type=update`, {
            method: "POST",
            body: JSON.stringify(cmd),

        }).then((res) => res.json()).then((data) => {
            if (data.error) {
                Swal.fire({
                    icon: "error",
                    title: "Failed updating command",
                    text: data.error,
                })
                setTimeout(() => {
                    up.disabled = false
                }, 2000)
                return
            }
            Swal.mixin({
                toast: true,
                position: 'top-end',
                showConfirmButton: false,
                timer: 3000,
                timerProgressBar: true,
                didOpen: (toast) => {
                    toast.addEventListener('mouseenter', Swal.stopTimer)
                    toast.addEventListener('mouseleave', Swal.resumeTimer)
                },

            }).fire({
                icon: "success",
                title: "Updated the command",

            })
            setTimeout(() => {
                up.disabled = false
            }, 2000)

        })
    }

    function DeleteCommand() {
        let del = document.getElementById("del") as HTMLButtonElement
        del.disabled = true
        fetch(`/api/guilds/${id}/commands/${command}`, {
            method: "DELETE",
            credentials: "include"

        }).then((res) => {
            res.json().then((data) => {
                Swal.fire({
                    icon: "error",
                    title: "Failed deleting command",
                    text: data.error,
                })
                setTimeout(() => {
                    del.disabled = false
                }, 2000)
            }).catch((e) => {
                window.location.href = `/guilds/${id}/commands`

            })

        })
    }

    function parsePlaceholder() {
        return (<>
        </>)
    }
    const handleDescChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        cmd.description = event.target.value
        setCommand(cmd)
    };
    const handleResponseChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        cmd.response = event.target.value
        setCommand(cmd)
    };
    return (
        <div className={"text-white"}>
            <NavBar/>
            <div className={"h-20"}/>
            <h1 className={"text-center w-screen"}>
                Edit {command} command
                <br/>
                <a className={"no-underline"} target="_blank" href={`https://glitchii.github.io/embedbuilder/?data=JTdCJTdE`}>Response editor</a>
                <br/>

            </h1>
            <div className="flex justify-center items-center">
                <button id={"upd"} onClick={() => UpdateCommand()} className={"font-bold mx-[3rem] p-3 border-2 bg-green-700 text-lg transition-all duration-300 border-gray-800 hover:border-white hover:border-3 rounded-md hover:rounded-xl m-[1rem] "}>
                    Update</button>
                <button id={"del"} onClick={() => DeleteCommand()} className={"font-bold mx-[3rem] p-3 border-2 bg-red-800 hover:bg-red-700 text-lg transition-all duration-300 border-gray-800 hover:border-white hover:border-3 rounded-md hover:rounded-xl m-[1rem]"} >
                    Delete</button>
            </div>
            <div className=" flex justify-center items-center text-lg">

                Placeholders:
            </div>
            <div className=" flex justify-center items-center">
                {/*<span className={"text-lg "}></span>*/}

                </div>
            <div className=" flex justify-center items-center">

                <div className={"w-[50vw]"}>
                    <InputGroup className="mb-3 ">
                        <InputGroup.Text id="basic-addon1">Description</InputGroup.Text>
                        <Form.Control
                            onChange={(event) => handleDescChange(event)}
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
                                      onChange={(event) => handleResponseChange(event)}
                        id={"response"}
                        />
                    </InputGroup>
                </div>


            </div>

        </div>
    )
}
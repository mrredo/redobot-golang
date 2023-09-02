import React, {useEffect, useState} from "react"
import NavBar from "../navbar";
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import {AiOutlinePlus} from "react-icons/ai";
import {SimpleSpinner, Spinner} from "../Spinner"
import {useParams} from "react-router";
import Swal1 from 'sweetalert2'
import withReactContent from "sweetalert2-react-content";
import {Simulate} from "react-dom/test-utils";
import error = Simulate.error;
let Swal = withReactContent(Swal1)
/*
Command name - 1-32 chars
command description 1-100 chars

*/
export default function GuildCommands() {
    let {id} = useParams()
    let [loaded, setLoaded] = useState(false)
    let [commands, setCommands]= useState({} as MapCommand)
    let [listCommands, setListOfCommands]= useState([] as Command[])
    let [commandCount, setCmdCount] = useState(0)
    let [notsynced, setSynced] = useState(false) as any
    async function SyncCommands() {
        fetch(`/api/guilds/${id}/commands/reregister`, {
            method: "POST",
            credentials: "include"
        }).catch(error => console.error)
    }
    if (notsynced) {
        SyncCommands()
    }
    const updateSynced = (keys: string[], data: MapCommand) => {

        for (let val of keys) {
            let cmd = data[val]
            if(cmd.registered == false) {
                setSynced(true);
                return
            }
        }
    };
    //set commands
    //ser command count
    useEffect(() => {

        async function fetchCommands()  {
            fetch(`/api/guilds/${id}/commands`, {credentials: "include"}).then(res => res.json()).then(data => {
                if (data.error) {
                    //setLoaded(true);
                    return
                }
                setCommands(data)
                setLoaded(true);
                updateSynced(Object.keys(data), (data as MapCommand))

            })
        }
       fetchCommands()
    }, []);

    function UpdateCommand(index: number, newCommand: Command) {

    }
    function ModalCreateCommand() {


        Swal.fire({
            title: <strong>Create a new command</strong>,
            html:
                <div>
                    <InputGroup className="mb-3">
                        <InputGroup.Text id="basic-addon1">/</InputGroup.Text>
                        <Form.Control

                            id="name"
                            placeholder="Name"
                            aria-label="cmd-name"
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                    <InputGroup className="mb-3">
                        <InputGroup.Text id="basic-addon1">/</InputGroup.Text>
                        <Form.Control
                            id="desc"
                            placeholder="Description"
                            aria-label="cmd-description"
                            aria-describedby="basic-addon1"
                        />
                    </InputGroup>
                </div>,
            showCancelButton: true,
            showConfirmButton: true,
            confirmButtonText: 'Create',
            preConfirm: () => {
                    let name = Swal.getPopup()?.querySelector("#name") as HTMLInputElement,
                        description = Swal.getPopup()?.querySelector("#desc") as HTMLInputElement
                    if (!(1 <= name.value.length && name.value.length  <= 32)) {
                        Swal.showValidationMessage("Name must be from 1-32 characters")
                    } else if (!(1 <= description.value.length  && description.value.length  <= 100)) {
                        Swal.showValidationMessage("Description must be from 1-100 characters")
                    }
                    return {name: name.value, description: description.value, response: `{"content": "This is the default message of this command {commandname}"}` } as Command
                    },
            }).then((res) => {
                if(!res.isConfirmed) return
                let cmd = res.value as Command
                Swal.fire({
                    title: <strong>Loading...</strong>,
                    html:<div className={"flex items-center justify-center h-20"}> <SimpleSpinner/></div>,
                    padding: 10,
                    showConfirmButton: false,

                })
            fetch(`/api/guilds/${id}/commands?type=register`, {
                credentials: "include",
                method: "POST",
                body: JSON.stringify(cmd)
            }).then(res => res.json()).then((data: any) => {
                if (data.error) {
                    Swal.fire({
                        icon: "error",
                        title: "Couldn't create command",
                        text: `${data.error}`,
                        showCloseButton: true,

                    })
                    return
                }
                Swal.close()
                setCommands((prevCommands) => ({
                    ...prevCommands,
                    [data.name]: data,
                }));

            })
        })

    }
    function CreateCommand(command: Command) {
        fetch(`/api/guilds/${id}/commands?type=register`, {
            credentials: "include",
            method: "POST",
            body: JSON.stringify(command)
        }).then(res => res.json()).then((data: any) => {

            if (data.error != undefined) {
                console.log(data.error)
                Swal.fire({
                    icon: "error",
                    title: "Couldn't create command",
                    text: `${data.error}`,
                    showCloseButton: true,


                })
            } else {
                Swal.fire({
                    icon: "success",
                    title: "Successfully created command!",
                    showCloseButton: true,
                    timer: 2000,
                })
                let cmd = data as Command
                setCommands((prevCommands) => ({
                    ...prevCommands,
                    [data.name]: cmd,
                }));

            }


        })
    }

    function ValidateInputName(ev: InputEvent) {

    }
    return (
        <div className={"text-white"}>
            <NavBar/>
            <div className={"m-16"}></div>
            <h1 className={"text-center w-screen"}>
                Add your own commands to a server!
                <br/>
                <a className={"no-underline"} target="_blank" href={`https://glitchii.github.io/embedbuilder/?data=JTdCJTdE`}>Response editor</a>
            </h1>
            <div className="flex justify-between m-10 border-2 p-2 rounded-lg">
                <div className="flex items-center">
                    <h3>Add a command</h3>
                </div>
                <button onClick={() => {
                    ModalCreateCommand()
                    // CreateCommand({name: "hello", description: "hwddwdwwd", response: `{\"content\":\"hello bozo\"}`})
                }
                } className="ml-auto border-2 rounded-lg  transition-all duration-300 hover:rounded-xl bg-gray-700 hover:bg-green-500">
                    <AiOutlinePlus className="w-10 h-10" />
                </button>
            </div>

            <div className={"commands grid grid-cols-3 md:grid-cols-2 m-[2.4rem] gap-2 sm:grid-cols-1"}>
                {!loaded ? (
                    <Spinner />
                ) : (
                    Object.keys(commands).map((value, index) => (


                            <a className={"text-center border-gray-900 rounded-lg border-2 hover:border-3 hover:border-white p-4 no-underline text-white bg-gray-800"}  href={`/guilds/${id}/commands/${commands[value].name}`}>
                                <strong>/{commands[value].name}</strong>
                                <br/>
                                <span className="text-gray-600">{commands[value].description}</span>
                            </a>



                    ))
                )}

            </div>


        </div>
    )
}
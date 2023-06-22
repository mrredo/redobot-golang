import React, {useEffect, useState} from "react"
import NavBar from "../navbar";
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import {AiOutlinePlus} from "react-icons/ai";
import Spinner from "../Spinner"
import {useParams} from "react-router";
import Swal1 from 'sweetalert2'
import withReactContent from "sweetalert2-react-content";
let Swal = withReactContent(Swal1)
interface Command {
    name: string
    description: string
    response: string
    id?: string
    registered?: boolean
}
interface MapCommand {
    [key: string]: Command
}
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
    //get commands
    //set commands
    //ser command count
    useEffect(() => {

        function fetchCommands()  {
            fetch(`/api/guilds/${id}/commands`, {credentials: "include"}).then(res => res.json()).then(data => {
                if (data.error) {
                    //setLoaded(true);
                    return
                }
                setCommands(data)
                setLoaded(true);
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
                let cmd = res.value as Command

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
            <h1 className={"text-center"}>
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
            <div className={"commands grid grid-cols-3 md:grid-cols-2 m-2 gap-2 sm:grid-cols-1"}>
                {!loaded ? (
                    <Spinner />
                ) : (
                    Object.keys(commands).map((value, index) => (
                        <div className={"border-2 rounded-md"}>


                            <div className={"cmdname p-2"}>
                                {/*<InputGroup className="mb-3">*/}
                                {/*    <InputGroup.Text id="basic-addon1">/</InputGroup.Text>*/}
                                {/*    <Form.Control*/}
                                {/*        value={commands[value].name}*/}
                                {/*        disabled*/}
                                {/*        placeholder="Name"*/}
                                {/*        aria-label="cmd-name"*/}
                                {/*        aria-describedby="basic-addon1"*/}
                                {/*    />*/}
                                {/*</InputGroup>*/}
                                {/*<InputGroup className="mb-3">*/}
                                {/*    <InputGroup.Text id="basic-addon1">/</InputGroup.Text>*/}
                                {/*    <Form.Control*/}
                                {/*        value={commands[value].description}*/}
                                {/*        placeholder="Description"*/}
                                {/*        aria-label="cmd-description"*/}
                                {/*        aria-describedby="basic-addon1"*/}
                                {/*    />*/}
                                {/*</InputGroup>*/}

                                {/*<InputGroup>*/}
                                {/*    <InputGroup.Text>Response</InputGroup.Text>*/}
                                {/*    <Form.Control value={commands[value].response} as="textarea" aria-label="With textarea" />*/}
                                {/*</InputGroup>*/}

                                {/*<button className={"border-2 m-2 p-2 text-lg transition-all duration-300 rounded-md hover:bg-green-600 hover:rounded-xl"} >Update</button>*/}
                            </div>


                        </div>
                    ))
                )}

            </div>


        </div>
    )
}
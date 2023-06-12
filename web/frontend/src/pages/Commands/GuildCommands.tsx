import React, {useEffect, useState} from "react"
import NavBar from "../navbar";
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import {AiOutlinePlus} from "react-icons/ai";
interface Command {
    name: string
    description: string
    response: string

}
/*
Command name - 1-32 chars
command description 1-100 chars

*/
export default function GuildCommands() {
    let [commands, setCommands]= useState(([
        {
            name: "hello",
            description: "this is hello command",
            response: JSON.stringify({ content: "hello bozo"})

        },
        {
            name: "who",
            description: "this is who command",
            response: JSON.stringify({ content: "who are you"})

        },
        {
            name: "why",
            description: "this is why command",
            response: JSON.stringify({ content: "why did u do that"})

        },
        {
            name: "why1",
            description: "this is why1 command",
            response: JSON.stringify({ content: "why1 did u do that"})

        },
    ] as Command[]))
    let [commandCount, setCmdCount] = useState(0)
    //get commands
    //set commands
    //ser command count
    function UpdateCommand(index: number, newCommand: Command) {
        console.log(index, newCommand)
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
                <button className="ml-auto border-2 rounded-lg transition-all duration-300 bg-gray-700 hover:bg-green-500">
                    <AiOutlinePlus className="w-10 h-10" />
                </button>
            </div>
            <div className={"commands grid grid-cols-3 md:grid-cols-2 m-2 gap-2 sm:grid-cols-1"}>
                {commands.map((command:Command, index   ) => (
                    <div className={"border-2 rounded-md"}>


                        <div className={"cmdname p-2"}>
                            <InputGroup className="mb-3">
                                <InputGroup.Text id="basic-addon1">/</InputGroup.Text>
                                <Form.Control
                                    value={command.name}
                                    placeholder="Name"
                                    aria-label="cmd-name"
                                    aria-describedby="basic-addon1"
                                />
                            </InputGroup>
                            <InputGroup className="mb-3">
                                <InputGroup.Text id="basic-addon1">/</InputGroup.Text>
                                <Form.Control
                                    value={command.description}
                                    placeholder="Description"
                                    aria-label="cmd-description"
                                    aria-describedby="basic-addon1"
                                />
                            </InputGroup>

                            <InputGroup>
                                <InputGroup.Text>Response</InputGroup.Text>
                                <Form.Control value={command.response} as="textarea" aria-label="With textarea" />
                            </InputGroup>
<button onClick={() => UpdateCommand(index, command)}>submit</button>
                        </div>
                        <hr />
                        <div className={"cmddescription"}>

                        </div>

                    </div>
                    )

                )}
            </div>


        </div>
    )
}
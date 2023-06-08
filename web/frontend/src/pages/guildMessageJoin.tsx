import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import NavBar from "./navbar";
export default function GuildMessagesJOIN() {
    const { id } = useParams();
    const [guild, setGuild] = useState({}) as any;
    const [join, setJoin] = useState({}) as any;
    const [channels, setChannel] = useState([]) as any;
    const [msg, setMessage] = useState({}) as any
    useEffect(() => {
        // const fetchData = async () => {
        // fetch(`/api/guilds/${id}`)
        //   .then(response => response.json())
        //   .then(data => setGuild(data));
        // };
        // fetchData();
        // const fetchJoin = async () => {
        //   fetch(`/api/guilds/${id}/messages/join/`)
        //     .then(response => response.json())
        //     .then(data => setJoin(data));
        //   };
        //   fetchJoin();
        const fetchChannels = async () => {
            fetch(`/api/guilds/${id}/channels?type=text`)
                .then(response => response.json())
                .then(data => setChannel(data));
        };
        fetchChannels();
        const fetchMessages = async () => {
            let  textb = document.getElementById("jsondata") as HTMLTextAreaElement
            let selectchan = document.getElementById("channels") as HTMLSelectElement
            let enabled = document.getElementById("enabled") as HTMLInputElement
            fetch(`/api/guilds/messages/${id}/join`, {
                method: "GET",
                credentials: "include",
            })                .then(response => response.json())
                .then(data => {
                    if (!data.error) {
                        selectchan.value = data.channel_id
                        textb.value = data.json_data
                        enabled.checked = data.enabled
                    }
                });
        };
        fetchMessages();

    }, []);
    function SendDataMessage() {
        let  textb = document.getElementById("jsondata")
        let selectchan = document.getElementById("channels")
        let enabled = document.getElementById("enabled") as HTMLInputElement
        fetch(`/api/guilds/${id}/${(selectchan as HTMLSelectElement).value}/join?enabled=${enabled.checked}`, {
            method: "PUT",
            credentials: "include",
            body: (textb as HTMLTextAreaElement).value
        })
    }
    function GetDataMessage() {

    }
    return(
        <div className="main">
            <div className="grid grid-cols-2 sm:grid-cols-1 m-3">
                <div className="joins border-white border-2">
                    <h1 className="text-center rounded-lg text-white font-bold py-2">
                        <span className="">Join messages</span>

                    </h1>
                    <div className="content text-white text-center">

                        <div className="title">

                            <input id={"enabled"} type={"checkbox"} /> Enabled?
                            <br/>
                            Select channel where to send message
                            <br/>
                            <select id={"channels"} className="text-black">
                                {channels.length != 0? channels.map((channel: any) => (
                                    <option value={channel.id}>{channel.name}</option>
                                )) : ""}
                            </select>
                            <br/>
                            Embed JSON data
                            <br />


                            <textarea id={"jsondata"} className={"resize w-[20vw] h-[5vw] rounded-lg bg-gray-800"}></textarea>
                            <br/>
                            <button onClick={() => SendDataMessage()} className="text-3xl border-2 m-2 p-2 rounded-lg hover:rounded-xl hover:bg-green-500 transition-all duration-300" >Save!</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import Swal from 'sweetalert2'
import NavBar from "./navbar";
import GuildMessageComponent from "./guildMessageComponent";
export default function GuildMessages() {
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
        const fetchMessages1 = async () => {
            let  textb = document.getElementById("jsondata1") as HTMLTextAreaElement
            let selectchan = document.getElementById("channels1") as HTMLSelectElement
            let enabled = document.getElementById("enabled1") as HTMLInputElement
            fetch(`/api/guilds/messages/${id}/leave`, {
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
        fetchMessages1();

      }, []);
    function SendDataMessage() {
       let  textb = document.getElementById("jsondata")
        let selectchan = document.getElementById("channels")
        let enabled = document.getElementById("enabled") as HTMLInputElement
        fetch(`/api/guilds/${id}/${(selectchan as HTMLSelectElement).value}/join?enabled=${enabled.checked}`, {
            method: "PUT",
            credentials: "include",
            body: (textb as HTMLTextAreaElement).value
        }).then(response => response.json()).then(data => {
            if(data.error) {
                Swal.fire({
                    icon: "error",
                    title: "Couldn't save data",
                    text: `Reason: ${data.error}`,

                })
            } else {
                Swal.fire({
                    icon: "success",
                    title: "Successfully saved data!",
                    showCloseButton: true,
                    timer: 2000,
                })
            }
        })
    }
    function SendDataMessage1() {
        let  textb = document.getElementById("jsondata1")
        let selectchan = document.getElementById("channels1")
        let enabled = document.getElementById("enabled1") as HTMLInputElement
        fetch(`/api/guilds/${id}/${(selectchan as HTMLSelectElement).value}/leave?enabled=${enabled.checked}`, {
            method: "PUT",
            credentials: "include",
            body: (textb as HTMLTextAreaElement).value
        }).then(response => response.json()).then(data => {
            if(data.error) {
                Swal.fire({
                    icon: "error",
                    title: "Couldn't save data",
                    text: `Reason: ${data.error}`,

                })
            } else {
                Swal.fire({
                    icon: "success",
                    title: "Successfully saved data!",
                    showCloseButton: true,
                    timer: 2000,
                })
            }
        })
    }
    return(
        <div className="main">
            <NavBar />
            <div className="text-center pb-2 border-b-white text-white text-5xl mt-16 mb-3 sm:text-[2.7rem]">
                Greetings
                <br/>
                <a target="_blank" href={`https://glitchii.github.io/embedbuilder/?data=JTdCJTdE`}>Embed builder</a>
            </div>
            <div className="grid grid-cols-2 sm:grid-cols-1 m-3 gap-3">

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
                <div className="joins border-white border-2">
                    <h1 className="text-center rounded-lg text-white font-bold py-2">
                        <span className="">Leave messages</span>

                    </h1>
                    <div className="content text-white text-center">

                        <div className="title">

                            <input id={"enabled1"} type={"checkbox"} /> Enabled?
                            <br/>
                            Select channel where to send message
                            <br/>
                            <select id={"channels1"} className="text-black">
                                {channels.length != 0? channels.map((channel: any) => (
                                    <option value={channel.id}>{channel.name}</option>
                                )) : ""}
                            </select>
                            <br/>
                            Embed JSON data
                            <br />


                            <textarea id={"jsondata1"} className={"resize w-[20vw] h-[5vw] rounded-lg bg-gray-800"}></textarea>
                            <br/>
                            <button onClick={() => SendDataMessage1()} className="text-3xl border-2 m-2 p-2 rounded-lg hover:rounded-xl hover:bg-green-500 transition-all duration-300" >Save!</button>
                        </div>
                    </div>
                </div>
                {/*<GuildMessageComponent type={"Join"} ></GuildMessageComponent>*/}
            </div>
        </div>
    )
}
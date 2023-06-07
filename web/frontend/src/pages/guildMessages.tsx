import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import NavBar from "./navbar";
export default function GuildMessages() {
    const { id } = useParams();
    const [guild, setGuild] = useState({}) as any;
    const [join, setJoin] = useState({}) as any;
    const [channels, setChannel] = useState([]) as any;
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
      }, []);
    function SendDataMessage() {
       let  textb = document.getElementById("jsondata")
        let selectchan = document.getElementById("channels")
        fetch(`/api/guilds/${id}/${(selectchan as HTMLSelectElement).value}/join`, {
            method: "PUT",
            credentials: "include",
            body: (textb as HTMLTextAreaElement).value
        })
    }
    return(
        <div className="main">
            <NavBar />
            <div className="text-center pb-2 border-b-white text-white text-5xl mt-16 mb-3 sm:text-[2.7rem]">
                Edit redobot responses
                <br/>
                <a target="_blank" href={`https://glitchii.github.io/embedbuilder/?data=JTdCJTdE`}>Embed builder</a>
            </div>
            <div className="grid grid-cols-2 sm:grid-cols-1 m-3">
              <div className="joins border-white border-2">
                <h1 className="text-center rounded-lg text-white font-bold py-2">
                  <span className="">Join messages</span>
                </h1>
                <div className="content text-white text-center">

                  <div className="title">
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
                      <button onClick={() => SendDataMessage()} className="text-3xl border-2 m-2 p-2 rounded-lg hover:rounded-xl hover:bg-green-700 transition-all duration-300" >Save!</button>
                  </div>
                </div>
              </div>
            </div>
        </div>
    )
}
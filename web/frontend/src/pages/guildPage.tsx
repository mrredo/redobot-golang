import React, {useState} from "react";
import Guilds from "./guilds";
import NavBar from "./navbar";
import {Spinner} from "./Spinner";
export default function GuildPage() {
    const [user1, setUser] = useState({} as any)

    window.addEventListener("keydown", function(e) {
        if(e.key === "b") {
            const guilds = document.getElementsByClassName("1-guild")
            for(let i = 0; i < guilds.length; i++) {
                const guild = guilds[i]
                guild.classList.add("animate-spin")
            }
            setTimeout(() => {
                for(let i = 0; i < guilds.length; i++) {
                    const guild = guilds[i]
                    guild.classList.remove("animate-spin")
                }
            }, 1000 * 3)

        }
    })
    console.log(user1)
    return (
        <div className="page">
            <NavBar discordloginpopup={true} path={"/guilds"} setuser={function (user) {
                setUser(user)
                console.log(user)
            }} />
            <div className="text-center text-white text-5xl mt-16 mb-3 sm:text-[2.7rem]">
                Select a server
            </div>
            {user1.username? (
                <Guilds />
            ) : (
                <Spinner/>
            )}

        </div>
    )
}
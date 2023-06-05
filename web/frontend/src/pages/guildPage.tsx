import React from "react";
import Guilds from "./guilds";
import NavBar from "./navbar";
export default function GuildPage() {
    window.addEventListener("keydown", function(e) {
        if(e.key === "r" && "e" && "d") {
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
    return (
        <div className="page">
            <NavBar />
            <div className="text-center text-white text-5xl mt-16 mb-3 sm:text-[2.7rem]">
                Select a server
            </div>
            <Guilds />
        </div>
    )
}
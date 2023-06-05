import { useEffect, useState } from "react";
import { GuildData } from "src/interfaces/GuildData";

const fetchData = async (setGuilds: any) => {
    fetch("/api/guilds")
        .then((response) => response.json())
        .then((data) => setGuilds(data));
};

export default function Guilds() {
    const [guilds, setGuilds] = useState([]);
    const [updateTrigger, setUpdateTrigger] = useState(false); // New state variable

    const url =
        "https://www.freepnglogos.com/uploads/discord-logo-png/discord-logo-logodownload-download-logotipos-1.png";

    useEffect(() => {
        fetchData(setGuilds);
    }, [updateTrigger]); // Use updateTrigger as a dependency

    function OpenBotadd(guildid: string) {
        let win = window.open(
            `http://localhost:4000/addbot?guild_id=${guildid}`,
            "popup",
            "width=400,height=600"
        );
        var popupTick = setInterval(function() {
            if (win?.closed) {
                window.location.href =  `/guilds/${guildid}/info`;
            }
        }, 500);


    }

    return (
        <div className="bg-gray-700 grid-cols-3 sm:grid-cols-1 md:grid-cols-2 grid auto-rows-auto guilds p-13">
            {guilds.map((guild: any) => (
                <div
                    className="1-guild rounded-3xl grid place-items-center w-12/12 m-2 p-3 border-2 border-gray-800 shadow-2xl"
                    key={guild.id}
                >
                    {guild.icon ? (
                        <img
                            width="90"
                            height="90"
                            src={`https://cdn.discordapp.com/icons/${guild.id}/${guild.icon}`}
                            className="rounded-full"
                        />
                    ) : (
                        <img
                            width="90"
                            height="90"
                            src={url}
                            className="rounded-full"
                        />
                    )}
                    <p className="text-white text-3xl sm:text-3xl mt-2">{guild.name}</p>
                    {guild.botInServer ? (
                        <a
                            href={`/guilds/${guild.id}/info`}
                            className="transition-all text-gray-300 rounded-lg hover:rounded-[1rem] no-underline text-2xl border-gray-800 mt-1 mb-1 border-4 p-2 hover:no-underline duration-300 hover:bg-gray-800 hover:text-white hover:shadow-2xl"
                        >
                            configure
                        </a>
                    ) : (
                        <a
                            onClick={() => OpenBotadd(`${guild.id}`)}
                            className="hover:cursor-pointer transition-all text-gray-300 rounded-lg hover:rounded-[1rem] no-underline text-2xl border-gray-800 mt-1 mb-1 border-4 p-2 hover:no-underline duration-300 hover:bg-gray-800 hover:text-white hover:shadow-2xl"
                        >
                            Add bot
                        </a>
                    )}
                </div>
            ))}
        </div>
    );
}

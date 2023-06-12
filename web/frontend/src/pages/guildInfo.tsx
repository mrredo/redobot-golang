import { useEffect, useState } from "react";
import { useParams } from "react-router";
import NavBar from "./navbar";
export default function GuildInfo() {
    const { id } = useParams();
    const [guild, setGuild] = useState({});
    useEffect(() => {
        const fetchData = async () => {
        fetch(`/api/guilds/${id}`)
          .then(response => response.json())
          .then(data => setGuild(data));
        };
        fetchData();
      }, []);
      return (
          <div className="data">
              <NavBar />
            <h1 className="text-white rounded-3xl lg:text-5xl sm:text-4xl p-3 ml-4 mr-4 mt-16 border-4 border-gray-800 text-center">
                Edit {(guild as any).name} server settings
            </h1>
            <div className="links grid grid-cols-3 md:grid-cols-2 m-2">
                    <a className="no-underline border-2 rounded-lg w-[95%] py-4 text-3xl md:text-2xl mt-3 font-bold text-center hover:no-underline transition-all duration-500 hover:bg-white hover:text-black" 
                        href={`/guilds/${id}/settings`}>
                        Settings
                    </a>
                    <a className="no-underline border-2 w-[95%] rounded-lg text-center py-4 text-3xl md:text-2xl mt-3 font-bold hover:no-underline transition-all duration-200 hover:bg-white hover:text-black" 
                        href={`/guilds/${id}/messages`}>
                            Messages
                    </a>
                <a className="no-underline border-2 w-[95%] rounded-lg text-center py-4 text-3xl md:text-2xl mt-3 font-bold hover:no-underline transition-all duration-200 hover:bg-white hover:text-black"
                   href={`/guilds/${id}/commands`}>
                    Custom commands
                </a>
            </div>
        </div>
      );
}
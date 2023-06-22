import Guilds from "./guilds";
import '../styles/scrollbar.css';
import NavBar from "./navbar";
import {} from "react-router-dom";
import { Parallax, ParallaxProvider, ParallaxBanner } from 'react-scroll-parallax';
export default function Home() {
  document.title = "Home";
  return (
        <body>
          <NavBar />
          <div className="home-page">
            <h1>Redobot, multipurpose bot</h1>
          </div>
        </body>
    )
}
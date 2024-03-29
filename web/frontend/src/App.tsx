import React, { useEffect, useState } from "react";
import Home from "./pages/home";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import NotFound from "./pages/notFound";
import GuildPage from "src/pages/guildPage";
import GuildInfo from "./pages/guildInfo";
import Temp from "./pages/Temp";
import GuildSettings from "./pages/guildSettings";
import GuildMessages from "./pages/guildMessages";
import { ParallaxProvider } from 'react-scroll-parallax';
import GuildCommands from "./pages/Commands/GuildCommands";
import GuildCommand from "./pages/Commands/GuildCommand";
import './styles/darkswal.scss'
import PaymentPremiumPage from "./pages/premium/PremiumPage";
import {DiscordLoginBtn} from "./pages/popups/discordloginpopup";
import SyntaxHighlight from "./pages/syntaxHighlight";
import ColoredWordsEditor from "./pages/TEST";
function App() {

  return (
  <ParallaxProvider>
    <BrowserRouter>
      <Routes>
          <Route path="/" element={<Home />} />
          <Route path="*" element={<NotFound />} />
          <Route path="/guilds" element={<GuildPage />} />
          <Route path="/guilds/:id/commands/" element={<GuildCommands />} />
          <Route path="/guilds/:id/commands/:command" element={<GuildCommand />} />
          <Route path="/guilds/:id/info/" element={<GuildInfo />} />
          <Route path="/guilds/:id/settings/" element={<GuildSettings />} />
          <Route path="/guilds/:id/messages/" element={<GuildMessages />} />
          <Route path="/majas/" element={<DiscordLoginBtn />}></Route>
          <Route path="/premium" element={<PaymentPremiumPage/>}></Route>
          <Route path="/test" element={<SyntaxHighlight/>}></Route>
          <Route path="/test1" element={<ColoredWordsEditor/>}></Route>
        </Routes>
    </BrowserRouter>
  </ParallaxProvider>
  );
}

export default App;

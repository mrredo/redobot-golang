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
function App() {

  return (
  <ParallaxProvider>
    <BrowserRouter>
      <Routes>
          <Route path="/" element={<Home />} />
          <Route path="*" element={<NotFound />} />
          <Route path="/guilds" element={<GuildPage />} />
          <Route path="/guilds/:id/commands/" element={<GuildCommands />} />
          <Route path="/guilds/:id/info/" element={<GuildInfo />} />
          <Route path="/guilds/:id/settings/" element={<GuildSettings />} />
          <Route path="/guilds/:id/messages/" element={<GuildMessages />} />
          <Route path="/majas/" element={<Temp />}></Route>
        </Routes>
    </BrowserRouter>
  </ParallaxProvider>
  );
}

export default App;

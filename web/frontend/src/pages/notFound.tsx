import React, { useEffect, useState } from "react"
import NavBar from "./navbar"
import "bootstrap/dist/css/bootstrap.min.css"
import "../styles/NavBar.css"
export default function NotFound() {
  document.title = "Error: 404 not found";
  const [data, setFetchedData] = useState({}) as any;
  useEffect(() => {
    const getData = async () => {
      const datas = await fetch("/api/url");
      setFetchedData(await datas.json());
    };
    getData();
  }, []);
    return (
      <body>
      <NavBar />
      <div
      className="
        flex
        items-center
        justify-center
        w-screen
        h-screen
        bg-gradient-to-r
        from-indigo-600
        to-blue-400
      "
    >
      <div className="px-40 py-20 md:px-20 md:py-15 sm:px-12 sm:py-10 bg-white rounded-md shadow-xl">
        <div className="flex flex-col items-center">
          <h1 className="font-bold text-blue-600 text-9xl sm:text-8xl">404</h1>
    
          <h6 className="mb-2 text-4xl font-bold text-center text-gray-800 sm:text-3xl">
            <span className="text-red-500">Oops!</span> Page not found
          </h6>
    
          <p className="mb-8 text-center text-gray-500 sm:text-lg">
            The page you’re looking for doesn’t exist.
          </p>
    
          <a
            href={data.current_url ?? "/"}
            className="transition-all duration-300 hover:rounded-lg px-6 py-2 no-underline text-sm font-semibold hover:bg-blue-500 text-blue-800 bg-blue-100"
            >Go home</a>
        </div>
      </div>
    </div>
    </body>
    )
}
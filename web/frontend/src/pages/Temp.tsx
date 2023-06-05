import React from "react";
import skice from "../stuff/skice.png";
export default function Temp() {
    document.title = "Programmēšana"
    return (
        <div className="bg-gradient-to-r from-yellow-500 to-red-700 w-screen h-screen">
            <nav className="">
                <div className="grid grid-cols-2">
                <a href="/majas" className="no-underline text-center my-4 py-2 border-blue-50 text-blue-50 font-bold bg-gray-900 text-4xl sm:text-3xl m-3 border-2 rounded-lg transition-all duration-500 hover:bg-white hover:border-black hover:text-black hover:font-bold">
                    Kā iesākās?</a>
                <a href="/majas" className="no-underline text-center my-4 py-2 border-blue-50 text-blue-50 font-bold bg-gray-900 text-4xl sm:text-3xl m-3 border-2 rounded-lg transition-all duration-500 hover:bg-white hover:border-black hover:text-black hover:font-bold">
                    Kur sāku?</a>
                </div>
            </nav>
            <h1 className="text-center border-2 text-white bg-gray-800 py-2 rounded-lg mx-4 sm:text-3xl">Kāpēc izvēlējos kļūt par programmētāju?</h1>
            <p className="text-center text-2xl py-2 border-2 mx-4 rounded-lg text-white bg-gray-800">Man vienkārši patika taisīt jaunas lietas</p>
            <h2 className="text-center border-2 text-white bg-gray-800 py-2 rounded-lg mx-4 sm:text-3xl">Cik gados sāku programmēt?</h2>
            <p className="text-center text-2xl py-2 border-2 mx-4 rounded-lg text-white bg-gray-800">Es sāku kad man bija 11 gadi, bet pametu un kad gandrīz bija 13 gadi es atsāku programmēšanu</p>
            <div className="numurets">
                <h2 className="text-center border-2 text-white bg-gray-800 py-2 rounded-lg mx-4 sm:text-3xl">Labākās vietas kur var programmēt!</h2>
                <div className="list border-blue-50 rounded-lg border-2 grid place-items-center text-2xl mx-4 mt-2 mb-1 pt-2  text-blue-50 bg-gray-800">
                    <ol className="list-decimal">
                        <li>
                            No mājām 
                        </li>
                        <li>
                            Offisā
                        </li>
                    </ol>
                </div>
            </div>
            <div className="aizimets">
                <h2 className="text-center border-2 text-white bg-gray-800 py-2 rounded-lg mx-4 sm:text-3xl">Kompānijas kuras pieņem darbā programmētājus!</h2>
                <div className="list border-blue-50 rounded-lg border-2 grid place-items-center text-2xl mx-4 mt-2 mb-1 pt-2  text-blue-50 bg-gray-800">
                    <ul className="list-disc">
                        <li>
                            Google
                        </li>
                        <li>
                            Facebook
                        </li>
                        <li>
                            Reddit
                        </li>
                        <li>
                            Discord
                        </li>
                        <li>
                            Draugiem Group
                        </li>
                    </ul>
                </div>
            </div>
            <div className="tabula">
                <h2 className="text-center border-2 text-white bg-gray-800 py-2 rounded-lg mx-4 sm:text-3xl">
                    Labākās kompānijas kur strādāt</h2>
                <div className="grid place-items-center">
            <div className="border-b border-gray-800 shadow">
                <div className="">
                    <div className="">
                        <table>
                            <thead className="bg-gray-800">
                                <tr>
                                    <th className="px-6 py-2 text-xs text-white">
                                        Kompānija
                                    </th>
                                    <th className="px-6 py-2 text-xs text-white">
                                        Alga(EUR)
                                    </th>
                                    <th className="px-6 py-2 text-xs text-white">
                                        Dibināšanas Datums
                                    </th>
                                    <th className="px-6 py-2 text-xs text-white">
                                        Saīsināts Datums
                                    </th>
                                </tr>
                            </thead>
                            <tbody className="bg-gray-600">
                                <tr className="whitespace-nowrap">
                                    <td className="px-6 py-4 text-sm text-white">
                                        Discord
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="text-sm text-white">
                                            90,000 - 214,676
                                        </div>
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="text-sm text-white">
                                            2015. gads 6. marts
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-sm text-white">
                                        2015-03-6
                                    </td>
                                </tr>
                                <tr className="whitespace-nowrap">
                                    <td rowSpan={20} className="px-6 py-4 text-sm text-white">
                                        Google
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="text-sm text-white">
                                            190,000 - 1,000,000
                                        </div>
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="text-sm text-white">
                                            1998. gads 4. semptembris
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-sm text-white">
                                        1998-9-04
                                    </td>
                                </tr>
                                <tr className="whitespace-nowrap">
                                    <td className="px-6 py-4 text-sm text-white">
                                        Reddit
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="text-sm text-white">
                                            88,784 - 190,584
                                        </div>
                                    </td>
                                    <td className="px-6 py-4">
                                        <div className="text-sm text-white">
                                            2005. gads 23. jūnijs
                                        </div>
                                    </td>
                                    <td className="px-6 py-4 text-sm text-white">
                                        2005-6-23
                                    </td>
                                </tr>


                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
                </div>
            </div>
            <div className="img grid place-content-center">
                <img src="data:image/jpeg;base64,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"
                    alt="Hakošanas foto"
                    width="300"
                    height="300"
                    className="border-4 border-black rounded-lg"
                />
            </div>
            <div className="skice grid place-items-center">
                <a download="RV_7B_skice" href={skice} className="no-underline p-4 border-2 m-2 border-gray-50 text-blue-50 text-2xl transition-all hover:bg-white font-bold hover:text-black duration-700">Nospiežiet šeit lai ielādētu skici!</a>
            </div>
        </div>
    )
}
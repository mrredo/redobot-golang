import React, {useState} from 'react';
import  "../styles/toggle.css"
import {useParams} from "react-router";
type Proptype = "Join" | "Leave"
interface props {
    type: Proptype
}
interface Message {
    id: string
    json_data: string
    channel_id: string
    type: Proptype
    enabled: boolean
}
const GuildMessageComponent: React.FC<props> = (props) => {
    let messageType = props.type
    const { id } = useParams()
    let [msg, setMessage] = useState()
    function SaveMessageData() {

    }
    return (
        <div className="border-2">
            <div className="d-flex justify-content-between align-items-center">
                <h1 className="text-left mx-2 rounded-lg text-white font-bold py-2">
                    <span>{messageType} messages</span>
                </h1>

                <div className="form-check form-switch">
                    <input className="form-check-input custom-control-input custom-switch" type="checkbox"
                           id="flexSwitchCheckDefault"/>
                    <label className="custom-control-label" htmlFor="flexSwitchCheckDefault"></label>
                </div>
            </div>
        </div>


    );
};

export default GuildMessageComponent;

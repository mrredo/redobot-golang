import Swal1 from 'sweetalert2'
import withReactContent from "sweetalert2-react-content";
// @ts-ignore
import Discord from "../../stuff/discord.svg"

let Swal = withReactContent(Swal1)
export function DiscordLoginPopup() {
    Swal.fire({
        title: "You must be logged in to view this Page!",
        html: <div className={"grid place-items-center"}><DiscordLoginBtnTest/></div>,
        showCancelButton: false,
        showConfirmButton: false,
        allowOutsideClick: false,
    })
}
export function DiscordLoginBtn() {
    return (
        <button onClick={() => DiscordLoginPopup()}>
            Test
        </button>
    )
}
export function DiscordLoginBtnTest() {
    return (
        <button
            onClick={() => {location.href = "/auth/login"}}
            style={{ /*backgroundColor: "#234de7",*/ borderColor: "#1e43de",  }}
            className="border-2 p-3 flex items-center text-white rounded-md font-bold transition-all duration-300 hover:bg-[#234de7] bg-[#000000]"
        >
            <img src={Discord} className="mr-2" alt="Discord Logo" />
            <span>Login with Discord</span>
        </button>
    )
}
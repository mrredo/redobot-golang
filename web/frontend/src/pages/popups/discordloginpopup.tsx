import Swal1 from 'sweetalert2'
import withReactContent from "sweetalert2-react-content";


let Swal = withReactContent(Swal1)
function DiscordLogin() {
    Swal.fire({
        title: "You must be logged in to view this Page!"

    })
}
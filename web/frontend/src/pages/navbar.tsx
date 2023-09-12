import { useState, useEffect } from 'react';
import { Navbar, Container, Nav, NavDropdown } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.min.css"
import "../styles/NavBar.css"

// @ts-ignore
import redobot from '../stuff/redobot.png';
import {DiscordLoginPopup} from "./popups/discordloginpopup";
interface Props {
  discordloginpopup?: boolean
  path?: string
  setuser?: (user: any) => void
}
export default function NavBar(props: Props) {
  const [user, setFetchedData] = useState({}) as any;
  let [logged, setLogged] = useState(false)

  useEffect(() => {
    const getData = async () => {
      const datas = await fetch("/api/user");
      let jsd = await datas.json()
      setFetchedData(jsd);
      setLogged(datas.status == 200)
      if (props.setuser) {
        props.setuser(jsd); // Make sure to replace userData with the actual user data
      }

      if(!(datas.status == 200) && props.discordloginpopup) {
        DiscordLoginPopup(/*(props.path == ""? "/" : props.path) as string*/)
      }
    };
    getData();
  }, []);

  return (
    <Navbar bg="dark" fixed="top" expand="lg">
      <Container>
        <Navbar.Brand href="/" className="text-white">
          <img
            src={redobot}
            width="30"
            height="30"
            className="d-inline-block align-top mr-1"
            alt="Redobot logo"
          />
          Redobot
        </Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link className="text-white" href="#home">About</Nav.Link>
            <Nav.Link className="text-white" href="#link">Documentations</Nav.Link>
            {logged? (
                <NavDropdown title={
                <span className="text-white">
                  {/*<img className="rounded-full" width="30" src={user.avatar.url} alt="profile picture"/>*/}
                  {user.username !== ""? user.username : "Loading..."}
                  </span>} id="basic-nav-dropdown">
                  <NavDropdown.Item className="text-white bg-gray-600" href="/guilds">Dashboard</NavDropdown.Item>
                  <NavDropdown.Item className="text-white bg-gray-600" href="/guilds">Your info</NavDropdown.Item>
                  <NavDropdown.Divider />
                  <NavDropdown.Item className="text-white" href="/auth/logout?redirect=/">Logout</NavDropdown.Item>
                </NavDropdown>
                ) : (
              <Nav.Link className="text-white" href="/auth/login">Login</Nav.Link>
              )
              }
            
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}
import React, { useEffect, useState } from "react";
import NavBar from "../navbar";
import {Spinner} from "../Spinner";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
// @ts-ignore
import Redobot from "../../stuff/redobot.png"
import TableElement from "./premiumpricingtable";
let description = `
This premium package offers for all servers alot of perks.
`

const PaymentPremiumPage = () => {
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        // Simulate a loading delay for demonstration purposes
        const loadingTimeout = setTimeout(() => {
            setIsLoading(false);
        }, 2000); // Adjust the delay as needed

        return () => clearTimeout(loadingTimeout);
    }, []);



    return (
        <>
            <NavBar />
            <div className="h-20" />
            <div className="grid sm:grid-cols-1 grid-cols-2 place-items-center">
                <div>
                    <span className={"text-2xl"}>€2.99/month</span>
                    <Card style={{ width: '20rem' }}>
                        <Card.Img variant="top my-2" src={Redobot} />
                        <Card.Body>
                            <Card.Title>Redobot Premium</Card.Title>
                            <Card.Text>
                                {description}
                            </Card.Text>
                            <Button onClick={() => {location.href = "/checkout"}} variant="primary">Subscribe</Button>
                        </Card.Body>
                    </Card>
                </div>
                <div>
                    <span className={"text-2xl"}>€28.99/year</span>
                    <Card style={{ width: '20rem' }}>
                        <Card.Img variant="top my-2" src={Redobot} />
                        <Card.Body>
                            <Card.Title>Redobot Premium</Card.Title>
                            <Card.Text>
                                Year subscription saves you 2 months off.
                                {description}
                            </Card.Text>
                            <Button onClick={() => {location.href = "/checkout?pr=year"}} variant="primary">Subscribe</Button>
                        </Card.Body>
                    </Card>
                </div>
            </div>
                <TableElement />

        </>
    );
};

export default PaymentPremiumPage;

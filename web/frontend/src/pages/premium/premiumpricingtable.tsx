import Table from 'react-bootstrap/Table';
import React from "react";

const TableElement = () => {
    return (
        <div className="w-[40vw] mx-auto my-10">
            <Table striped bordered hover>
                <thead>
                <tr>
                    <th>Pricing and Features</th>
                    <th>Default</th>
                    <th>Premium</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <td>Custom command count</td>
                    <td>10</td>
                    <td>50</td>
                </tr>

                </tbody>
            </Table>
        </div>
    );
}

export default TableElement;

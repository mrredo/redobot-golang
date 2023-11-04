import React, {useEffect, useState} from 'react';
import {Accordion} from "react-bootstrap"
function ColoredWordsEditor() {
    const [data, setData] = useState({});
    let [event, setEvent] = useState(0)
    useEffect(() => {

        setData({
            key1: {
                hello: "eeee",
                key: "beef",
                jj: {
                    eeee: 222,
                },
            },
            key2: "e",
            ffff: {
                hello: {
                    hello: {
                        hello: "eeee",
                        key: "beef",
                        jj: {
                            eeee: 222,
                        },
                    },
                    key: "beef",
                    jj: {
                        eeee: 222,
                    },
                },
                key: "beef",
                jj: {
                    eeee: 222,
                },
            },
        });
        // setData({
        //     key3: {
        //         ee: {
        //             keyg: "eeeee",
        //         },
        //     },
        //     key2: "3e"
        // })
    }, []);
    let eventKey=0
    function isObject(obj: any) {
        return typeof obj === "object" && !Array.isArray(obj);
    }

    function repeatString(str: string, times: number): string {
        if (times <= 0 || times === undefined) return ""
        return new Array(times + 1).join(str);
    }
    function AccordionMaker(items: any) {
        return (
            <Accordion defaultActiveKey="0-0-0">
                {items}
            </Accordion>
        )
    }
    function AccordionItemMaker(name: any, content: any, eventKey: string, withoutbody?: boolean) {
        return (
                <Accordion.Item eventKey={eventKey}>
                    <Accordion.Header>{name}</Accordion.Header>
                    {withoutbody? "" : <Accordion.Body>{content}</Accordion.Body>}
                </Accordion.Item>
        );
    }
    function DetailsAndSummary(name: any, content: any, nestingLevel: number) {
        nestingLevel = nestingLevel == 0 ? 0 : nestingLevel
        let color = nestingLevel % 2 == 0? "border-white" : "border-red-600"
        return (
            <>
                <details className={`border-2 ${color}`}>
                    <summary>{name}</summary>
                    {content}
                </details>
            </>
        );
    }

    function Normaltype(key: any, value: any) {
        return (
                <div key={key}>
                    {key}: {value}
                </div>
                );
    }


    function Load(ob: any): any[] {
        const content: any = []
        let keys = Object.keys(ob)
        for (const num in keys) {
            event += 1
            even++
            if (isObject(ob[keys[num]])) {
                content.push(AccordionItemMaker(keys[num], Load(ob[keys[num]]), `${even}`))
            } else {
                content.push(AccordionItemMaker(keys[num], "", `${even}`, true))
                // content.push(Normaltype(keys[num], ob[keys[num]]))
            }
        }
        return content
    }


    return (
        <div>
            <button onClick={() => setData(data)}>reload</button>
            {AccordionMaker(Load(data))}
        </div>
    );
}
let even = 1
export default ColoredWordsEditor;

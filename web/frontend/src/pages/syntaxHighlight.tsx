import React, {useEffect, useRef} from 'react';
let regex = /\{([^{}]+)}/gm
let NumRegex = /[0-9]+/gm
let SymbolsRegex = /(\|\||&&|==|!=|>=|<=|\+|-|\*|\/|%|&|\||\^|<<|>>|!)/gm
const SyntaxHighlight: React.FC = () => {
    const DivRef = useRef<HTMLDivElement>(null);
    const AreaRef = useRef<HTMLTextAreaElement>(null);
    const onInput = () => {
        let element = document.getElementById("input") as HTMLTextAreaElement
        let text = element.value
        let matches = text!.match(regex)
        let newText = text
        matches!.forEach((val, i) => {

            newText = newText.replace(SymbolsRegex, `<span style="color: #001ef8">$1</span>`)

             newText = newText.replace(val, `<span style="color: #ff0000">$1</span>`)
        })
        DivRef.current!.innerHTML = newText
    }
    useEffect(() => {
        onInput();
    }, []);
    return (
        <>
            <textarea id={"input"} onInput={onInput} ref={AreaRef}/>
            <div
                ref={DivRef}
            />
        </>
    );
};

export default SyntaxHighlight;
/*
* MAKE A check for current node length then if the required length is lower than this required we do previous child node length and we minus the required length
* correctCaretPos = false
* for !correctCaretPos {
*
* if current.length < requiredLength {
* requiredLength -= current.length
* current = NextNode()
* } else {
* correctCaretPos
* }
* }
*
*
* */
import React, {useRef} from 'react';

const SyntaxHighlight: React.FC = () => {
    const DivRef = useRef<HTMLDivElement>(null);
    const MainRef = useRef<HTMLDivElement>(null);

    const colorizeCode = (text: string): string => {
        return text.replace(/\{\{(.*?)\}\}/g, '{{<span style="color: #0dcaf0">$1</span>}}');
    };

    const OnInput = () => {
        if (DivRef.current && DivRef.current.textContent) {
            const NodePrevious = window.getSelection()?.focusNode

            const caret = window.getSelection()?.getRangeAt(0);
            const caretStart = caret?.startOffset || 0;
            const caretEnd = caret?.endOffset || 0;

            const savedCaretPosition = { start: caretStart, end: caretEnd };
            console.log(savedCaretPosition)
            const textContent = DivRef.current.textContent || '';
            DivRef.current.innerHTML = colorizeCode(textContent);
            const sel = window.getSelection();
            sel?.removeAllRanges();
            let node = DivRef.current!.firstChild as Node
            let correctCaretPos = false
            let offset = Math.min(savedCaretPosition.start, savedCaretPosition.end)
            while(!correctCaretPos) {
                let curl = node.textContent!.length
                if(node.textContent!.length < offset) {
                    node = node.nextSibling as Node
                } else {
                    correctCaretPos = true
                    continue
                }
                offset -= curl
            }

            if (DivRef.current!.childNodes.length > 1) {
                node = NodePrevious as Node
            }
            const newRange = document.createRange();
            const textNode = node as Node

            if (textNode.nodeType === Node.TEXT_NODE) {
                newRange.setStart(textNode, Math.min(offset, textNode.textContent!.length));
                newRange.setEnd(textNode, Math.min(offset, textNode.textContent!.length));
            }

            sel?.addRange(newRange);
        }
    };

    return (
        <div ref={MainRef}>
            <div
                contentEditable={true}
                onInput={OnInput}
                ref={DivRef}
            />
        </div>
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
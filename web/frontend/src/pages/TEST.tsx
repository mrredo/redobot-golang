import React, { Component, ChangeEvent } from 'react';

class ColoredWordsEditor extends Component {
    constructor(props: {}) {
        super(props);
        this.state = {
            content: [],
        };
    }

    handleContentChange = (e: ChangeEvent<HTMLDivElement>) => {
        const content = e.target.textContent || '';
        const coloredContent = this.colorWords(content);
        this.setState({ content: coloredContent });
    };

    colorWords = (text: string) => {
        // Split the text into words using a regular expression
        const words = text.split(/\s+/).filter((word) => word.trim() !== '');

        // Generate a random color for each word
        const coloredWords = words.map((word, index) => {
            const randomColor = this.getRandomColor();
            return (
                <span key={index} style={{ color: randomColor }}>
          {word}{' '}
        </span>
            );
        });

        return coloredWords;
    };

    getRandomColor = () => {
        const letters = '0123456789ABCDEF';
        let color = '#';
        for (let i = 0; i < 6; i++) {
            color += letters[Math.floor(Math.random() * 16)];
        }
        return color;
    };

    render() {
        return (
            <div
                contentEditable
                onInput={this.handleContentChange}
                style={{ border: '1px solid black', padding: '10px' }}
            >
                {(this.state as any  ).content}
            </div>
        );
    }
}

export default ColoredWordsEditor;

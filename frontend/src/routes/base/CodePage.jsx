import React from 'react';
import MonacoEditor from 'react-monaco-editor';

class CodeEditor extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            code: '// write your code...',
        };

        // 绑定方法
        this.editorDidMount = this.editorDidMount.bind(this);
        this.onChange = this.onChange.bind(this);
    }

    editorDidMount(editor, monaco) {
        // console.log('editorDidMount', editor);
        editor.focus();
    }

    onChange(newValue, e) {
        // console.log('onChange', newValue, e);
        this.setState({code: newValue});
    }

    render() {
        const {code} = this.state;
        const options = {
            selectOnLineNumbers: true
        };

        return (
            <MonacoEditor
                width="100%"
                height="100%"  // 或任何你希望的高度
                language="golang"
                theme="vs-dark"
                value={code}
                options={options}
                onChange={this.onChange}
                editorDidMount={this.editorDidMount}
            />
        );
    }
}

export default CodeEditor;

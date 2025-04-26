import React, {useState} from 'react'
import {Avatar, Popover} from 'antd'
import ChatWindow from './ChatWindow'
import '../css/ChatWindow.css'

const FloatingChatButton = () => {
    const [chatVisible, setChatVisible] = useState(false)
    const [connected, setConnected] = useState(false)

    const handleButtonClick = () => {
        setChatVisible(!chatVisible)
    }

    const chatFloatStyle = {
        right: 40,
        bottom: 30,
        position: 'fixed',
    }

    const dotStyle = {
        bottom: 30,
        width: '15x',
        height: '15px',
        position: 'fixed',
        borderRadius: '50%',
        display: 'inline-block',
        backgroundColor: connected ? 'red' : 'green',
    }

    const buttonContent = (
        <Popover>
            <div style={chatFloatStyle}>
                <Avatar
                    className="avatar"
                    size={64}
                    onClick={handleButtonClick}
                />
                <div style={dotStyle}></div>
                {chatVisible && (
                    <ChatWindow
                        onClose={() => setChatVisible(false)}
                        onConnect={() => setConnected(true)}
                        onDisconnect={() => setConnected(false)}
                    />
                )}
            </div>
        </Popover>
    )

    return <div>{buttonContent}</div>
}

export default FloatingChatButton

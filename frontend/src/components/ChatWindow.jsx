import React, {useEffect, useRef, useState} from 'react'
import {Avatar, Button, Input, List, message, Modal} from 'antd'
import {CheckCircleTwoTone, CloseCircleFilled, CloseCircleTwoTone} from '@ant-design/icons'
import html2canvas from 'html2canvas'
import ChangeConfig from "../components/ChatConfig"
import WS from "../shared/stream.js"
import '../css/ChatWindow.css'
import {SDB} from "../shared/token.js"

const {TextArea} = Input

const ChatWindow = ({onClose, onDisconnect}) => {
    const [messages, setMessages] = useState([])
    const [inputMessage, setInputMessage] = useState('')
    const [connected, setConnected] = useState(true)
    const [isModalVisible, setIsModalVisible] = useState(false)
    const [contextMenu, setContextMenu] = useState({visible: false, x: 0, y: 0})
    const messageContainerRef = useRef(null)

    // 设置滚动到底部函数
    const scrollToBottom = () => {
        if (messageContainerRef.current) {
            messageContainerRef.current.scrollTop = messageContainerRef.current.scrollHeight
        }
    }

    // 右键菜单的事件处理器
    const handleContextMenu = (event) => {
        event.preventDefault()
        setContextMenu({
            visible: true,
            x: event.clientX,
            y: event.clientY
        })
    }

    // 关闭右键菜单
    const closeContextMenu = () => {
        setContextMenu({visible: false, x: 0, y: 0})
    }


    const showChangeAIDialog = () => {
        setIsModalVisible(true)
    }

    const handleCancel = () => {
        setIsModalVisible(false)
    }


    const changeAI = () => {
        closeContextMenu()
        showChangeAIDialog()
    }

    const clearChat = () => {
        closeContextMenu()
        setMessages([])
        SDB.set('chatMessages', JSON.stringify([]))
    }

    const exportChatAsImage = () => {
        closeContextMenu()
        html2canvas(document.querySelector(".ant-list-items")).then(canvas => {
            // 创建一个图片元素并下载
            const image = canvas.toDataURL("image/png").replace("image/png", "image/octet-stream")
            const link = document.createElement('a')
            link.download = 'chat.png'
            link.href = image
            link.click()
        })
    }

    useEffect(() => {
        const handleDocumentClick = (event) => {
            // 检查点击事件是否发生在菜单外部
            if (contextMenu.visible && !event.target.closest('.context-menu')) {
                closeContextMenu()
            }
        }

        // 绑定事件监听器
        document.addEventListener('click', handleDocumentClick)

        // 清理事件监听器
        return () => {
            document.removeEventListener('click', handleDocumentClick)
        }
    }, [contextMenu.visible])


    useEffect(() => {
        // 从localStorage加载历史消息
        const savedMessages = SDB.get('chatMessages')
        if (savedMessages) {
            setMessages(JSON.parse(savedMessages))
        }
    }, [])

    useEffect(() => {
        // 设置连接成功回调
        WS.setChatCallback((msg) => handleAgentMsg(msg))
        WS.setTaskCallback((msg) => handleAgentMsg(msg))

        // 设置连接断开回调
        WS.setLogCallback(() => setConnected(false))

    }, [onDisconnect, messages])

    const handleUserSendMsg = () => {
        if (inputMessage.trim() !== '') {
            const newMessage = {
                text: inputMessage,
                fromUser: true,
            }
            const newMessages = [...messages, newMessage]
            setMessages(newMessages)
            SDB.set('chatMessages', JSON.stringify(newMessages)) // 使用SimpleDB保存消息

            setInputMessage('')
            ParseSendMessage(inputMessage)
        }
    }

    // 处理代理发送的消息
    const handleAgentMsg = (msg) => {
        const chatMsg = {
            text: msg,
            fromUser: false,
        }
        const newMessages = [...messages, chatMsg]
        setMessages(newMessages)
        SDB.set('chatMessages', JSON.stringify(newMessages)) // 使用SimpleDB保存消息

        setTimeout(() => {
            scrollToBottom() // 在代理发送消息后触发滚动
        }, 0)
    }

    const handleKeyDown = (e) => {
        if (e.key === 'Enter' && e.ctrlKey) {
            handleUserSendMsg()
            scrollToBottom()
        }
    }

    const copyMessageToClipboard = async (text) => {
        try {
            await navigator.clipboard.writeText(text)
            message.success('消息已复制')
        } catch (err) {
            console.error('复制失败', err)
        }
    }

    let chatName = 'Kbot'

    return (
        <div className="chat-window">
            <div className="chat-header">
                <div style={{flex: 1}}>
                    <Avatar className="avatar-ct" size={64}/>
                </div>
                <div style={{flex: 3}}>
                    <span className="chat-username">{chatName}</span>
                    <ConStatus flag={connected}/>
                </div>
                <Button type='default' onClick={onClose} icon={<CloseCircleFilled/>}/>
            </div>
            <div className="message-container" ref={messageContainerRef} onContextMenu={handleContextMenu}>
                <List
                    dataSource={messages}
                    renderItem={(message, index) => (
                        <List.Item
                            key={index}
                            className={message.fromUser ? 'user' : 'agent'}
                            onDoubleClick={() => copyMessageToClipboard(message.text)} // 添加双击处理器
                        >
                            <div className='chat-line-box-style'>
                                {message.text}
                            </div>
                        </List.Item>
                    )}
                />

            </div>
            <div className="chat-input">
                <TextArea
                    placeholder="输入你的消息..."
                    value={inputMessage}
                    onChange={(e) => setInputMessage(e.target.value)}
                    onKeyDown={handleKeyDown}
                    autoSize={{minRows: 1, maxRows: 6}}
                />
                <Button
                    type="primary"
                    shape='round'
                    onClick={handleUserSendMsg}
                    disabled={inputMessage.trim() === ''}
                    className={inputMessage.trim() !== '' ? 'btn-active' : ''}
                >
                    发送
                </Button>

                {/* 右键菜单 */}
                <div
                    className={`context-menu ${contextMenu.visible ? 'active' : ''}`}
                    style={{top: contextMenu.y, left: contextMenu.x}}
                    onClick={closeContextMenu}
                >
                    <div className="context-menu-item" onClick={changeAI}>更换GPT</div>
                    <div className="context-menu-item" onClick={clearChat}>清空聊天</div>
                    <div className="context-menu-item" onClick={exportChatAsImage}>导出图片</div>
                </div>
            </div>
            {/* 更换AI的模态框 */}
            <Modal
                title="更换AI配置"
                style={{textAlign: "center"}}
                open={isModalVisible}
                onCancel={handleCancel}
                footer={null}
            >
                <ChangeConfig onClose={handleCancel}/>
            </Modal>
        </div>
    )
}

const ConStatus = ({flag}) => {
    return flag ? (
        <span className="chat-status">
            <CheckCircleTwoTone twoToneColor="#52c41a"/> 在线
        </span>
    ) : (
        <span className="chat-status">
            <CloseCircleTwoTone twoToneColor="#eb2f96"/> 不在线
        </span>
    )
}


const ParseSendMessage = (message) => {
    let sendMsg = {
        type: 'chat',
        payload: message,
    }

    let extract = (str) => {
        const regex = /\/task\s+(.*)/
        const matches = regex.exec(str);
        return matches && matches.length > 1 ? matches[1] : null
    }

    let info = extract(message)
    if (info !== null) {
        console.log(info)
        sendMsg.type = 'task'
        sendMsg.payload = info
    }else{
        sendMsg.payload = message
    }

    // 发送任务信息
    WS.socket.send(JSON.stringify(sendMsg))
}

export default ChatWindow

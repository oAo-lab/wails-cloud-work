class WebSocketClient {
    constructor(url) {
        this.url = url
        this.socket = null
        this.heartbeatInterval = 3000 // 心跳间隔，默认为 3 秒
        this.logCallback = null
        this.chatCallback = null
        this.taskCallback = null
        this.heartbeatTimer = null
    }

    connect() {
        this.socket = new WebSocket(this.url)

        this.socket.onopen = () => {
            console.log('WebSocket connection opened.')
            this.startHeartbeat()
        }

        this.socket.onmessage = (event) => {
            const message = JSON.parse(event.data)

            switch (message.type) {
                case 'log':
                    if (this.logCallback) {
                        this.logCallback(message.payload)
                    }
                    break
                case 'chat':
                    if (this.chatCallback) {
                        this.chatCallback(message.payload)
                    }
                    break
                case 'task':
                    if (this.taskCallback) {
                        this.taskCallback(message.payload)
                    }
                    break
                default:
                    break
            }
        }

        this.socket.onclose = () => {
            console.log('WebSocket connection closed.')
            this.stopHeartbeat()
        }
    }

    // 日志回调
    setLogCallback(callback) {
        this.logCallback = callback
    }

    // 聊天回调
    setChatCallback(callback) {
        this.chatCallback = callback
    }

    // 任务回调
    setTaskCallback(callback) {
        this.taskCallback = callback
    }

    // 开始心跳
    startHeartbeat() {
        this.heartbeatTimer = setInterval(() => {
            this.sendPing()
        }, this.heartbeatInterval)
    }

    // 停止心跳
    stopHeartbeat() {
        if (this.heartbeatTimer) {
            clearInterval(this.heartbeatTimer)
        }
    }

    // 设置ping
    sendPing() {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.send(JSON.stringify({type: 'ping', payload: 'pong'}))
        }
    }

    // 关闭连接
    close() {
        if (this.socket) {
            this.socket.close()
        }
    }
}

const WS = new WebSocketClient("ws://localhost:7001/api/v1/ws")

// 连接WebSocket
WS.connect()

export default WS

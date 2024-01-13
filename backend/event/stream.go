package event

import (
	"could-work/backend/core/define"
	"could-work/backend/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketMessage struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func MonitorWS(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		define.Log.Error("WebSocket upgrade error:", err)
		return
	}

	defer func(conn *websocket.Conn) {
		_ = conn.Close()
	}(conn)

	// 定义心跳间隔
	heartbeatInterval := 3 * time.Second

	// 创建一个定时器
	heartbeatTimer := time.NewTicker(heartbeatInterval)

	for {
		select {
		case <-heartbeatTimer.C:
			if err := conn.WriteJSON(WebSocketMessage{Type: "ping"}); err != nil {
				define.Log.Error("WebSocket write error:", err)
				return
			}

		default:
			if !util.MsgQueue.IsEmpty() {
				message := util.MsgQueue.Pop()
				if message.Type != "ping" {
					if err := conn.WriteJSON(&message); err != nil {
						define.Log.Error("WebSocket write error:", err)
						return
					}
				}
			}

			var receive *WebSocketMessage

			if err := conn.ReadJSON(&receive); err != nil {
				define.Log.Error("WebSocket message parsing error:", err)
				continue
			}

			switch receive.Type {
			case "ping":
				handlePing(receive)
			case "chat":
				handleChat(receive)
			case "task":
				handleTask(receive)
			default:
				define.Log.Warn("Unknown message type:", receive.Type)
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}

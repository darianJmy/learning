package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理WebSocket连接请求
func websocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to upgrade to WebSocket: "+err.Error())
		return
	}
	// 在新协程中处理WebSocket
	go func(conn *websocket.Conn) {
		defer conn.Close()
		for {
			// 读取WebSocket消息
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// 处理WebSocket消息
			// 发送WebSocket消息
			err = conn.WriteMessage(websocket.TextMessage, []byte("Received: "+string(message)))
			if err != nil {
				return
			}
		}
	}(conn)
}
func main() {
	r := gin.Default()
	r.GET("/websocket", websocketHandler)
	r.Run()
}

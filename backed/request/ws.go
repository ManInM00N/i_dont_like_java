package request

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//	func WsInit() {
//		R.
//	}
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(c *gin.Context) {
	// 升级 HTTP 连接为 WebSocket 连接
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	for {
		// 从 WebSocket 连接读取客户端发送的消息
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// 处理客户端的请求并调用 LL Model
		userInput := string(msg)
		llmResponse := callLLModel(userInput)

		// 将 LL Model 的响应发送给客户端
		err = ws.WriteMessage(websocket.TextMessage, []byte(llmResponse))
		if err != nil {
			log.Println(err)
			break
		}
	}
}

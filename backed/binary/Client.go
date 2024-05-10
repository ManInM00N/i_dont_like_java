package binary

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Client struct {
	ID         string
	UserId     interface{}
	Socket     *websocket.Conn
	Send       chan []byte
	Start      time.Time
	ExpireTime time.Duration // 一段时间没有接收到心跳则过期
}

type ClientManager struct {
	Clients    map[string]*Client // 记录在线用户
	Broadcast  chan []byte        //触发消息广播
	Register   chan *Client       // 触发新用户登陆
	UnRegister chan *Client       // 触发用户退出
}
type WsMessage struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

var (
	Manager ClientManager
)

// Read 读取客户端发送过来的消息
func (c *Client) Read() {
	// 出现故障后把当前客户端注销
	defer func() {
		_ = c.Socket.Close()
		Manager.UnRegister <- c
	}()
	for {
		_, data, err := c.Socket.ReadMessage()
		if err != nil {
			log.Fatalln(err)
			break
		}
		var msg WsMessage
		err = json.Unmarshal(data, &msg)
		if err != nil {
			log.Println(err)
			break
		}

		switch msg.Type {
		case 6:
			// 如果是心跳监测消息（利用心跳监测来判断对应客户端是否在线）
			resp, _ := json.Marshal(&WsMessage{Type: 6, Data: "pong"})
			c.Start = time.Now() // 重新刷新时间
			c.Send <- resp
		case 1:
			// 获取在线人数
			count := len(Manager.Clients)
			resp, _ := json.Marshal(&WsMessage{Type: 1, Data: count})
			c.Send <- resp
		case 2:
			// 获取消息历史记录
			_data := ChatRecord() //你的获取消息记录的操作
			resp, _ := json.Marshal(&WsMessage{Type: 2, Data: _data})
			c.Send <- resp
		case 3:
			// 发送文本消息
			resp, _ := json.Marshal(&WsMessage{Type: 3, Data: msg.Data})
			Manager.Broadcast <- resp

		case 4:
			// 你的撤回消息的操作
			c.Send <- []byte("回复消息")

		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
		Manager.UnRegister <- c
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				err := c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Println(err)
					return
				}
				return
			}
			err := c.Socket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (c *Client) Check() {
	for {
		now := time.Now()
		var duration = now.Sub(c.Start)
		if duration >= c.ExpireTime {
			Manager.UnRegister <- c
			break
		}
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-Manager.Register:
			Manager.Clients[conn.ID] = conn
			// 如果有新用户连接则发送最近聊天记录和在线人数给他
			count := len(Manager.Clients)
			Manager.InitSend(conn, count)
		}
	}
}
func (manager *ClientManager) InitSend(cur *Client, count int) {
	resp, _ := json.Marshal(&WsMessage{Type: 1, Data: count})
	Manager.Broadcast <- resp

	_data := YouChatHistoryList() //获取聊天室历史消息记录操作
	resp, _ = json.Marshal(&WsMessage{Type: 2, Data: _data})
	cur.Send <- resp
}
func (manager *ClientManager) BroadcastSend() {
	for {
		select {
		case msg := <-Manager.Broadcast:
			for _, conn := range Manager.Clients {
				conn.Send <- msg
			}
		}
	}
}
func (manager *ClientManager) Quit() {
	for {
		select {
		case conn := <-Manager.UnRegister:
			delete(Manager.Clients, conn.ID)
			// 给客户端刷新在线人数
			resp, _ := json.Marshal(&WsMessage{Type: 1, Data: len(Manager.Clients)})
			manager.Broadcast <- resp
		}
	}
}
func WebSocketHandle(c *gin.Context) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		log.Println(err)
		return
	}
	userid := c.Values["a_userid"]
	ip := c.ClientIP()
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		log.Println(err)
		return
	}
	ua := c.GetHeader("User-Agent")
	id := ip + ua
	idMd5 := fmt.Sprintf("%x", md5.Sum([]byte(id)))
	client := &Client{
		ID:     idMd5,
		Socket: conn, Send: make(chan []byte),
		UserId:     userid,
		Start:      time.Now(),
		ExpireTime: time.Minute * 1,
	}
	Manager.Register <- client
	go client.Read() // 以goroutine的方式调用Client的Read、Write、Check方法
	go client.Write()
	go client.Check()
}

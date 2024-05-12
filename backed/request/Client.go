package request

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"main/binary"
	"strings"
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
type ReturnMessage struct {
	Name  string `json:"name"`
	Inner string `json:"inner"`
}

var (
	Manager = &ClientManager{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte, 2000),
		Register:   make(chan *Client, 2000),
		UnRegister: make(chan *Client, 2000),
	}
)

func ChatHistoryList() []Send {
	var res []Send
	db.Find(&res)
	binary.InfoLog.Println(res)
	return res
}

func (c *Client) Read() {
	defer func() {
		_ = c.Socket.Close()
		Manager.UnRegister <- c
	}()
	for {
		_, data, err := c.Socket.ReadMessage()
		//var data WsMessage
		//err := c.Socket.ReadJSON(&data)

		if err != nil {
			if strings.Contains(err.Error(), "websocket: close 1005") || strings.Contains(err.Error(), "websocket: close 1001") {
				err = nil
				binary.InfoLog.Println(c.UserId, " close connection")
				return
			}
			binary.DebugLog.Fatalln(err)
			break
		}
		var msg WsMessage
		err = json.Unmarshal(data, &msg)
		var newMsg Send
		binary.InfoLog.Println(msg.Data.(map[string]interface{}))
		newMsg.Name = (msg.Data.(map[string]interface{}))["name"].(string)
		newMsg.Inner = (msg.Data.(map[string]interface{}))["inner"].(string)
		err = db.Create(&newMsg).Error
		if err != nil {
			binary.DebugLog.Println(err)
			break
		}
		switch msg.Type {
		case 6:
			resp, _ := json.Marshal(&WsMessage{Type: 6, Data: "pong"})
			c.Start = time.Now()
			c.Send <- resp
		case 1:
			count := len(Manager.Clients)
			resp, _ := json.Marshal(&WsMessage{Type: 1, Data: count})
			c.Send <- resp
		case 2:
			_data := ChatHistoryList()
			resp, _ := json.Marshal(&WsMessage{Type: 2, Data: _data})
			c.Send <- resp
		case 3:
			resp, _ := json.Marshal(&WsMessage{Type: 3, Data: msg.Data})

			Manager.Broadcast <- resp
		case 4:
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
					binary.DebugLog.Println(err)
					return
				}
				return
			}
			err := c.Socket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				binary.DebugLog.Println(err)
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
			count := len(Manager.Clients)
			Manager.InitSend(conn, count)
		}
	}
}
func (manager *ClientManager) InitSend(cur *Client, count int) {
	resp, _ := json.Marshal(&WsMessage{Type: 1, Data: count})
	//Manager.Broadcast <- resp
	cur.Send <- resp
	_data := ChatHistoryList()
	ndata := make([]ReturnMessage, 0)
	for _, v := range _data {
		ndata = append(ndata, ReturnMessage{
			Inner: v.Inner,
			Name:  v.Name,
		})
	}
	resp, _ = json.Marshal(&WsMessage{Type: 2, Data: ndata})
	cur.Send <- resp
}
func (manager *ClientManager) BroadcastSend() {
	for {
		select {
		case msg := <-Manager.Broadcast:
			var tt WsMessage
			err := json.Unmarshal(msg, &tt)
			if err != nil {
				binary.DebugLog.Println(err)
			}
			binary.InfoLog.Println("start broadcast", len(Manager.Clients))

			for _, conn := range Manager.Clients {
				conn.Send <- msg
				binary.InfoLog.Println("post to", conn.UserId.(string))
			}
		}
	}
}
func (manager *ClientManager) Quit() {
	for {
		select {
		case conn := <-Manager.UnRegister:
			delete(Manager.Clients, conn.ID)
			resp, _ := json.Marshal(&WsMessage{Type: 1, Data: len(Manager.Clients)})
			manager.Broadcast <- resp
		}
	}
}

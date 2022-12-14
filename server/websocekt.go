package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yuyuancha/chatroom-service/model"
	"github.com/yuyuancha/chatroom-service/redis"
	"github.com/yuyuancha/chatroom-service/service"
)

var messageRedis redis.Message
var messageChannel = make(chan model.Message)

func handleWebsocket(c *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	name := c.Param("name")
	if name == "" {
		name = c.ClientIP()
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Websocket 建立錯誤。err:", err.Error())

		return
	}

	defer func() {
		_ = ws.Close()
	}()

	oldMessages := messageRedis.Get()
	_ = ws.WriteJSON(oldMessages)

	go ReadMessage(ws, name)
	SendMessage(ws)
}

func ReadMessage(ws *websocket.Conn, name string) {
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			return
		}

		messagePointer := service.MapMessageByWebsocket(name, string(data))
		if messagePointer == nil {
			continue
		}

		message := *messagePointer
		messageRedis.Create(message)

		messageChannel <- message
		<-messageChannel
	}
}

func SendMessage(ws *websocket.Conn) {
	for message := range messageChannel {
		_ = ws.WriteJSON(message)
		messageChannel <- message
	}
}

package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yuyuancha/chatroom-service/redis"
	"github.com/yuyuancha/chatroom-service/service"
)

var messageRedis redis.Message

func handleWebsocket(c *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
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

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			return
		}

		messagePointer := service.MapMessageByWebsocket(string(data))
		if messagePointer == nil {
			continue
		}

		fmt.Println(*messagePointer)
		_ = ws.WriteJSON(struct {
			Reply string `json:"rep2ly"`
		}{
			Reply: "Echo...",
		})
	}
}

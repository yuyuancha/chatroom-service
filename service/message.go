package service

import (
	"encoding/json"
	"fmt"

	"github.com/yuyuancha/chatroom-service/model"
	"github.com/yuyuancha/chatroom-service/redis"
)

var messageRedis redis.Message

func MapMessageByWebsocket(data string) *model.Message {
	var message model.Message

	err := json.Unmarshal([]byte(data), &message)
	if err != nil {
		fmt.Println("websocket 資料轉換結構錯誤。err:", err.Error())

		return nil
	}

	return &message
}

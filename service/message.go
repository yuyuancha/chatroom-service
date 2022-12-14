package service

import (
	"time"

	"github.com/yuyuancha/chatroom-service/model"
	"github.com/yuyuancha/chatroom-service/redis"
)

var messageRedis redis.Message

func MapMessageByWebsocket(name, content string) *model.Message {
	var message model.Message

	message.Author = name
	message.Content = content
	message.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	return &message
}

package redis

import (
	"encoding/json"
	"fmt"

	"github.com/yuyuancha/chatroom-service/config"
	"github.com/yuyuancha/chatroom-service/model"
)

type Message struct{}

func (Message) Create(message model.Message) {
	envConfig := config.GetEnvConfig()

	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("message 轉換結構錯誤。err:", err.Error())

		return
	}

	index, err := redisClient.RPush(ctx, keyMessages, data).Result()
	if err != nil {
		fmt.Println("建立訊息錯誤。err:", err.Error())
	}

	if index > envConfig.GetInt64("message_save_amount") {
		redisClient.LPop(ctx, keyMessages)
	}
}

func (Message) Get() []model.Message {
	data, err := redisClient.LRange(ctx, keyMessages, -5, -1).Result()
	if err != nil {
		fmt.Println("取得訊息錯誤。err:", err.Error())

		return nil
	}

	messages := make([]model.Message, 0, len(data))
	for i := 0; i < len(data); i++ {
		var message model.Message

		_ = json.Unmarshal([]byte(data[i]), &message)

		messages = append(messages, message)
	}

	return messages
}

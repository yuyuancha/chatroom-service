package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/yuyuancha/chatroom-service/config"
)

var ctx = context.Background()
var redisClient *redis.Client

func init() {
	envConfig := config.GetEnvConfig()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     envConfig.GetString("redis.host") + ":" + envConfig.GetString("redis.port"),
		Password: envConfig.GetString("redis.password"),
		DB:       envConfig.GetInt("redis.db"),
	})
}

func Pong() {
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("連線 redis 失敗: %s", err.Error())
	}

	fmt.Println("Pong:", pong)
}

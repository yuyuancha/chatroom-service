package main

import (
	"fmt"

	"github.com/yuyuancha/chatroom-service/config"
	"github.com/yuyuancha/chatroom-service/redis"
)

func init() {
	printAPPInfo()
}

func main() {
	fmt.Println("Hello world!")

	redis.Pong()
}

func printAPPInfo() {
	c := config.GetEnvConfig()
	fmt.Printf("APP information:\nname: %s\nenviroment: %s\n\n", c.GetString("app.name"), c.GetString("app.enviroment"))
}

package main

import (
	"fmt"

	"github.com/yuyuancha/chatroom-service/config"
)

func init() {
	printAPPInfo()
}

func main() {
	fmt.Println("Hello world!")
}

func printAPPInfo() {
	c := config.GetEnvConfig()
	fmt.Printf("APP information:\nname: %s\nenviroment: %s\n\n", c.GetString("app.name"), c.GetString("app.enviroment"))
}

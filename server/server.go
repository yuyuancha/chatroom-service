package server

import "github.com/gin-gonic/gin"

func SetServer() {
	server := gin.New()

	server.Use(gin.Recovery())
	_ = server.SetTrustedProxies(nil)

	server.GET("/messages/:name", handleWebsocket)

	_ = server.Run(":80")
}

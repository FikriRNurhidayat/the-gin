package main

import (
	"gin-thing/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.CustomRecovery(controller.Recover))

	r.GET("/", controller.Main)
	r.POST("/v1/login", controller.Login)
	r.NoRoute(controller.Default)

	r.Run()
}

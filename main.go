package main

import (
	"gin-thing/controller"
	"gin-thing/model/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.CustomRecovery(controller.Recover))

	r.GET("/", controller.Main)
	r.POST("/v1/login", controller.Login)
	r.POST("/v1/register", controller.Register)
	r.POST("/v1/refresh", controller.Refresh)
	r.NoRoute(controller.Default)

	database.Init()
	r.Run()
}

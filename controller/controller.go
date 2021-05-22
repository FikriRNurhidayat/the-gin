package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Main(c *gin.Context) {
	c.JSON(http.StatusOK, ok("I'm okay, thanks!", nil))
}

func Default(c *gin.Context) {
	c.JSON(
		http.StatusNotFound,
		fail("Route not found!",
			gin.H{
				"method": c.Request.Method,
				"url":    c.Request.RequestURI,
			}),
	)
}

func Recover(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		c.JSON(http.StatusInternalServerError, error(err))
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}

func fail(message string, data interface{}) gin.H {
	return gin.H{
		"status":  "FAIL",
		"message": message,
		"data":    data,
	}
}

func ok(message string, data interface{}) gin.H {
	return gin.H{
		"status":  "OK",
		"message": message,
		"data":    data,
	}
}

func error(message string) gin.H {
	return gin.H{
		"status":  "ERROR",
		"message": message,
	}
}

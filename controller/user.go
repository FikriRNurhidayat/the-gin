package controller

import (
	m "gin-thing/model"
	s "gin-thing/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body m.Credential
	c.BindJSON(&body)

	token, err := s.Login(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, fail(err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, ok("Token created!", token))
}

func Refresh(c *gin.Context) {
	var body m.Refresh
	c.BindJSON(&body)

	token, err := s.Refresh(body.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, fail(err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, ok("Token refreshed!", token))
}

func Register(c *gin.Context) {
	var body m.Credential
	c.BindJSON(&body)

	user, err := s.Register(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, fail(err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, ok("Registered!", user))
}

package controller

import (
	m "gin-thing/model"
	s "gin-thing/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var body m.Credential
	var token m.Token

	c.BindJSON(&body)

	token, err := s.Login(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, fail(err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, ok("Token created!", token))
}

package helper

import (
	"errors"
	"fmt"
	"gin-thing/config"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	var payload *jwt.StandardClaims

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.SIGNING_SECRET, nil
	})

	if err != nil {
		return payload, err
	}

	if !token.Valid {
		return payload, errors.New("token is not valid")
	}

	payload = token.Claims.(*jwt.StandardClaims)

	fmt.Println(payload)

	return payload, nil
}

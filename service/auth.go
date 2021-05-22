package service

import (
	"errors"
	"fmt"
	"time"

	config "gin-thing/config"
	"gin-thing/helper"
	m "gin-thing/model"
	mdb "gin-thing/model/database"

	"github.com/dgrijalva/jwt-go"
)

func Login(username string, password string) (m.Token, error) {
	var u mdb.User
	var token m.Token
	user := u.SelectOneByUsername(username)

	if !user.Check(password) {
		return token, errors.New("wrong password")
	}

	return createToken(*user), nil
}

func Refresh(tokenString string) (m.Token, error) {
	var u mdb.User
	var token m.Token
	payload, err := helper.ParseToken(tokenString)
	if err != nil {
		return token, err
	}
	user := u.SelectOneByID(payload.Id)
	token = createToken(*user)
	return token, nil
}

func createToken(user mdb.User) m.Token {
	var token m.Token

	token.ExpiredAt = time.Now().Add(time.Minute * 15).Unix()
	token.AccessToken = generateToken(user, config.SIGNING_SECRET, token.ExpiredAt)
	token.RefreshToken = generateToken(user, config.SIGNING_SECRET, time.Now().Add(time.Hour*168).Unix())

	return token
}

func generateToken(payload mdb.User, secret []byte, expiresAt int64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        fmt.Sprint(payload.ID),
		Subject:   payload.Username,
		ExpiresAt: expiresAt,
	})

	tokenString, err := token.SignedString(secret)
	fmt.Println(err)
	return tokenString
}

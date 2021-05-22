package service

import (
	"errors"
	"fmt"
	"time"

	config "gin-thing/config"
	m "gin-thing/model"
	mdb "gin-thing/model/database"

	jwt "github.com/dgrijalva/jwt-go"
)

func Login(username string, password string) (m.Token, error) {
	var u mdb.User
	var token m.Token
	var expires_in_minutes time.Duration = 15

	user := u.SelectOne(username)

	if !user.Check(password) {
		return token, errors.New("wrong password")
	}

	token.ExpiredAt = generateExpirationTime(time.Minute * expires_in_minutes)
	token.AccessToken = generateToken(user, config.SIGNING_SECRET, token.ExpiredAt)
	token.RefreshToken = generateToken(user, config.SIGNING_SECRET, generateExpirationTime(time.Hour*168))

	return token, nil
}

func generateToken(payload *mdb.User, secret []byte, expires_at int64) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        fmt.Sprint(payload.ID),
		Subject:   payload.Username,
		ExpiresAt: expires_at,
	})

	tokenString, err := token.SignedString(secret)
	fmt.Println(err)
	return tokenString
}

func generateExpirationTime(expires_in time.Duration) int64 {
	return time.Now().Add(expires_in).Unix()
}

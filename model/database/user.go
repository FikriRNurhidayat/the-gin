package database

import "fmt"

type User struct {
	ID                uint64 `json:"id"`
	Username          string `json:"username"`
	EncryptedPassword string `json:"encrypted_password"`
}

var user User = User{
	ID:                1,
	Username:          "fain",
	EncryptedPassword: "123456",
}

func (u User) SelectOneByUsername(username string) *User {
	if username != user.Username {
		return nil
	}
	return &user
}

func (u User) SelectOneByID(id string) *User {
	if id != fmt.Sprint(user.ID) {
		return nil
	}
	return &user
}

func (u User) Check(password string) bool {
	return user.EncryptedPassword == password
}

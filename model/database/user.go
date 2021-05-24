package database

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username          string `json:"username" gorm:"uniqueIndex"`
	EncryptedPassword string `json:"encrypted_password" gorm:"not null"`
}

func (u User) SelectOneByUsername(username string) User {
	var user User
	db.Where("username = ?", username).First(&user)
	return user
}

func (u User) SelectOneByID(id string) User {
	var user User
	db.First(&user, id)
	return user
}

func (u User) Create(username string, encryptedPassword string) (User, *gorm.DB) {
	u.Username = strings.ToLower(username)
	u.EncryptedPassword = encryptedPassword
	trx := db.Create(&u)
	return u, trx
}

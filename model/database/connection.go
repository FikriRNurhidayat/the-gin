package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

func Init() {
	db.AutoMigrate(&User{})
}

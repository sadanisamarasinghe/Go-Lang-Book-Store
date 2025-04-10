package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(sqlite.Open("books.db"),&gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database!")
	}
	DB = database
}
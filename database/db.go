package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("go_gin_app.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected successfully")

	Db = db
}

func GetDb() *gorm.DB {
	return Db
}

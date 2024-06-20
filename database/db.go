package database

import (
	"fmt"
	"log"

	"github.com/hazitgi/go_gin_server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	var err error
	DB, err = gorm.Open(sqlite.Open("go_gin_app.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected successfully")

	if err := DB.AutoMigrate(&models.User{}, &models.Skill{}, &models.SkillGroup{}, &models.UserSkillRank{}); err != nil {
		log.Fatalln("failed to migrate tables")
	}
}

func GetDb() *gorm.DB {
	if DB == nil {
		log.Fatalln("Db connection isn't initialized!")
	}
	return DB
}

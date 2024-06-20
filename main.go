package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/hazitgi/go_gin_server/apis"
	"github.com/hazitgi/go_gin_server/database"
	"github.com/hazitgi/go_gin_server/handlers"
	"github.com/hazitgi/go_gin_server/managers"
	"github.com/gin-contrib/static"
)

func main() {
	// Create a gin router
	router := gin.Default()

	_ = godotenv.Load(".env")
	env := os.Getenv("ENV")
	if env == "development" {
		gin.SetMode(gin.DebugMode)
		router.Use(cors.Default())
	} else {
		gin.SetMode(gin.ReleaseMode)
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"https://hazitgi.github.io/", "http://localhost:8000"}
		router.Use(cors.Default())
	}

	// initioalize database
	initDB()

	router.Use(static.Serve("/", static.LocalFile("views/build", false)))
	router.Use(static.Serve("/login", static.LocalFile("views/build", false)))
	router.Use(static.Serve("/skill-groups/:groupId", static.LocalFile("views/build", false)))
	router.Use(static.Serve("/skill-groups", static.LocalFile("views/build", false)))
	router.Use(static.Serve("/skills/:skillId", static.LocalFile("views/build", false)))
	router.Use(static.Serve("/users", static.LocalFile("views/build", false)))
	router.Use(static.Serve("/users/*", static.LocalFile("views/build", false)))
	router.Use(static.Serve("/skills", static.LocalFile("views/build", false)))
	router.Use(static.Serve("/home", static.LocalFile("views/build", false)))

	userManager := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManager)
	userHandler.RegisterUserRoutes(router)

	if err := router.Run(":8000"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}

func initDB() {
	database.Initialize()
}

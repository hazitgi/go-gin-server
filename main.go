package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/hazitgi/go_gin_server/apis"
	"github.com/hazitgi/go_gin_server/database"
	"github.com/hazitgi/go_gin_server/handlers"
	"github.com/hazitgi/go_gin_server/managers"
)

func main() {
	// Create a gin router
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	// init CORS middleware
	// Configure CORS middleware
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Replace with your frontend URL
        AllowMethods:     []string{"GET, POST, PATCH, PUT, DELETE, OPTIONS, HEAD"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	// initioalize database
	initDB()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"version": "v1.0.0",
		})
	})

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

package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hazitgi/go_gin_server/database"
)

func main() {
	// Create a gin router
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	// init CORS middleware
	router.Use(cors.Default())

	// initioalize database
	initDB()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"version": "v1.0.0",
		})
	})

	if err := router.Run(":8000"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}

func initDB() {
	database.ConnectDB()
}

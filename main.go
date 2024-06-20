package main

import (
	"log"
	// "time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/hazitgi/go_gin_server/apis"
	"github.com/hazitgi/go_gin_server/database"
	"github.com/hazitgi/go_gin_server/handlers"
	"github.com/hazitgi/go_gin_server/managers"
	// "github.com/hazitgi/go_gin_server/middleware"
)

func main() {
	// Create a gin router
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	// router.Use(middleware.ConfigureCORSMiddleware())
	// init CORS middleware
	router.Use(cors.Default())
	// router.Use(cors.New(cors.Config{
    //     AllowOrigins:     []string{"*"}, // Replace with your frontend URL
    //     AllowMethods:     []string{"GET, POST, PATCH, PUT, DELETE, OPTIONS, HEAD"},
    //     AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
    //     ExposeHeaders:    []string{"Content-Length"},
    //     AllowCredentials: true,
    //     MaxAge:           12 * time.Hour,
    // }))

	// router.OPTIONS("/", func(c *gin.Context) {
	// 	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	// 	c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	// 	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
	// 	c.Header("Access-Control-Allow-Credentials", "true")
	// 	c.JSON(200, gin.H{"message": "OPTIONS request handled"})
	// })
	

	// initioalize database
	initDB()

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

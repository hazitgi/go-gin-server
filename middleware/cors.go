package middleware

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureCORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Referrer-Policy"},
		ExposeHeaders:    []string{"Content-Length", "Referrer-Policy"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func ConfigureCORSMiddleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		fmt.Println(origin, ">>>>>>>>>>>>>>>")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{"message": "OPTIONS request handled"})
			return
		}

		c.Next()
	}
}

package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/hazitgi/go_gin_server/models"
)

func listUser(ctx *gin.Context) {
	user1 := models.User{
		FullName: "Mohamed Haseeb",
		Email:    "hazitgi@gmail.com",
	}
	user2 := models.User{
		FullName: "Mohamed Haseeb",
		Email:    "hazitgi@gmail.com",
	}
	user3 := models.User{
		FullName: "Mohamed Haseeb",
		Email:    "hazitgi@gmail.com",
	}
	ctx.JSON(200, []models.User{user1, user2, user3})
}

func createUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ok",
		"status":  "success",
	})
}

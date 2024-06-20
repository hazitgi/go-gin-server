package apis

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("api/users/")
	userGroup.GET("", listUser)
	userGroup.POST("", createUser)

}

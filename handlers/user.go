package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazitgi/go_gin_server/common"
	// "github.com/hazitgi/go_gin_server/database"
	"github.com/hazitgi/go_gin_server/managers"
	// "github.com/hazitgi/go_gin_server/models"
)

type UserHandler struct {
	groupName   string
	UserManager *managers.UserManager
}

func NewUserHandlerFrom(userManager *managers.UserManager) *UserHandler {
	return &UserHandler{
		"/api/users",
		userManager,
	}
}

func (userHandler *UserHandler) RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group(userHandler.groupName)
	// userGroup.POST("", userHandler.ListUser)
	userGroup.POST("", userHandler.Create)
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {
	userData := common.NewUserCreationInput()

	if err := ctx.BindJSON(&userData); err != nil {
		fmt.Println("failed to bind json: ", err)
	}
	fmt.Println(userData)

	newUser, err := userHandler.UserManager.Create(userData)
	if err != nil {
		fmt.Println("failed to create user: ", err)
	}

	ctx.JSON(http.StatusOK, newUser)
}

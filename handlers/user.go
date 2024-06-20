package handlers

import (
	"fmt"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazitgi/go_gin_server/common"
	// "github.com/hazitgi/go_gin_server/database"
	"github.com/hazitgi/go_gin_server/managers"
	// "github.com/hazitgi/go_gin_server/models"
)

type UserHandler struct {
	groupName   string
	UserManager managers.UserManager
}

func NewUserHandlerFrom(userManager managers.UserManager) *UserHandler {
	return &UserHandler{
		"/api/users",
		userManager,
	}
}

func (userHandler *UserHandler) RegisterUserRoutes(r *gin.Engine) {
	userGroup := r.Group(userHandler.groupName)
	userGroup.GET("", userHandler.ListUsers)
	userGroup.GET(":userId", userHandler.Detail)
	userGroup.DELETE(":userId", userHandler.Delete)
	userGroup.POST("", userHandler.Create)
	userGroup.PUT(":userId", userHandler.Update)
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {
	userData := common.NewUserCreationInput()
	if err := ctx.BindJSON(&userData); err != nil {
		fmt.Println("failed to bind json: ", err)
		common.BadResponse(ctx, "failed to bind json", err.Error())
		return
	}
	fmt.Println(userData)
	newUser, err := userHandler.UserManager.Create(userData)
	if err != nil {
		fmt.Println("failed to create user: ", err)
		common.InternalServerError(ctx, "failed to create user", err.Error())
		return
	}
	common.SuccessResponse(ctx, "user created", newUser)
}

func (userHandler *UserHandler) ListUsers(ctx *gin.Context) {
	users, err := userHandler.UserManager.List()
	if err != nil {
		fmt.Println("failed to list users: ", err)
		common.InternalServerError(ctx, "failed to get users", err.Error())
		return
	}
	common.SuccessResponse(ctx, "success", users)
}

func (userHandler *UserHandler) Detail(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")
	if !ok {
		fmt.Println("invalid userId")
		common.InternalServerError(ctx, "userId is required", nil)
	}
	user, err := userHandler.UserManager.Get(userId)

	if user.ID == 0 {
		common.InternalServerError(ctx, "requested user not found", err.Error())
		return
	}
	if err != nil {
		fmt.Println("failed to list user: ", err)
	}
	common.SuccessResponse(ctx, "success", user)
}

func (userHandler *UserHandler) Delete(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")
	if !ok {
		fmt.Println("invalid userId")
	}
	err := userHandler.UserManager.Delete(userId)
	if err != nil {
		fmt.Println("failed to list users: ", err)
		common.InternalServerError(ctx, "failed to delete user", err.Error())
		return
	}
	common.SuccessResponse(ctx, "deleted successfully", nil)
}

func (userHandler *UserHandler) Update(ctx *gin.Context) {
	userData := common.NewUserUpdateInput()
	if err := ctx.BindJSON(&userData); err != nil {
		fmt.Println("failed to bind json: ", err)
		common.BadResponse(ctx, "failed to bind json", err.Error())
		return
	}
	fmt.Println(userData)
	userId, ok := ctx.Params.Get("userId")
	if !ok {
		fmt.Println("invalid userId")
		common.BadResponse(ctx, "userId is required", nil)
	}
	newUser, err := userHandler.UserManager.Update(userId, userData)
	if err != nil {
		fmt.Println("failed to update user: ", err)
		common.InternalServerError(ctx, "failed to update user", err.Error())
		return
	}
	common.SuccessResponse(ctx, "user updated", newUser)
}

package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCreationInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserUpdateInput struct {
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

func NewUserUpdateInput() *UserUpdateInput {
	return &UserUpdateInput{}
}

type RequestResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}
type ErrorRequestResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(ctx *gin.Context, msg string, data interface{}) {
	response := RequestResponse{
		msg,
		http.StatusOK,
		data,
	}
	ctx.JSON(http.StatusOK, response)
}

func BadResponse(ctx *gin.Context, msg string, err interface{}) {
	response := ErrorRequestResponse{
		msg,
		http.StatusBadRequest,
		err,
	}
	ctx.JSON(http.StatusBadRequest, response)
}

func InternalServerError(ctx *gin.Context, msg string, err interface{}) {
	response := ErrorRequestResponse{
		msg,
		http.StatusInternalServerError,
		err,
	}
	ctx.JSON(http.StatusInternalServerError, response)
}

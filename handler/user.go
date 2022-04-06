package handler

import (
	"api_testing/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.ServiceUser
}

func NewUserHandler(userService user.ServiceUser) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Hello",
	})
}

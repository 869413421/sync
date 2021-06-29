package controllers

import (
	"github.com/gin-gonic/gin"
	"sync/pkg/model/user"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (*UserController) Index(context *gin.Context) {
	users, _ := user.All()
	context.JSON(200, users)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/pkg/auth"
)

type AuthController struct {
	BaseController
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// Login 登录
func (controller *AuthController) Login(context *gin.Context) {
	//1.获取表单数据
	email := context.Request.PostFormValue("email")
	password := context.Request.PostFormValue("password")

	//2.用户认证
	token, errors := auth.Attempt(email, password)
	if len(errors) > 0 {
		context.JSON(http.StatusForbidden, controller.Data(http.StatusForbidden, "认证失败", errors))
		return
	}

	//3.登录成功
	context.JSON(http.StatusForbidden, controller.Data(http.StatusOK, "", token))
}

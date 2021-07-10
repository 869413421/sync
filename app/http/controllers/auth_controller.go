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

type LoginJson struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 登录
func (controller *AuthController) Login(ctx *gin.Context) {
	//1.获取表单数据
	loginJson := LoginJson{}
	if err := ctx.ShouldBind(&loginJson); err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, err.Error(), []string{})
		return
	}

	//2.用户认证
	data, errors := auth.Attempt(loginJson.UserName, loginJson.Password)
	if len(errors) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "认证失败", []string{})
		return
	}

	//3.登录成功
	controller.ResponseJson(ctx, http.StatusOK, "", data)
}

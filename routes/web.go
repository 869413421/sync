package routes

import (
	"github.com/gin-gonic/gin"
	. "sync/app/http/controllers"
	"sync/app/http/middlewares"
)

var userController = NewUserController()
var authController = NewAuthController()

//RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors())
	router.POST("/login", authController.Login)

	// 用户管理理由
	userApi := router.Group("/user").Use(middlewares.Jwt())
	{
		userApi.GET("/index", userController.Index)
	}
}

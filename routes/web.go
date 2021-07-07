package routes

import (
	"github.com/gin-gonic/gin"
	. "sync/app/http/controllers"
	"sync/app/http/middlewares"
)

var userController = NewUserController()
var authController = NewAuthController()
var imageController = NewImageController()

//RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors())
	router.POST("/login", authController.Login)

	// 用户管理路由
	userApi := router.Group("/user").Use(middlewares.Jwt())
	{
		userApi.GET("/index", userController.Index)
		userApi.GET("/:id", userController.Show)
		userApi.PUT("/:id", userController.Update)
	}

	// 图片管理路由
	imgApi := router.Group("/image").Use(middlewares.Jwt())
	{
		imgApi.POST("", imageController.Store)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	. "sync/app/http/controllers"
	"sync/app/http/middlewares"
)

var userController = NewUserController()
var authController = NewAuthController()
var imageController = NewImageController()
var casbinController = NewCasbinController()

var middlewareHandlers []gin.HandlerFunc

//RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	router.Use(middlewares.Cors())
	router.POST("/login", authController.Login)

	middlewareHandlers = append(middlewareHandlers, middlewares.Jwt(), middlewares.Enforcer())
	// 用户管理路由
	userApi := router.Group("/user").Use(middlewareHandlers...)
	{
		userApi.GET("", userController.Index)
		userApi.POST("", userController.Store)
		userApi.GET("/:id", userController.Show)
		userApi.PUT("/:id", userController.Update)
		userApi.DELETE("/:id", userController.Delete)
	}

	// 图片管理路由
	imgApi := router.Group("/image").Use(middlewareHandlers...)
	{
		imgApi.POST("", imageController.Store)
	}

	// casbin管理路由
	casbinApi := router.Group("/casbin").Use(middlewareHandlers...)
	{
		casbinApi.GET("", casbinController.Index)
		casbinApi.GET("/:id", casbinController.Show)
		casbinApi.GET("/tree", casbinController.Tree)
		casbinApi.POST("", casbinController.Store)
		casbinApi.PUT("/:id", casbinController.Update)
		casbinApi.DELETE("/:id", casbinController.Delete)
	}
}

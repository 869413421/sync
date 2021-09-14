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
var permissionController = NewPermissionController()
var roleController = NewRoleController()

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

	// 权限管理路由
	permissionApi := router.Group("/permission").Use(middlewareHandlers...)
	{
		permissionApi.GET("", permissionController.Index)
		permissionApi.GET("/:id", permissionController.Show)
		permissionApi.GET("/tree", permissionController.Tree)
		permissionApi.POST("", permissionController.Store)
		permissionApi.PUT("/:id", permissionController.Update)
		permissionApi.DELETE("/:id", permissionController.Delete)
	}

	// 角色管理路由
	roleApi := router.Group("/role").Use(middlewareHandlers...)
	{
		roleApi.GET("", roleController.Index)
		roleApi.GET("/:id", roleController.Show)
		roleApi.POST("", roleController.Store)
		roleApi.PUT("/:id", roleController.Update)
		roleApi.DELETE("/:id", roleController.Delete)
	}
}

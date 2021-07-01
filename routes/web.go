package routes

import (
	"github.com/gin-gonic/gin"
	. "sync/app/http/controllers"
)

type Router struct {
	Name    string
	Method  string
	Pattern string
	Handle  gin.HandlerFunc
}

var userController = NewUserController()
var authController = NewAuthController()

type WebRouters []Router

var Routes = WebRouters{
	{
		Name:    "index",
		Method:  "get",
		Pattern: "/",
		Handle:  userController.Index,
	},
	{
		Name:    "login",
		Method:  "post",
		Pattern: "/login",
		Handle:  authController.Login,
	},
}

//RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {
	//1.装载所有路由
	for _, route := range Routes {
		switch route.Method {
		case "get":
			router.GET(route.Pattern, route.Handle)
			break
		case "post":
			router.POST(route.Pattern, route.Handle)
			break
		case "delete":
			router.DELETE(route.Pattern, route.Handle)
			break
		case "put":
			router.PUT(route.Pattern, route.Handle)
			break
		case "head":
			router.HEAD(route.Pattern, route.Handle)
			break
		case "options":
			router.OPTIONS(route.Pattern, route.Handle)
			break

		}

		//2.注册全局中间件
		//router.Use(middlewares.StartSession, middlewares.Auth)
	}
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync/pkg/logger"
	"sync/pkg/model/user"
	"sync/pkg/types"
)

type UserController struct {
	BaseController
}

type UserID struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// Index 用户列表
func (*UserController) Index(context *gin.Context) {
	users, _ := user.All()
	context.JSON(200, users)
}

// Show 用户详情
func (controller *UserController) Show(context *gin.Context) {
	//1.获取路由中参数
	id := context.Param("id")
	if id == "" {
		context.JSON(http.StatusBadRequest, controller.Data(http.StatusBadRequest, "route id required", []string{}))
		context.Abort()
		return
	}

	//2.根据ID获取用户信息
	user, err := user.GetByID(types.StringToUInt64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, controller.Data(http.StatusNotFound, "user not found", []string{}))
			context.Abort()
			return
		}
		logger.Danger(err, "user controller get user err")
	}
	//3.返回用户信息
	context.JSON(http.StatusOK, controller.Data(0, "", user))
}

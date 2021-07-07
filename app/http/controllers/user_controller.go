package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync/app/http/requests"
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
func (controller *UserController) Show(ctx *gin.Context) {
	//1.获取路由中参数
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, controller.Data(http.StatusBadRequest, "route id required", []string{}))
		ctx.Abort()
		return
	}

	//2.根据ID获取用户信息
	user, err := user.GetByID(types.StringToUInt64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, controller.Data(http.StatusNotFound, "user not found", []string{}))
			ctx.Abort()
			return
		}
		logger.Danger(err, "user controller get user err")
	}
	//3.返回用户信息
	ctx.JSON(http.StatusOK, controller.Data(0, "", user))
}

func (controller *UserController) Update(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")
	name := ctx.Request.PostFormValue("name")
	email := ctx.Request.PostFormValue("email")
	password := ctx.Request.PostFormValue("password")
	avatar := ctx.Request.PostFormValue("avatar")

	//2.构建用户信息
	_user, err := user.GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, controller.Data(http.StatusNotFound, "user not found", []string{}))
		ctx.Abort()
		return
	}
	_user.Name = name
	_user.Email = email
	_user.Password = password
	_user.Avatar = avatar

	//3.验证提交信息
	errs := requests.ValidateUserEditForm(_user)
	if len(errs) > 0 {
		ctx.JSON(http.StatusForbidden, controller.Data(http.StatusForbidden, "validate error", errs))
		ctx.Abort()
		return
	}

	//4.更新用户
	rowsAffected, err := _user.Update()
	if rowsAffected == 0 {
		ctx.JSON(http.StatusForbidden, controller.Data(http.StatusForbidden, "更新用户失败", err))
		ctx.Abort()
		return
	}

	//5.更新成功，响应信息
	ctx.JSON(http.StatusOK, controller.Data(0, "", _user))
}

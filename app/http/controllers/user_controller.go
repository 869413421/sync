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

//type RequestJson struct {
//	Name     string `json:"Name"`
//	Avatar   string `json:"Avatar"`
//	Email    string `json:"email"`
//	Password string `json:"Password"`
//	Status   int    `json:"Status"`
//}

func NewUserController() *UserController {
	return &UserController{}
}

// Index 用户列表
func (controller *UserController) Index(ctx *gin.Context) {
	//1.获取分页数据
	users, pagerData, err := user.Pagination(ctx, controller.PerPage(ctx))
	if err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, err.Error(), []string{})
		return
	}

	//响应数据
	data := make(map[string]interface{})
	data["PagerData"] = pagerData
	data["users"] = users
	controller.ResponseJson(ctx, http.StatusOK, "", data)
}

// Show 用户详情
func (controller *UserController) Show(ctx *gin.Context) {
	//1.获取路由中参数
	id := ctx.Param("id")
	if id == "" {
		controller.ResponseJson(ctx, http.StatusForbidden, "route id required", []string{})
		return
	}

	//2.根据ID获取用户信息
	user, err := user.GetByID(types.StringToUInt64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			controller.ResponseJson(ctx, http.StatusForbidden, "user not found", []string{})
			return
		}
		logger.Danger(err, "user controller get user err")
	}
	//3.返回用户信息
	controller.ResponseJson(ctx, http.StatusOK, "", user)
}

// Store 新增用户
func (controller *UserController) Store(ctx *gin.Context) {

	//1.构建用户信息
	_user := user.User{}
	ctx.ShouldBind(&_user)

	//2.验证提交信息
	errs := requests.ValidateUserEditForm(_user)
	if len(errs) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "validate error", errs)
		return
	}

	//3.新建用户
	err := _user.Store()
	if err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, "新建用户失败", err)
		return
	}

	//4.更新成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", _user)
}

// Update 更新用户
func (controller *UserController) Update(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")

	//2.构建用户信息
	_user, err := user.GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		controller.ResponseJson(ctx, http.StatusForbidden, "user not found", _user)
		return
	}
	ctx.ShouldBind(&_user)

	//3.验证提交信息
	errs := requests.ValidateUserEditForm(_user)
	if len(errs) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "validate error", errs)
		return
	}

	//4.更新用户
	rowsAffected, err := _user.Update()
	if rowsAffected == 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "更新用户失败", err)
		return
	}

	//5.更新成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", _user)
}

// Delete 删除用户
func (controller *UserController) Delete(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")

	//2.构建用户信息
	_user, err := user.GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		controller.ResponseJson(ctx, http.StatusForbidden, "user not found", err)
		return
	}

	//3.删除用户
	rowsAffected, err := _user.Delete()
	if rowsAffected == 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "删除用户失败", err)
		return
	}

	//5.删除成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", []string{})
}

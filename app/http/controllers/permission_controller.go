package controllers

import "C"
import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync/app/http/requests"
	"sync/pkg/logger"
	. "sync/pkg/model/permission"
	"sync/pkg/types"
	"sync/service/permission_service"
)

type PermissionController struct {
	BaseController
}

func NewPermissionController() *PermissionController {
	return &PermissionController{}
}

// Index 获取列表
func (controller *PermissionController) Index(ctx *gin.Context) {
	//1.构建查询条件
	where := make(map[string]interface{})

	//2.获取分页数据
	rules, pagerData, err := Pagination(ctx, where, controller.PerPage(ctx))
	if err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, err.Error(), []string{})
		return
	}

	//3.响应数据
	data := make(map[string]interface{})
	data["PagerData"] = pagerData
	data["permissions"] = rules
	controller.ResponseJson(ctx, http.StatusOK, "", data)
}

// Show 规则详情
func (controller *PermissionController) Show(ctx *gin.Context) {
	//1.获取路由中参数
	id := ctx.Param("id")
	if id == "" {
		controller.ResponseJson(ctx, http.StatusForbidden, "route id required", []string{})
		return
	}

	//2.根据ID获取规则信息
	rule, err := GetByID(types.StringToUInt64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			controller.ResponseJson(ctx, http.StatusForbidden, "permission not found", []string{})
			return
		}
		logger.Danger(err, "permission controller get permission err")
	}
	//3.返回规则信息
	controller.ResponseJson(ctx, http.StatusOK, "", rule)
}

// Store 新增
func (controller *PermissionController) Store(ctx *gin.Context) {
	//1.获取请求参数
	_permission := Permission{}
	ctx.ShouldBind(&_permission)

	//2.验证提交信息
	errs := requests.ValidatePermission(_permission)
	if len(errs) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "validate error", errs)
		return
	}

	//3.规则
	err := _permission.Store()
	if err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, "新建规则失败", err)
		return
	}

	//4.更新成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", _permission)
}

// Update 更新
func (controller *PermissionController) Update(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")

	//2.构建用户信息
	_permission, err := GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		controller.ResponseJson(ctx, http.StatusForbidden, "permission not found", []string{})
		return
	}
	ctx.ShouldBind(&_permission)

	//3.验证提交信息
	errs := requests.ValidatePermission(_permission)
	if len(errs) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "validate error", errs)
		return
	}

	//4.更新规则
	rowsAffected, err := _permission.Update()
	if rowsAffected == 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "更新规则失败,没有任何更改", err)
		return
	}

	//5.更新成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", _permission)
}

// Delete 删除规则
func (controller *PermissionController) Delete(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")

	//2.构建规则信息
	_permission, err := GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		controller.ResponseJson(ctx, http.StatusForbidden, "user not found", err)
		return
	}

	//3.删除用户
	rowsAffected, err := _permission.Delete()
	if rowsAffected == 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "删除权限失败", err)
		return
	}

	//5.删除成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", []string{})
}

// Tree 获取权限树木
func (controller *PermissionController) Tree(ctx *gin.Context) {
	//1.获取权限树
	tree := permission_service.GetPermissionTree()

	//2.响应数据
	controller.ResponseJson(ctx, http.StatusOK, "", tree)
}

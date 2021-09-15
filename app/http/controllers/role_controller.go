package controllers

import "C"
import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync/app/http/requests"
	"sync/pkg/logger"
	. "sync/pkg/model/role"
	"sync/pkg/types"
	"sync/service/role_service"
)

type RoleController struct {
	BaseController
}

type PermissionsJson struct {
	Permissions []interface{} `json:"permissions"`
}

func NewRoleController() *RoleController {
	return &RoleController{}
}

// Index 获取列表
func (controller *RoleController) Index(ctx *gin.Context) {
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
	data["roles"] = rules
	controller.ResponseJson(ctx, http.StatusOK, "", data)
}

// Show 规则详情
func (controller *RoleController) Show(ctx *gin.Context) {
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
			controller.ResponseJson(ctx, http.StatusForbidden, "Role not found", []string{})
			return
		}
		logger.Danger(err, "Role controller get Role err")
	}
	//3.返回规则信息
	controller.ResponseJson(ctx, http.StatusOK, "", rule)
}

// Store 新增
func (controller *RoleController) Store(ctx *gin.Context) {
	//1.获取请求参数
	_Role := Role{}
	ctx.ShouldBind(&_Role)

	//2.验证提交信息
	errs := requests.ValidateRole(_Role)
	if len(errs) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "validate error", errs)
		return
	}

	//3.规则
	err := _Role.Store()
	if err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, "新建规则失败", err)
		return
	}

	//4.更新成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", _Role)
}

// Update 更新
func (controller *RoleController) Update(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")
	permissions := PermissionsJson{}
	ctx.BindJSON(&permissions)

	//2.构建用户信息
	_Role, err := GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		controller.ResponseJson(ctx, http.StatusForbidden, "Role not found", []string{})
		return
	}
	ctx.ShouldBind(&_Role)

	//3.验证提交信息
	errs := requests.ValidateRole(_Role)
	if len(errs) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "validate error", errs)
		return
	}

	//4.更新规则
	rowsAffected, err := _Role.Update()
	if rowsAffected == 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "更新规则失败,没有任何更改", err)
		return
	}

	//5 更新权限
	role_service.AddPermissionsByRole(_Role.ID, permissions.Permissions)

	//5.更新成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", _Role)
}

// Delete 删除规则
func (controller *RoleController) Delete(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")

	//2.构建规则信息
	_Role, err := GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		controller.ResponseJson(ctx, http.StatusForbidden, "role not found", err)
		return
	}

	//3.删除用户
	rowsAffected, err := _Role.Delete()
	if rowsAffected == 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "删除权限失败", err)
		return
	}

	//5.删除成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", []string{})
}

package controllers

import "C"
import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sync/app/http/requests"
	"sync/pkg/logger"
	"sync/pkg/model/casbin_rule"
	"sync/pkg/types"
	"sync/service/casbin_service"
)

type CasbinController struct {
	BaseController
}

func NewCasbinController() *CasbinController {
	return &CasbinController{}
}

// Index 用户列表
func (controller *CasbinController) Index(ctx *gin.Context) {
	//1.构建查询条件
	ptype := ctx.Request.FormValue("ptype")
	where := make(map[string]interface{})
	where["ptype"] = ptype

	//2.获取分页数据
	rules, pagerData, err := casbin_rule.Pagination(ctx, where, controller.PerPage(ctx))
	if err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, err.Error(), []string{})
		return
	}

	//3.响应数据
	data := make(map[string]interface{})
	data["PagerData"] = pagerData
	data["rules"] = rules
	controller.ResponseJson(ctx, http.StatusOK, "", data)
}

// Show 规则详情
func (controller *CasbinController) Show(ctx *gin.Context) {
	//1.获取路由中参数
	id := ctx.Param("id")
	if id == "" {
		controller.ResponseJson(ctx, http.StatusForbidden, "route id required", []string{})
		return
	}

	//2.根据ID获取规则信息
	rule, err := casbin_rule.GetByID(types.StringToUInt64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			controller.ResponseJson(ctx, http.StatusForbidden, "rule not found", []string{})
			return
		}
		logger.Danger(err, "user controller get user err")
	}
	//3.返回规则信息
	controller.ResponseJson(ctx, http.StatusOK, "", rule)
}

// Store 新增规则
func (controller *CasbinController) Store(ctx *gin.Context) {
	//1.获取请求参数
	_rule := casbin_rule.CasbinRule{}
	ctx.ShouldBind(&_rule)

	//2.验证提交信息
	errs := requests.ValidateRuleEditForm(_rule)
	if len(errs) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "validate error", errs)
		return
	}

	//3.规则
	err := _rule.Store()
	if err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, "新建规则失败", err)
		return
	}

	//4.更新成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", _rule)
}

// Update 更新规则
func (controller *CasbinController) Update(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")

	//2.构建用户信息
	_rule, err := casbin_rule.GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		controller.ResponseJson(ctx, http.StatusForbidden, "rule not found", []string{})
		return
	}
	ctx.ShouldBind(&_rule)

	//3.验证提交信息
	errs := requests.ValidateRuleEditForm(_rule)
	if len(errs) > 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "validate error", errs)
		return
	}

	//4.更新规则
	rowsAffected, err := _rule.Update()
	if rowsAffected == 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "更新规则失败,没有任何更改", err)
		return
	}

	//5.更新成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", _rule)
}

// Delete 删除规则
func (controller *CasbinController) Delete(ctx *gin.Context) {
	//1.获取请求参数
	id := ctx.Param("id")

	//2.构建规则信息
	_rule, err :=  casbin_rule.GetByID(types.StringToUInt64(id))
	if err == gorm.ErrRecordNotFound {
		controller.ResponseJson(ctx, http.StatusForbidden, "user not found", err)
		return
	}

	//3.删除用户
	rowsAffected, err := _rule.Delete()
	if rowsAffected == 0 {
		controller.ResponseJson(ctx, http.StatusForbidden, "删除用户失败", err)
		return
	}

	//5.删除成功，响应信息
	controller.ResponseJson(ctx, http.StatusOK, "", []string{})
}

// Tree 获取权限树木
func (controller *CasbinController) Tree(ctx *gin.Context)  {
	casbin_service.GetPerssionTree()
}

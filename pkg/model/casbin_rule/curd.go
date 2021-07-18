package casbin_rule

import (
	"github.com/gin-gonic/gin"
	"sync/pkg/logger"
	"sync/pkg/model"
	"sync/pkg/pagination"
)

// Pagination 获取所有规则
func Pagination(ctx *gin.Context, where map[string]interface{}, perPage int) (rules []CasbinRule, viewData pagination.ViewData, err error) {
	//1.初始化分页实例
	db := model.DB.Table(CasbinRuleTable)
	for key, val := range where {
		db.Where(key+"=?", val)
	}
	_pager := pagination.New(ctx, db, "/casbin", perPage)

	// 2. 获取视图数据
	viewData = _pager.Paging()

	// 3. 获取数据
	_pager.Results(&rules)

	return rules, viewData, nil
}

// GetByID 根据id获取
func GetByID(id uint64) (rule CasbinRule, err error) {
	err = model.DB.Table(CasbinRuleTable).Where("id=?", id).First(&rule).Error
	return
}

// Store 新增规则
func (rule *CasbinRule) Store() (err error) {
	result := model.DB.Table(CasbinRuleTable).Create(&rule)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model rule create error")
	}
	return
}

// Update 更新规则
func (rule *CasbinRule) Update() (rowsAffected int64, err error) {
	result := model.DB.Table(CasbinRuleTable).Save(&rule)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model rule update error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}

// Delete 删除规则
func (rule *CasbinRule) Delete() (rowsAffected int64, err error) {
	result :=  model.DB.Table(CasbinRuleTable).Delete(&rule)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model rule delete error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}

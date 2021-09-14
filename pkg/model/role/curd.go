package role

import (
	"github.com/gin-gonic/gin"
	"sync/pkg/logger"
	"sync/pkg/model"
	"sync/pkg/pagination"
)

// Pagination 获取分页
func Pagination(ctx *gin.Context, where map[string]interface{}, perPage int) (roles []Role, viewData pagination.ViewData, err error) {
	//1.初始化分页实例
	db := model.DB.Model(&Role{})
	for key, val := range where {
		db.Where(key+"=?", val)
	}
	_pager := pagination.New(ctx, db, "/role", perPage)

	// 2. 获取视图数据
	viewData = _pager.Paging()

	// 3. 获取数据
	err = _pager.Results(&roles)
	if err != nil {
		return nil, pagination.ViewData{}, err
	}

	return roles, viewData, nil
}

// GetByID 根据id获取
func GetByID(id uint64) (role Role, err error) {
	err = model.DB.Model(&Role{}).Where("id=?", id).First(&role).Error
	return
}

// GetByWhere 根据条件获取行数
func GetByWhere(where map[string]interface{}) (role Role, err error) {
	db := model.DB.Model(&Role{})
	for key, val := range where {
		db.Where(key+"=?", val)
	}
	err = db.First(&role).Error
	return
}

// GetList 根据条件获取列表
func GetList(where map[string]interface{}) []Role {
	db := model.DB.Model(&Role{})
	for key, val := range where {
		db.Where(key+"=?", val)
	}
	var role []Role
	db.Scan(&role)
	return role
}

// Store 新增
func (role *Role) Store() (err error) {
	result := model.DB.Model(&Role{}).Create(&role)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model role create error")
	}
	return
}

// Update 更新
func (role *Role) Update() (rowsAffected int64, err error) {
	result := model.DB.Model(&Role{}).Where("id = ?",role.ID).Save(&role)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model role update error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}

// Delete 删除
func (role *Role) Delete() (rowsAffected int64, err error) {
	result := model.DB.Model(&Role{}).Delete(&role)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model role delete error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}

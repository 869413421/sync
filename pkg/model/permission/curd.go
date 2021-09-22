package permission

import (
	"github.com/gin-gonic/gin"
	"sync/pkg/logger"
	"sync/pkg/model"
	"sync/pkg/pagination"
)

// Pagination 获取分页
func Pagination(ctx *gin.Context, where map[string]interface{}, perPage int) (permissions []Permission, viewData pagination.ViewData, err error) {
	//1.初始化分页实例
	db := model.DB.Model(&Permission{})
	for key, val := range where {
		db.Where(key+"=?", val)
	}
	_pager := pagination.New(ctx, db, "/permission", perPage)

	// 2. 获取视图数据
	viewData = _pager.Paging()

	// 3. 获取数据
	err = _pager.Results(&permissions)
	if err != nil {
		return nil, pagination.ViewData{}, err
	}

	return permissions, viewData, nil
}

// GetByID 根据id获取
func GetByID(id uint64) (permission Permission, err error) {
	err = model.DB.Model(&Permission{}).Where("id=?", id).First(&permission).Error
	return
}

// GetByWhere 根据条件获取行数
func GetByWhere(where map[string]interface{}) (permission Permission, err error) {
	db := model.DB.Model(&Permission{})
	for key, val := range where {
		db.Where(key+"=?", val)
	}
	err = db.First(&permission).Error
	return
}

// GetList 根据条件获取列表
func GetList(where map[string]interface{}) []Permission {
	db := model.DB.Model(&Permission{})
	for key, val := range where {
		db.Where(key+"=?", val)
	}
	var permission []Permission
	db.Scan(&permission)
	return permission
}

// Store 新增
func (permission *Permission) Store() (err error) {
	result := model.DB.Model(&Permission{}).Create(&permission)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model permission create error")
	}
	return
}

// Update 更新
func (permission *Permission) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&permission)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model permission update error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}

// Delete 删除
func (permission *Permission) Delete() (rowsAffected int64, err error) {
	result := model.DB.Model(&Permission{}).Delete(&permission)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model permission delete error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}


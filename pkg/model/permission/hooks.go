package permission

import (
	"gorm.io/gorm"
	"strings"
	"sync/pkg/logger"
	"sync/pkg/model"
	"sync/pkg/types"
)

// BeforeSave 保存前模型事件
func (permission *Permission) BeforeSave(*gorm.DB) (err error) {
	//1.判断是否有上级
	if permission.ParentId != 0 {
		//2.获取上级
		parent, getError := GetByID(permission.ParentId)

		if getError != nil {
			logger.Danger(getError, "Permission BeforeSave Error")
			err = getError
			return
		}

		//3.构建ids
		parentIds := parent.ParentIds + "," + types.UInt64ToString(parent.ID)
		parentIds = strings.Trim(parentIds, ",")
		permission.ParentIds = parentIds
	}
	return
}

// AfterDelete 删除后钩子
func (permission *Permission) AfterDelete(tx *gorm.DB) (err error) {
	//1.删除所有下级
	skipHookDB := model.DB.Session(&gorm.Session{
		//设置跳过hook
		SkipHooks: true,
	})
	var permissionList []Permission
	skipHookDB.Where("parent_ids LIKE ?", types.UInt64ToString(permission.ID)+"%").Find(&permissionList)

	//2.删除所有用户拥有该权限数据
	skipHookDB.Table("casbin_rule").Where("v1 = ?", permission.Url).Where("v2 = ?", permission.Method).Where("ptype = ?", "p").Delete(nil)
	for _, item := range permissionList {
		skipHookDB.Table("casbin_rule").Where("v1 = ?", item.Url).Where("v2 = ?", item.Method).Where("ptype = ?", "p").Delete(nil)
	}

	//3.删除下级所有权限
	skipHookDB.Where("parent_ids LIKE ?", types.UInt64ToString(permission.ID)+"%").Delete(&Permission{})

	return
}

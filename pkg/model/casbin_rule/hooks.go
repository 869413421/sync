package casbin_rule

import (
	"gorm.io/gorm"
	"strings"
	"sync/config"
	"sync/pkg/logger"
	"sync/pkg/model"
	"sync/pkg/types"
)

// BeforeSave 保存前模型事件
func (rule *CasbinRule) BeforeSave(*gorm.DB) (err error) {
	//1.判断是否有上级
	if rule.ParentId != 0 {
		//2.获取上级
		parent, getError := GetByID(rule.ParentId)
		if getError != nil {
			logger.Danger(getError, "CasbinRule BeforeSave Error")
			err = getError
			return
		}

		//3.构建ids
		parentIds := parent.ParentIds + "," + types.UInt64ToString(parent.ID)
		parentIds = strings.Trim(parentIds, ",")
		rule.ParentIds = parentIds
	}
	return
}

// AfterDelete 删除后钩子
func (rule *CasbinRule) AfterDelete(*gorm.DB) (err error) {
	if rule.Ptype == "p" {
		//1.删除所有下级
		database := config.LoadConfig().Db.Database
		deleteSql := "DELETE FROM `" + database + "`.`casbin_rule` WHERE parent_ids LIKE '" + types.UInt64ToString(rule.ID) + "%'"
		model.DB.Exec(deleteSql)

		//2.删除所有拥有该权限数据
		deleteSql = "DELETE FROM `" + database + "`.`casbin_rule` WHERE ptype='p' AND v1 = '" + rule.V1 + "'"
		model.DB.Exec(deleteSql)
	}
	return
}

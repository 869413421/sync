package role

import (
	"gorm.io/gorm"
	"sync/pkg/model"
	"sync/pkg/types"
)


// AfterDelete 删除后钩子
func (role *Role) AfterDelete(tx *gorm.DB) (err error) {

	//1.删除所有下级
	skipHookDB := model.DB.Session(&gorm.Session{
		//设置跳过hook
		SkipHooks:true,
	})
	skipHookDB.Table("casbin_rule").Where("v0 = ?",types.UInt64ToString(role.ID)).Where("ptype = ?","p").Delete(nil)
	return
}

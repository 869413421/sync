package user

import (
	"gorm.io/gorm"
	"sync/pkg/enforcer"
	"sync/pkg/password"
)

// BeforeSave 保存前模型事件
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	//1.如果密码没加密，进行一次加密
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}

	//2.删除更新用户角色
	e := enforcer.Enforcer
	_, err = e.DeleteRolesForUser(user.Name)
	if err != nil {
		return err
	}
	_, err = e.AddRoleForUser(user.Name, user.Role)
	if err != nil {
		return err
	}
	return
}

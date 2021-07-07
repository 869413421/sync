package user

import (
	"gorm.io/gorm"
	"sync/pkg/password"
)

// BeforeSave 保存前模型事件
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !password.IsHashed(user.Password) {
		user.Password = password.Hash(user.Password)
	}
	return
}

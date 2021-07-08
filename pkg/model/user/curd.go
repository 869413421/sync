package user

import (
	"sync/pkg/logger"
	"sync/pkg/model"
)

// All 获取所有用户
func All() (users []User, err error) {
	err = model.DB.Find(&users).Error
	return
}

// GetByEmail 根据邮件获取用户
func GetByEmail(email string) (user User, err error) {
	err = model.DB.Where("email=?", email).First(&user).Error
	return
}

// GetByName 根据名称获取用户
func GetByName(name string) (user User, err error) {
	err = model.DB.Where("name=?", name).First(&user).Error
	return
}

// Store 新增用户
func (user *User) Store() (err error) {
	result := model.DB.Create(&user)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model user create error")
	}
	return
}

// Update 更新用户
func (user *User) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&user)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model user update error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}

// Delete 删除用户
func (user *User) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&user)
	err = result.Error
	if err != nil {
		logger.Danger(err, "model user delete error")
		return
	}
	rowsAffected = result.RowsAffected
	return
}

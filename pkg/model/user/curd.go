package user

import (
	"sync/pkg/logger"
	"sync/pkg/model"
)

func All() (users []User, err error) {
	err = model.DB.Find(&users).Error
	return
}

func GetByEmail(email string) (user User, err error) {
	err = model.DB.Where("email=?", email).First(&user).Error
	return
}

func GetByName(name string) (user User, err error) {
	err = model.DB.Where("name=?", name).First(&user).Error
	return
}

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

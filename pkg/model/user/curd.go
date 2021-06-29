package user

import (
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

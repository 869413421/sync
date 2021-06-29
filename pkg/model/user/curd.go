package user

import (
	"sync/pkg/model"
)

func All() (users []User, err error) {
	err = model.DB.Find(&users).Error
	return
}

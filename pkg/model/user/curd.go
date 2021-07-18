package user

import (
	"github.com/gin-gonic/gin"
	"sync/pkg/logger"
	"sync/pkg/model"
	"sync/pkg/pagination"
)

// Pagination 获取所有用户
func Pagination(ctx *gin.Context, perPage int) (users []User, viewData pagination.ViewData, err error) {
	//1.初始化分页实例
	db := model.DB.Model(User{}).Order("created_at desc")
	_pager := pagination.New(ctx, db, "/user", perPage)

	// 2. 获取视图数据
	viewData = _pager.Paging()

	// 3. 获取数据
	_pager.Results(&users)

	return users, viewData, nil
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

// GetByID 根据id获取用户
func GetByID(id uint64) (user User, err error) {
	err = model.DB.Where("id=?", id).First(&user).Error
	return
}

func (user *User) Create() (err error) {
	err = model.DB.Create(&user).Error
	if err != nil {
		logger.Danger(err, "model user create error")
	}
	return
}


package user

import (
	"sync/pkg/model"
	"sync/pkg/password"
)

type User struct {
	model.BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null;unique" valid:"name" json:"name"`
	Email    string `gorm:"column:email;type:varchar(255) not null;unique" valid:"email" json:"email"`
	Password string `gorm:"column:password;type:varchar(255);not null" valid:"password" json:"password"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null;default:''" valid:"avatar" json:"avatar"`
	Status   int    `gorm:"column:status;type:tinyint(1);not null;default:0" json:"status"`
	Role     string `gorm:"-" json:"role"`
	// gorm:"-" 使用这个注解GORM读写会忽略这个字段
	PasswordComfirm string `gorm:"-" valid:"passwordComfirm"`
}

// ComparePassword 比较用户密码是否匹配
func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}

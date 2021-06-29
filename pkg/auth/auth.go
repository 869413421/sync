package auth

import (
	"gorm.io/gorm"
	"sync/pkg/logger"
	"sync/pkg/model/user"
	"sync/pkg/types"
)

type Errors map[string]interface{}

func Attempt(email string, password string) Errors {
	//1.根据邮箱获取用户信息
	user, err := user.GetByEmail(email)

	errors := make(Errors)
	//2.判断用户信息是否出错
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errors["email"] = []string{"账户不存在"}
			return errors
		} else {
			logger.Danger(err, "get user by email error")
		}
	}

	//3.比较用户密码
	if !user.ComparePassword(password) {
		errors["password"] = []string{"密码错误"}
		return errors
	}

	//4.写入session
	Login(user)

	return errors
}

// User 返回认证后的用户
func User() (user.User, error) {
	user, err := user.GetByID(types.StringToUInt64(GetUID()))
	return user, err
}

// Login 登录
func Login(user user.User) {
	session.Put("uid", user.GetStringID())
}

// Logout 退出登录
func Logout(user user.User) {
	session.Forget("uid")
}

// Check 检查用户是否已经登录
func Check() bool {
	return len(GetUID()) > 0
}

// GetUID 从session中获取用户信息
func GetUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if len(uid) > 0 && ok {
		return uid
	}

	return ""
}

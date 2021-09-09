package auth

import (
	"gorm.io/gorm"
	"sync/pkg/jwt"
	"sync/pkg/logger"
	"sync/pkg/message"
	"sync/pkg/model/user"
)

func Attempt(name string, password string) (data map[string]interface{}, errors message.ResponseErrors) {
	defer func() {
		if err := recover(); err != nil {
			logger.Danger(err, "auth Attempt error")
		}
	}()

	//1.根据名称获取用户信息
	user, err := user.GetByName(name)
	errors = make(message.ResponseErrors)
	//2.判断用户信息是否出错
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errors["name"] = []string{"账户不存在"}
			return
		} else {
			logger.Danger(err, "get user by email error")
		}
	}

	//3.比较用户密码
	if !user.ComparePassword(password) {
		errors["password"] = []string{"密码错误"}
		return
	}

	//4.返回token
	token, err := jwt.GenerateToken(user)
	if err != nil {
		errors["name"] = []string{err.Error()}
		return
	}
	data = make(map[string]interface{})
	data["token"] = token
	data["user"] = user

	return
}

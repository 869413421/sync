package auth

import (
	"fmt"
	"gorm.io/gorm"
	"sync/pkg/jwt"
	"sync/pkg/logger"
	"sync/pkg/message"
	"sync/pkg/model/user"
)

func Attempt(email string, password string) (token string, errors message.ResponseErrors) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//1.根据邮箱获取用户信息
	user, err := user.GetByEmail(email)
	errors = make(message.ResponseErrors)
	//2.判断用户信息是否出错
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			errors["email"] = []string{"账户不存在"}
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
	token, err = jwt.GenerateToken(user)
	if err != nil {
		errors["name"] = []string{err.Error()}
		return
	}
	return
}

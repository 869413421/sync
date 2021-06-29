package password

import (
	"golang.org/x/crypto/bcrypt"
	"sync/pkg/logger"
)

// Hash hash加密
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.Danger(err, "hash password error")
	}

	return string(bytes)
}

//CheckHash 检查密码是否与hash值匹配
func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsHashed 检查是否已经加密过
func IsHashed(str string) bool {
	return len(str) == 60
}

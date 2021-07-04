package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	. "sync/config"
	"sync/pkg/model/user"
	"time"
)

var secret = []byte(LoadConfig().Jwt.Secret)

type Claims struct {
	User user.User
	jwt.StandardClaims
}

// GenerateToken 生成jwt
func GenerateToken(user user.User) (token string, err error) {
	fmt.Println(LoadConfig().Jwt.Secret)
	//1.设置过期时间
	nowTime := time.Now()
	expireTime := nowTime.Add(LoadConfig().Jwt.ExpireTime * time.Second)

	//2.构建jwt数据
	claims := Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "sync",
		},
	}

	//3.生成token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenClaims.SignedString(secret)
	return
}

// ParseToken 验证token
func ParseToken(token string) (claims *Claims, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims == nil {
		err = errors.New("ParseToken error")
	}

	fmt.Println(tokenClaims)

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	err = errors.New("claims nil")
	return
}

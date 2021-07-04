package jwt

import (
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
func ParseToken(tokenString string) (user user.User, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		user = claims.User
		return
	} else {
		return
	}

	return
}

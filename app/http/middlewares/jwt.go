package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync/app/http/controllers"
	"sync/pkg/jwt"
)

var base = controllers.NewBaseController()

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//1.获取token
		token := ctx.GetHeader("Authorization")
		if token != "" {
			tokenS := strings.Split(token, " ")
			token = tokenS[1]
		} else {
			token = ctx.Param("token")
		}
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, base.Data(http.StatusUnauthorized, "无法获取token", []string{}))
			ctx.Abort()
			return
		}

		//2.解析token
		_user, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, base.Data(http.StatusUnauthorized, err.Error(), []string{}))
			ctx.Abort()
			return
		}

		//3.通过验证
		ctx.Set("authUser", _user)
		ctx.Next()
	}
}

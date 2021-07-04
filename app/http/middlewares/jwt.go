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
	return func(context *gin.Context) {
		//1.获取token
		token := context.GetHeader("Authorization")
		if token != "" {
			tokenS := strings.Split(token, " ")
			token = tokenS[1]
		} else {
			token = context.Param("token")
		}
		if token == "" {
			context.JSON(http.StatusForbidden, base.Data(http.StatusForbidden, "无法获取token", []string{}))
			context.Abort()
			return
		}

		//2.解析token
		_, err := jwt.ParseToken(token)
		if err != nil {
			context.JSON(http.StatusForbidden, base.Data(http.StatusForbidden, err.Error(), []string{}))
			context.Abort()
			return
		}

		//3.通过验证
		context.Next()
	}
}

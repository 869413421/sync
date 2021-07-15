package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/pkg/enforcer"
	"sync/pkg/logger"
	"sync/pkg/model/user"
)

func Enforcer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//1.检验器
		e := enforcer.Enforcer

		//2.请求url
		obj := ctx.Request.RequestURI

		//3.请求方法
		act := ctx.Request.Method

		//4.获取当前用户角色
		sub := "guest"
		value, authed := ctx.Get("authUser")
		if authed {
			_user := value.(user.User)
			roles, err := e.GetRolesForUser(_user.Name)
			if err != nil {
				logger.Danger(err, "get roles error")
			}
			if len(roles) > 0 {
				sub = roles[0]
			}
		}

		//5.判断是否在策略中存在
		err := e.LoadPolicy()
		if err != nil {
			logger.Danger(err, "加载策略失败")
			base.ResponseJson(ctx, http.StatusForbidden, "加载策略失败", []string{})
			return
		}
		if ok, _ := e.Enforce(sub, obj, act); ok {
			//4，1 放行
			ctx.Next()
			return
		}

		//6.不具备权限
		base.ResponseJson(ctx, http.StatusForbidden, "没有权限访问", []string{})
		return
	}
}

package route

import (
	"github.com/gin-gonic/gin"
)

var route *gin.Engine

// SetRoute 为当前包提供使用对象
func SetRoute(r *gin.Engine) {
	route = r
}

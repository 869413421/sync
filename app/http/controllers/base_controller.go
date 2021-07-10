package controllers

import (
	"github.com/gin-gonic/gin"
	"sync/pkg/message"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (*BaseController) ResponseJson(ctx *gin.Context, code int, errorMsg string, data interface{}) {
	responseData := message.ResponseData{
		Code:     code,
		ErrorMsg: errorMsg,
		Data:     data,
	}

	ctx.JSON(code, responseData)
	ctx.Abort()
}

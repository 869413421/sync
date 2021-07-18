package controllers

import (
	"github.com/gin-gonic/gin"
	"sync/config"
	"sync/pkg/message"
	"sync/pkg/types"
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

func (*BaseController) PerPage(ctx *gin.Context) int {
	page := ctx.Request.FormValue("pageSize")
	if page == "" {
		return config.LoadConfig().Pagination.PerPage
	}
	return types.StringToInt(page)
}

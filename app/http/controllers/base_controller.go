package controllers

import (
	"sync/pkg/message"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (*BaseController) Data(code int, errorMsg string, data interface{}) message.ResponseData {
	responseData := message.ResponseData{
		Code:     code,
		ErrorMsg: errorMsg,
		Data:     data,
	}
	return responseData
}

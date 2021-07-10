package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	config2 "sync/config"
	"sync/pkg/logger"
	"sync/pkg/utils"
	"time"
)

type ImageController struct {
	BaseController
}

func NewImageController() *ImageController {
	return &ImageController{}
}

// Store 图片上传
func (controller *ImageController) Store(ctx *gin.Context) {
	//1.获取文件
	file, err := ctx.FormFile("file")
	if err != nil {
		controller.ResponseJson(ctx, http.StatusForbidden, err.Error(), []string{})
		return
	}

	//2.保存文件
	config := config2.LoadConfig().App
	dateStr := time.Now().Format("2006-01-02")
	dir := config.Static + "/img/" + dateStr + "/"
	exist, _ := utils.PathExist(dir)
	if exist == false {
		os.Mkdir(dir, os.ModeDir)
	}
	fileName := dir + file.Filename
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		logger.Danger(err, "save file error")
		controller.ResponseJson(ctx, http.StatusForbidden, err.Error(), []string{})
		return
	}

	//返回url
	url := config.Protocol + "://" + config.Address + "/static/img/" + dateStr + "/" + file.Filename
	controller.ResponseJson(ctx, http.StatusOK, "", url)
}

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
func (controller *ImageController) Store(context *gin.Context) {
	//1.获取文件
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, controller.Data(http.StatusBadRequest, err.Error(), []string{}))
		context.Abort()
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

	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		logger.Danger(err, "save file error")
		context.JSON(http.StatusBadRequest, controller.Data(http.StatusBadRequest, "文件保存失败", []string{}))
		context.Abort()
		return
	}

	url := config.Protocol + "://" + config.Address + "/static/img/" + dateStr + "/" + file.Filename

	context.JSON(http.StatusCreated, controller.Data(0, "", url))
}

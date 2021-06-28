package logger

import (
	"log"
	"os"
	"sync"
	. "sync/config"
)

var Logger *log.Logger
var once sync.Once

//初始化日志
func init() {
	once.Do(func() {
		//1.初始化日志文件
		config := LoadConfig()

		//2.判断文件夹是否存在，不存在创建
		_, err := os.Stat(config.App.Log)
		if err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(config.App.Log, 666)
			} else {
				log.Fatalln("Failed to make dir ", err)
			}
		}

		//3.创建或打开日志文件
		file, err := os.OpenFile(config.App.Log+"/sync.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 666)
		if err != nil {
			log.Fatalln("Failed to open log file ", err)
		}

		//4.初始化日志对象
		Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

//Info 记录详情
func Info(args ...interface{}) {
	Logger.SetPrefix("INFO ")
	Logger.Println(args...)
}

//Danger 记录错误级别
func Danger(args ...interface{}) {
	Logger.SetPrefix("ERROR ")
	Logger.Println(args...)
}

//Warning 记录警告级别
func Warning(args ...interface{}) {
	Logger.SetPrefix("WARNING ")
	Logger.Println(args...)
}

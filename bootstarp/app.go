package bootstarp

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	. "sync/config"
	"sync/pkg/river"
	"sync/pkg/route"
	"time"
)

func Run() {
	//1.加载配置
	fmt.Println("server starting")
	config := LoadConfig()

	//2.初始化gin路由
	gin.SetMode(config.App.Mode)
	router := SetupRoute()
	route.SetRoute(router)

	//3.设置静态资源访问目录
	router.StaticFS("static", http.Dir(config.App.Static))

	//3.初始化数据库
	SetupDB()

	fmt.Println(river.NewRiver())

	//4.如果是测试，不启动监听
	if config.App.RunTest {
		return
	}

	//5.启动服务端口
	fmt.Println("Server Running")
	server := &http.Server{
		Addr:         config.App.Address,
		Handler:      router,
		ReadTimeout:  config.App.ReadTimeout * time.Second,
		WriteTimeout: config.App.WriteTimeout * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Start Service Error ", err)
		}
	}()

	//6.阻塞信号，平滑关闭
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	fmt.Println("Shutdown Server ....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("shutdown err ", err)
	}

	fmt.Println("Server exit....")

}

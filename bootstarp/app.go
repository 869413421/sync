package bootstarp

import (
	"fmt"
	"net/http"
	. "sync/config"
	"sync/pkg/route"
)

func Run() {
	//1.加载配置
	fmt.Println("server starting")
	config := LoadConfig()

	//2.初始化gin路由
	router := SetupRoute()
	route.SetRoute(router)

	//3.设置静态资源访问目录
	router.StaticFS("static", http.Dir(config.App.Static))

	//3.初始化数据库
	SetupDB()

	//4.启动服务端口
	fmt.Println("Server Running")
	err := http.ListenAndServe(config.App.Address, router)
	if err != nil {
		fmt.Println("Start Service Error ", err)
	}
}

package bootstarp

import (
	"net/http"
	. "sync/config"
	"sync/pkg/route"
)

func Run() {
	//1.加载配置
	config := LoadConfig()

	//2.初始化gin路由
	router := SetupRoute()
	route.SetRoute(router)

	//3.设置静态资源访问目录
	router.StaticFS("static", http.Dir(config.App.Static))

	//3.初始化数据库

}

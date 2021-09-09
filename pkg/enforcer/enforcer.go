package enforcer

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"os"
	"sync/pkg/logger"
	"sync/pkg/model"
)

var Enforcer *casbin.Enforcer

func CreateEnforcer() {
	defer func() {
		err := recover()
		if err != nil {
			logger.Danger(err, "init Enforcer")
		}
	}()
	//1.初始化mysql适配器
	adapter, err := gormadapter.NewAdapterByDB(model.DB)
	if err != nil {
		logger.Danger(err, "create gormadapter error")
	}

	//2.通过适配器创建一个新的enforcer
	dir, _ := os.Getwd()
	modelPath := dir + "/config/rbac_model.conf"
	Enforcer, err = casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		logger.Danger(err, "create Enforcer error")
	}

	//3.开启日志记录
	Enforcer.EnableLog(false)
}

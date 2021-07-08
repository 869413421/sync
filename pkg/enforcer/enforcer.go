package enforcer

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"os"
	"sync/config"
	"sync/pkg/logger"
)

var Enforcer *casbin.Enforcer

func init() {
	//1.初始化mysql适配器
	config := config.LoadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Db.User, config.Db.Password, config.Db.Address, config.Db.Database)
	adapter, err := gormadapter.NewAdapter(config.Db.Driver, dsn, true)
	if err != nil {
		logger.Danger(err, "create gormadapter error")
	}

	//2.通过适配器创建一个新的enforcer
	dir, _ := os.Getwd()
	modelPath := dir + "/config/rbac_model.conf"
	Enforcer, _ = casbin.NewEnforcer(modelPath, adapter)

	//3.开启日志记录
	Enforcer.EnableLog(false)
}

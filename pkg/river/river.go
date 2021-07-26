package river

import (
	"context"
	"github.com/siddontang/go-mysql/canal"
	"sync"
	"sync/config"
	"sync/pkg/model/sync_rule"
)

type River struct {
	wg sync.WaitGroup

	syncCh chan interface{}

	ctx    context.Context
	cancel context.CancelFunc

	master *MasterInfo

	canal *canal.Canal
}

// NewRiver 构建一个河流对象
func NewRiver() (*River, error) {
	//1.初始化管道
	r := new(River)
	r.syncCh = make(chan interface{})
	r.ctx, r.cancel = context.WithCancel(context.Background())

	//2.加载binlog增量游标
	var err error
	if r.master, err = LoadMasterInfo("./var"); err != nil {
		return nil, err
	}

	//3.初始化运河对象
	if err = r.NewCanal(); err != nil {
		return nil, err
	}

	//4.预处理同步规则
	if err = r.prepareRule(); err != nil {
		return nil, err
	}

	return r, err
}

//NewCanal 初始化运河配置
func (r *River) NewCanal() error {
	//1.读取db配置
	dbConfig := config.LoadConfig().Db
	canalConfig := canal.NewDefaultConfig()

	//2.赋值canal配置
	canalConfig.Addr = dbConfig.Address
	canalConfig.User = dbConfig.User
	canalConfig.Password = dbConfig.Password
	canalConfig.Charset = dbConfig.Charset
	canalConfig.Flavor = dbConfig.Driver
	canalConfig.ServerID = dbConfig.ServerID
	canalConfig.Dump.ExecutionPath = dbConfig.DumpExec
	canalConfig.Dump.DiscardErr = dbConfig.DiscardErr
	canalConfig.Dump.SkipMasterData = dbConfig.SkipMasterData

	//3.从数据库中加载规则
	var err error
	rules, err := sync_rule.GetAll()
	if err != nil {
		return err
	}
	for _, rule := range rules {
		canalConfig.IncludeTableRegex = append(canalConfig.ExcludeTableRegex, rule.Schema+"\\."+rule.Table)
	}

	//4.赋值运河对象
	r.canal, err = canal.NewCanal(canalConfig)
	return err
}

func (r *River) prepareRule() error {

}

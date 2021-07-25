package river

import (
	"context"
	"github.com/siddontang/go-mysql/canal"
	"sync"
	"sync/config"
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

	r := new(River)

	r.syncCh = make(chan interface{})

	r.ctx, r.cancel = context.WithCancel(context.Background())

	var err error
	if r.master, err = LoadMasterInfo("./var"); err != nil {
		return nil, err
	}

	return r, err
}

//newCanal 初始化配置
func (r *River) newCanal() error {
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

	//3.从数据库中加载规则 TODO

	//4.赋值运河对象
	var err error
	r.canal, err = canal.NewCanal(canalConfig)
	return err
}

package river

import (
	"context"
	"errors"
	"fmt"
	"github.com/siddontang/go-mysql/canal"
	"regexp"
	"strings"
	"sync"
	"sync/config"
	"sync/pkg/model/sync_rule"
)

type River struct {
	wg sync.WaitGroup

	syncCh chan interface{}

	rules map[string]*Rule

	syncRules []sync_rule.SyncRule

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

	//3.加载数据库规则
	if r.syncRules, err = sync_rule.GetAll(); err != nil {
		return nil, err
	}
	//4.初始化运河对象
	if err = r.NewCanal(); err != nil {
		return nil, err
	}

	//5.预处理同步规则
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
	for _, rule := range r.syncRules {
		canalConfig.IncludeTableRegex = append(canalConfig.ExcludeTableRegex, rule.Schema+"\\."+rule.Table)
	}

	//4.赋值运河对象
	r.canal, err = canal.NewCanal(canalConfig)
	return err
}

func (r *River) prepareRule() error {
	wildTables, err := r.parseSource()
	if err != nil {
		return err
	}
}

// parseSource 初始化MYSQL配置规则源数据为同步规则
func (r *River) parseSource() (map[string][]string, error) {
	//1.校验配置的所有规则
	wildTables := make(map[string][]string, len(r.syncRules))
	if !r.isValidTables() {
		return nil, errors.New("库名表名包含通配符*或者长度为0")
	}

	//2.加载规则
	for _, rule := range r.syncRules {
		//2.1判断规则中是否包含通配符
		if regexp.QuoteMeta(rule.Table) != rule.Table {
			//2.1.1 检验通配符规则是否重复
			if _, ok := wildTables[ruleKey(rule.Schema, rule.Table)]; ok {
				return nil, errors.New(fmt.Sprintf("定义了重复的统配符规则：%s:%s", rule.Schema, rule.Table))
			}
			var tables []string

			//2.1.2 查找所有通配符包含的表
			sql := fmt.Sprintf(`SELECT table_name FROM information_schema.tables WHERE
					table_name RLIKE "%s" AND table_schema = "%s";`, buildTable(rule.Table), rule.Schema)
			res, err := r.canal.Execute(sql)
			if err != nil {
				return nil, err
			}
			for i := 0; i < res.Resultset.RowNumber(); i++ {
				f, _ := res.GetString(i, 0)
				err := r.newRule(rule.Schema, f)
				if err != nil {
					return nil, err
				}
				tables = append(tables, f)
			}

			//2.1.3 将统配的所有表加载到key中
			wildTables[ruleKey(rule.Schema, rule.Table)] = tables
		} else {
			err := r.newRule(rule.Schema, rule.Table)
			if err != nil {
				return nil, err
			}
		}
	}

	if len(r.rules) == 0 {
		return nil, errors.New("没有配置同步规则")
	}

	return wildTables, nil
}

// isValidTables 检查规则是否符合格式
func (r *River) isValidTables() bool {
	for _, rule := range r.syncRules {
		if len(rule.Schema) == 0 {
			return false
		}

		if rule.Schema == "*" {
			return false
		}

		if len(rule.Table) == 0 {
			return false
		}

		if rule.Table == "*" {
			return false
		}
	}

	return true
}

// newRule 创建规则
func (r *River) newRule(schema, table string) error {
	key := ruleKey(schema, table)
	if _, ok := r.rules[key]; ok {
		return errors.New(fmt.Sprintf("生成了重复的规则：%s:%s", schema, table))
	}

	r.rules[key] = newDefaultRule(schema, table)
	return nil
}

// ruleKey 获取规则的key
func ruleKey(schema string, table string) string {
	return strings.ToLower(fmt.Sprintf("%s:%s", schema, table))
}

// buildTable 获取表名称
func buildTable(table string) string {
	if table == "*" {
		return "." + table
	}
	return table
}

package river

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"regexp"
	"strings"
	"sync"
	"sync/config"
	"sync/pkg/logger"
	"sync/pkg/model/sync_rule"
	"sync/pkg/runtime_rule"
)

// ErrRuleNotExist is the error if rule is not defined.
var ErrRuleNotExist = errors.New("rule is not exist")

type River struct {
	wg sync.WaitGroup

	syncCh chan interface{}

	rules map[string]*runtime_rule.Rule

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
	r.rules = make(map[string]*runtime_rule.Rule)
	r.ctx, r.cancel = context.WithCancel(context.Background())

	//2.加载binlog对象
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

	//6.预处理运河对象
	if err = r.prepareCanal(); err != nil {
		return nil, err
	}

	//7.检查mysql binlog必须开启full row image
	if err = r.canal.CheckBinlogRowImage("FULL"); err != nil {
		return nil, err
	}

	return r, err
}

// Run 启动同步
func (r *River) Run() error {
	r.wg.Add(1)
	go r.syncLoop()
	pos := r.master.Position()
	if err := r.canal.RunFrom(pos); err != nil {
		logger.Danger("启动canal同步失败：", err)
		return err
	}

	return nil
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
		canalConfig.IncludeTableRegex = append(canalConfig.IncludeTableRegex, fmt.Sprintf("%s.%s", rule.Schema, rule.Table))
	}

	//4.赋值运河对象
	r.canal, err = canal.NewCanal(canalConfig)
	return err
}

// prepareRule 预处理同步规则
func (r *River) prepareRule() error {
	//1.初始化所有MYSQL配置的规则
	err := r.parseSource()
	if err != nil {
		return err
	}

	//2.启动前预处理规则
	rules := make(map[string]*runtime_rule.Rule)
	for key, rule := range r.rules {
		//2.1获取表详情
		if rule.TableInfo, err = r.canal.GetTable(rule.Schema, rule.Table); err != nil {
			return err
		}

		//2.2 检查表中是否有主键
		if len(rule.TableInfo.PKColumns) == 0 {
			return errors.New(fmt.Sprintf("%s.%s没有主键", rule.Schema, rule.Table))
		}

		rules[key] = rule
	}

	r.rules = rules
	return nil
}

// parseSource 初始化MYSQL配置规则源数据为同步规则
func (r *River) parseSource() error {
	//1.校验配置的所有规则
	if !r.isValidTables() {
		return errors.New("库名表名包含通配符*或者长度为0")
	}

	//2.加载规则
	for _, rule := range r.syncRules {
		//2.1判断规则中是否包含通配符
		if regexp.QuoteMeta(rule.Table) != rule.Table {
			//2.1.1 检验通配符规则是否重复
			if _, ok := r.rules[ruleKey(rule.Schema, rule.Table)]; ok {
				return errors.New(fmt.Sprintf("定义了重复的统配符规则：%s:%s", rule.Schema, rule.Table))
			}

			//2.1.2 查找所有通配符包含的表
			sql := fmt.Sprintf(`SELECT table_name FROM information_schema.tables WHERE
					table_name RLIKE "%s" AND table_schema = "%s";`, buildTable(rule.Table), rule.Schema)
			res, err := r.canal.Execute(sql)
			if err != nil {
				return err
			}
			for i := 0; i < res.Resultset.RowNumber(); i++ {
				f, _ := res.GetString(i, 0)
				err := r.newRule(rule.Schema, f)
				if err != nil {
					return err
				}
			}
		} else {
			err := r.newRule(rule.Schema, rule.Table)
			if err != nil {
				return err
			}
		}
	}

	if len(r.rules) == 0 {
		return errors.New("没有配置同步规则")
	}

	return nil
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

	r.rules[key] = runtime_rule.NewDefaultRule(schema, table)
	return nil
}

// prepareCanal 预处理运河对象
func (r *River) prepareCanal() error {
	//1.构建备份表信息
	var db string
	dbs := map[string]struct{}{}
	tables := make([]string, 0, len(r.rules))
	for _, rule := range r.rules {
		db = rule.Schema
		dbs[rule.Schema] = struct{}{}
		tables = append(tables, rule.Table)
	}

	//2.启动mysql dump
	if len(dbs) == 1 {
		r.canal.AddDumpTables(db, tables...)
	} else {
		keys := make([]string, 0, len(dbs))
		for key := range dbs {
			keys = append(keys, key)
		}

		r.canal.AddDumpDatabases(keys...)
	}

	//3.设置事件处理对象
	r.canal.SetEventHandler(&eventHandler{r: r})

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

func (r *River) updateRule(schema, table string) error {
	rule, ok := r.rules[ruleKey(schema, table)]
	if !ok {
		return ErrRuleNotExist
	}

	tableInfo, err := r.canal.GetTable(schema, table)
	if err != nil {
		return err
	}

	rule.TableInfo = tableInfo

	return nil
}

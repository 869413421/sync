package river

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"sync/pkg/elastic"
	"sync/pkg/logger"
	"time"
)

// binlog 游标对象
type posSaver struct {
	pos   mysql.Position
	force bool
}

//eventHandler  事件处理对象
type eventHandler struct {
	r *River
}

//OnRotate 当产生新的binlog日志后触发(在达到内存的使用限制后（默认为 1GB），会开启另一个文件，每个新文件的名称后都会有一个增量。)
func (h *eventHandler) OnRotate(e *replication.RotateEvent) error {
	pos := mysql.Position{
		Name: string(e.NextLogName),
		Pos:  uint32(e.Position),
	}

	h.r.syncCh <- posSaver{pos, true}

	return h.r.ctx.Err()
}

//OnTableChanged 创建、更改、重命名或删除表时触发，通常会需要清除与表相关的数据，如缓存。It will be called before OnDDL.
func (h *eventHandler) OnTableChanged(schema, table string) error {
	err := h.r.updateRule(schema, table)
	if err != nil && err != ErrRuleNotExist {
		return err
	}
	return nil
}

//OnDDL create alter drop truncate(删除当前表再新建一个一模一样的表结构)
func (h *eventHandler) OnDDL(nextPos mysql.Position, _ *replication.QueryEvent) error {
	h.r.syncCh <- posSaver{nextPos, true}
	return h.r.ctx.Err()
}

func (h *eventHandler) OnXID(nextPos mysql.Position) error {
	h.r.syncCh <- posSaver{nextPos, false}
	return h.r.ctx.Err()
}

//OnRow 监听数据记录
func (h *eventHandler) OnRow(e *canal.RowsEvent) error {
	//rule, ok := h.r.rules[ruleKey(e.Table.Schema, e.Table.Name)]
	//if !ok {
	//	return nil
	//}
	fmt.Println("row run")
	fmt.Println(time.Now())
	fmt.Println(e.Action)
	fmt.Println(e.Table.Name)
	fmt.Println(e.Table.Schema)
	fmt.Println(e.String())


	h.r.syncCh<-posSaver{mysql.Position{Name: "mysql-bin.000001",Pos: 6365910}, false}
	elastic.Create()
	return nil
}

func (h *eventHandler) OnGTID(gtid mysql.GTIDSet) error {
	return nil
}

//OnPosSynced 监听binlog日志的变化文件与记录的位置
func (h *eventHandler) OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error {
	return nil
}

func (h *eventHandler) String() string {
	return "ESRiverEventHandler"
}

func (r *River) syncLoop() {
	bulkSize := 100
	if bulkSize == 0 {
		bulkSize = 128
	}

	var interval time.Duration
	interval = 0
	if interval == 0 {
		interval = 200 * time.Millisecond
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer r.wg.Done()

	lastSavedTime := time.Now()
	//reqs := make([]*elastic.BulkRequest, 0, 1024)

	var pos mysql.Position

	for {
		needFlush := false
		needSavePos := false

		select {
		case v := <-r.syncCh:
			fmt.Println(v)
			switch v := v.(type) {
			case posSaver:
				now := time.Now()
				if v.force || now.Sub(lastSavedTime) > 3*time.Second {
					lastSavedTime = now
					needFlush = true
					needSavePos = true
					pos = v.pos
				}
				//case []*elastic.BulkRequest:
				//	reqs = append(reqs, v...)
				//	needFlush = len(reqs) >= bulkSize
			}
		case <-ticker.C:
			needFlush = true
		case <-r.ctx.Done():
			return
		}

		if needFlush {
			// TODO: retry some times?
			//if err := r.doBulk(reqs); err != nil {
			//	log.Errorf("do ES bulk err %v, close sync", err)
			//	r.cancel()
			//	return
			//}
			//reqs = reqs[0:0]
		}

		if needSavePos {
			if err := r.master.Save(pos); err != nil {
				logger.Danger("save sync position %s err %v, close sync", pos, err)
				r.cancel()
				return
			}
		}
	}
}

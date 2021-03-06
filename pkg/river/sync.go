package river

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/olivere/elastic"
	"sync/pkg/elasticsearch_client"
	"sync/pkg/logger"
	"sync/service/elasitc_service"
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
	//1.查找配置规则
	rule, ok := h.r.rules[ruleKey(e.Table.Schema, e.Table.Name)]
	if !ok {
		return nil
	}

	//2.根据动作分发数据
	var err error
	esService := elasitc_service.NewElasticService(rule, e.Rows)
	switch e.Action {
	case canal.InsertAction:
		reqs, _ := esService.MakeInsertRequest()
		h.r.syncCh <- reqs
	case canal.UpdateAction:
		reqs, _ := esService.MakeUpdateRequest()
		h.r.syncCh <- reqs
	case canal.DeleteAction:
		reqs, _ := esService.MakeDeleteRequest()
		h.r.syncCh <- reqs
	}

	if err != nil {
		fmt.Println(err)
	}

	return h.r.ctx.Err()
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
	bulkSize := 1

	var interval time.Duration
	interval = 0
	if interval == 0 {
		interval = 200 * time.Millisecond
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer r.wg.Done()

	lastSavedTime := time.Now()
	reqs := make([]elastic.BulkableRequest, 0, 1024)

	var pos mysql.Position

	for {
		needFlush := false
		needSavePos := false

		select {
		case v := <-r.syncCh:
			switch v := v.(type) {
			case posSaver:
				now := time.Now()
				if v.force || now.Sub(lastSavedTime) > 3*time.Second {
					lastSavedTime = now
					needFlush = true
					needSavePos = true
					pos = v.pos
				}
			case []elastic.BulkableRequest:
				reqs = v
				needFlush = len(reqs) >= bulkSize
			}
		case <-ticker.C:
			needFlush = true
		case <-r.ctx.Done():
			return
		}

		if needFlush {
			//TODO: retry some times?
			if _, err := elasticsearch_client.Bulk(reqs); err != nil {
				//r.cancel()
				logger.Danger(err)
				return
			}
			reqs = reqs[0:0]
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

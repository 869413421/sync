package elasitc_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/schema"
	"github.com/olivere/elastic"
	"strings"
	"sync/pkg/logger"
	"sync/pkg/runtime_rule"
	"time"
)

// ElasticService Es处理业务逻辑
type ElasticService struct {
	Rule *runtime_rule.Rule
	Rows [][]interface{}
}

// NewElasticService 初始化业务逻辑对象
func NewElasticService(rule *runtime_rule.Rule, rows [][]interface{}) *ElasticService {
	return &ElasticService{Rule: rule, Rows: rows}
}

// MakeInsertRequest 创建Es新增请求
func (service *ElasticService) MakeInsertRequest() ([]elastic.BulkableRequest, error) {
	return service.MakeRequest(canal.InsertAction)
}

// MakeUpdateRequest 创建Es更新请求
func (service *ElasticService) MakeUpdateRequest() ([]elastic.BulkableRequest, error) {
	return service.MakeRequest(canal.UpdateAction)
}

// MakeDeleteRequest 创建Es删除请求
func (service *ElasticService) MakeDeleteRequest() ([]elastic.BulkableRequest, error) {
	return service.MakeRequest(canal.DeleteAction)
}

// MakeRequest 创建Es请求
func (service *ElasticService) MakeRequest(action string) ([]elastic.BulkableRequest, error) {
	//1.创建请求数组
	reqs := make([]elastic.BulkableRequest, 0, 20)
	rule := service.Rule

	//2.从监听的binlog中构建数据
	for _, values := range service.Rows {
		//2.1获取数据中的主键
		id, err := service.getDocId(values)
		if err != nil {
			return nil, err
		}

		//2.2获取请求的mapping数据,转换为json
		mappingData := service.makeInsertReqData(values)
		body, err := json.Marshal(mappingData)
		if err != nil {
			logger.Danger(fmt.Sprintf("make es request json error:%s,data:%v", err, values))
			continue
		}
		requestJson := string(body)

		//2.3 根据触发类型创建请求
		switch action {
		case canal.InsertAction:
			req := elastic.NewBulkIndexRequest().Index(rule.Index).Type(rule.Type).Id(id).Doc(requestJson)
			reqs = append(reqs, req)
		case canal.UpdateAction:
			req := elastic.NewBulkUpdateRequest().Index(rule.Index).Type(rule.Type).Id(id).Doc(requestJson)
			reqs = append(reqs, req)
		case canal.DeleteAction:
			req := elastic.NewBulkDeleteRequest().Index(rule.Index).Type(rule.Type).Id(id)
			reqs = append(reqs, req)
		}
	}

	//3.返回请求数组
	return reqs, nil
}

// getDocId 获取binlog中的id
func (service *ElasticService) getDocId(row []interface{}) (string, error) {
	var (
		ids []interface{}
		err error
	)
	rule := service.Rule

	if rule.ID == nil {
		ids, err = rule.TableInfo.GetPKValues(row)
		if err != nil {
			return "", err
		}
	} else {
		ids = make([]interface{}, 0, len(rule.ID))
		for _, column := range rule.ID {
			value, err := rule.TableInfo.GetColumnValue(column, row)
			if err != nil {
				return "", err
			}
			ids = append(ids, value)
		}
	}

	var buf bytes.Buffer
	seq := ""
	for _, value := range ids {
		buf.WriteString(fmt.Sprintf("%s%v", seq, value))
		seq = ":"
	}

	return buf.String(), nil
}

// makeInsertReqData 构建请求数据
func (service *ElasticService) makeInsertReqData(values []interface{}) map[string]interface{} {
	data := make(map[string]interface{}, len(values))
	rule := service.Rule

	for i, c := range rule.TableInfo.Columns {
		//if !rule.CheckFilter(c.Name) {
		//	continue
		//}
		//mapped := false
		//for k, v := range rule.FieldMapping {
		//	mysql, elastic, fieldType := r.getFieldParts(k, v)
		//	if mysql == c.Name {
		//		mapped = true
		//		req.Data[elastic] = r.getFieldValue(&c, fieldType, values[i])
		//	}
		//}
		//if mapped == false {
		//	req.Data[c.Name] = r.makeReqColumnData(&c, values[i])
		//}

		data[c.Name] = service.MakeReqColumnData(&c, values[i])
	}

	return data
}

//MakeReqColumnData 获取列数据
func (service *ElasticService) MakeReqColumnData(col *schema.TableColumn, value interface{}) interface{} {
	switch col.Type {
	case schema.TYPE_ENUM:
		switch value := value.(type) {
		case int64:
			// for binlog, ENUM may be int64, but for dump, enum is string
			eNum := value - 1
			if eNum < 0 || eNum >= int64(len(col.EnumValues)) {
				// we insert invalid enum value before, so return empty
				logger.Danger(fmt.Sprintf("invalid binlog enum index %d, for enum %v", eNum, col.EnumValues))
				return ""
			}

			return col.EnumValues[eNum]
		}
	case schema.TYPE_SET:
		switch value := value.(type) {
		case int64:
			// for binlog, SET may be int64, but for dump, SET is string
			bitmask := value
			sets := make([]string, 0, len(col.SetValues))
			for i, s := range col.SetValues {
				if bitmask&int64(1<<uint(i)) > 0 {
					sets = append(sets, s)
				}
			}
			return strings.Join(sets, ",")
		}
	case schema.TYPE_BIT:
		switch value := value.(type) {
		case string:
			// for binlog, BIT is int64, but for dump, BIT is string
			// for dump 0x01 is for 1, \0 is for 0
			if value == "\x01" {
				return int64(1)
			}

			return int64(0)
		}
	case schema.TYPE_STRING:
		switch value := value.(type) {
		case []byte:
			return string(value[:])
		}
	case schema.TYPE_JSON:
		var f interface{}
		var err error
		switch v := value.(type) {
		case string:
			err = json.Unmarshal([]byte(v), &f)
		case []byte:
			err = json.Unmarshal(v, &f)
		}
		if err == nil && f != nil {
			return f
		}
	case schema.TYPE_DATETIME, schema.TYPE_TIMESTAMP:
		switch v := value.(type) {
		case string:
			vt, err := time.ParseInLocation(mysql.TimeFormat, string(v), time.Local)
			if err != nil || vt.IsZero() { // failed to parse date or zero date
				return nil
			}
			return vt.Format(time.RFC3339)
		}
	case schema.TYPE_DATE:
		switch v := value.(type) {
		case string:
			vt, err := time.Parse(runtime_rule.MysqlDateFormat, string(v))
			if err != nil || vt.IsZero() { // failed to parse date or zero date
				return nil
			}
			return vt.Format(runtime_rule.MysqlDateFormat)
		}
	}

	return value
}

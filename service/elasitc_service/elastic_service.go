package elasitc_service

import (
	"bytes"
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/olivere/elastic"
	elastic2 "sync/pkg/elastic"
	"sync/pkg/river"
)

// MakeInsertRequest 创建Es新增请求
func MakeInsertRequest(rule *river.Rule, rows [][]interface{}) (*elastic.BulkService, error) {
	return MakeRequest(rule, canal.InsertAction, rows)
}

// MakeRequest 创建Es请求
func MakeRequest(rule *river.Rule, action string, rows [][]interface{}) (*elastic.BulkService, error) {
	reqs := elastic2.EsClient.Bulk()
	for _, values := range rows {
		id, err := getDocId(rule, values)
		if err != nil {
			return nil, err
		}

		req := elastic.NewBulkIndexRequest().Index(rule.Index).Type(rule.Type).Id(id).Doc(values)
		reqs.Add(req)
	}

	return reqs, nil
}

// getDocId 获取binlog中的id
func getDocId(rule *river.Rule, row []interface{}) (string, error) {
	var (
		ids []interface{}
		err error
	)

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

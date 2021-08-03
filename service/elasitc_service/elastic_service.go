package elasitc_service

import (
	"bytes"
	"fmt"
	"sync/pkg/river"
)

func InsertSync(rule *river.Rule, rows [][]interface{}) {

}

func MakeRequest(rule *river.Rule, action string, rows [][]interface{}) {
	for _, values := range rows {
		id, err := getDocId(rule, values)
	}
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

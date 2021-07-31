package river

import (
	"github.com/go-mysql-org/go-mysql/schema"
	"strings"
)

type Rule struct {
	Schema string
	Table  string
	Index  string
	Type   string
	Parent string
	ID     []string

	FieldMapping map[string]string

	TableInfo *schema.Table

	Filter []string

	Pipeline string
}

// newDefaultRule 返回默认规则对象
func newDefaultRule(schema, table string) *Rule {
	rule := new(Rule)

	rule.Schema = schema
	rule.Table = table

	lowerTable := strings.ToLower(table)
	rule.Index = lowerTable
	rule.Type = lowerTable
	rule.FieldMapping = make(map[string]string)

	return rule
}

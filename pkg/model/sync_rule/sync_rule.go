package sync_rule

import (
	"sync/pkg/model"
	"sync/pkg/types"
)

// SyncRule 同步规则
type SyncRule struct {
	model.BaseModel
	Schema  string `gorm:"column:schema;type:varchar(50);not null;default:''" valid:"schema"`
	Table   string `gorm:"column:table;type:varchar(50);not null;default:''" valid:"table"`
	EsIndex string `gorm:"column:es_index;type:varchar(50);not null;default:''" valid:"es_index"`
	EsType  string `gorm:"column:es_type;type:varchar(50);not null;default:''" valid:"es_type"`
	EsKey   string `gorm:"column:es_key;type:varchar(50);not null;default:''" valid:"es_key"`
	Mapping string `gorm:"column:mapping;type:JSON" valid:"mapping" json:"mapping"`
}

// GetMapping 获取es映射
func (rule *SyncRule) GetMapping() map[string]interface{} {
	return types.JsonToMap(rule.Mapping)
}

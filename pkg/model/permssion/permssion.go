package permssion

import "sync/pkg/model"

type Permssion struct {
	model.BaseModel
	Url      string `gorm:"column:url;type:varchar(255);not null;default:'';uniqueIndex:uniq_name,priority:1" valid:"url" json:"url"`
	Name     string `gorm:"column:name;type:varchar(50);not null;uniqueIndex:uniq_name,priority:2" valid:"name" json:"name"`
	Method   string `gorm:"column:method;type:varchar(50);not null;default:'';uniqueIndex:uniq_name,priority:3" valid:"method" json:"method"`
	Desc     string `gorm:"column:desc;type:varchar(255);not null;default:''" valid:"desc" json:"desc"`
	Order    int    `gorm:"column:order;type:int(11);not null;default:0" valid:"order" json:"order"`
	ParentId uint64 `gorm:"column:parent_id;type:int(11);not null;default:0" valid:"parent_id" json:"parent_id"`
	ParentIds string `gorm:"column:parent_ids;type:varchar(500);not null;default:''" valid:"parent_ids" json:"parent_ids"`
}

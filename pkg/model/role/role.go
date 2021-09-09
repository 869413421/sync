package role

import "sync/pkg/model"

type Role struct {
	model.BaseModel
	Name  string `gorm:"column:name;type:varchar(50);not null;uniqueIndex:uniq_name" valid:"name" json:"name"`
	Desc  string `gorm:"column:desc;type:varchar(255);not null;default:''" valid:"desc" json:"desc"`
	Order int    `gorm:"column:order;type:int(11);not null;default:0" json:"order"`
}

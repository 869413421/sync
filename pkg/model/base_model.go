package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gloger "gorm.io/gorm/logger"
	. "sync/config"
	"sync/pkg/logger"
	"sync/pkg/types"
	"time"
)

type BaseModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

func (model BaseModel) GetStringID() string {
	return types.UInt64ToString(model.ID)
}

func (model BaseModel) CreatedAtDate() string {
	return model.CreatedAt.Format("2006-01-02")
}

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	//1.读取配置
	var err error
	config := LoadConfig()

	//2.构建dns
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Db.User, config.Db.Password, config.Db.Address, config.Db.Database)

	//3.连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gloger.Default.LogMode(gloger.Info),
	})
	if err != nil {
		logger.Danger(err, "gorm open error")
	}

	//4.返回db对象
	return DB
}

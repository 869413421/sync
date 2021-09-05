package bootstarp

import (
	"gorm.io/gorm"
	"sync/config"
	"sync/pkg/model"
	"sync/pkg/model/sync_rule"
	"sync/pkg/model/user"
)

// SetupDB 初始化gorm
func SetupDB() {
	//1.建立连接池
	dbConfig := config.LoadConfig()
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	//2.设置最大连接数
	sqlDB.SetMaxOpenConns(dbConfig.Db.MaxConnections)

	//3.设置最大空闲连接数
	sqlDB.SetMaxIdleConns(dbConfig.Db.MaxIdeConnections)

	//4.设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(dbConfig.Db.ConnectionMaxLifeTime)

	//5.执行数据迁移
	migration(db)
}

// migration 数据迁移
func migration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "Charset=utf8")
	db.AutoMigrate(&user.User{}, &sync_rule.SyncRule{})
	db.Exec("ALTER TABLE `sync`.`casbin_rule` ADD COLUMN `name` VARCHAR(255) NOT NULL DEFAULT '' AFTER `v5`;")
	db.Exec("ALTER TABLE `sync`.`casbin_rule` ADD COLUMN `desc` VARCHAR(255) NOT NULL DEFAULT '' AFTER `name`;")
}

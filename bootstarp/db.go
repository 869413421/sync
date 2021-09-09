package bootstarp

import (
	"fmt"
	"gorm.io/gorm"
	"sync/config"
	"sync/pkg/enforcer"
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

	//6.初始化enforcer链接
	enforcer.CreateEnforcer()
}

// migration 数据迁移
func migration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "Charset=utf8")
	db.AutoMigrate(&user.User{}, &sync_rule.SyncRule{})

	database := config.LoadConfig().Db.Database
	db.Exec(fmt.Sprintf("ALTER TABLE `%s`.`casbin_rule` ADD COLUMN `name` VARCHAR(255) NOT NULL DEFAULT '' AFTER `v5`;", database))
	db.Exec(fmt.Sprintf("ALTER TABLE `%s`.`casbin_rule` ADD COLUMN `desc` VARCHAR(255) NOT NULL DEFAULT '' AFTER `name`;", database))
	db.Exec(fmt.Sprintf("ALTER TABLE `%s`.`casbin_rule` ADD COLUMN `parent_id` INT(11) NOT NULL DEFAULT 0 AFTER `desc`;", database))
	db.Exec(fmt.Sprintf("ALTER TABLE `%s`.`casbin_rule` ADD COLUMN `parent_ids` VARCHAR(500) NOT NULL DEFAULT '' AFTER `parent_id`;", database))
}

package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"time"
)

func InitGorm() {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql，取消gorm连接")

	}
	dsn := global.Config.Mysql.Dsn()
	fmt.Println(dsn)
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		//开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		//只打印错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		global.Log.Fatal(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	global.DB = db
}

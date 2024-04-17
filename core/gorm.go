package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"log"
	"os"
	"time"
)

// 自定义日志记录器
var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // 打印日志到标准输出
	logger.Config{
		SlowThreshold: time.Second, // 慢查询阈值，设置为零表示禁用
		LogLevel:      logger.Info, // 日志级别
		Colorful:      true,        // 着色打印日志
	},
)

func InitGorm() {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql，取消gorm连接")

	}
	dsn := global.Config.Mysql.Dsn()
	fmt.Println(dsn)

	global.MysqlLog = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ // 启用详细日志记录,无论何时gorm都打印相应sql
		Logger: newLogger,
		//Logger:OtherLog(),
	})
	if err != nil {
		global.Log.Fatal(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Hour * 5)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	global.DB = db

}

func OtherLog() logger.Interface {
	var mysqlLogger logger.Interface
	mysqlLogger = logger.Default.LogMode(logger.Info)
	if global.Config.System.Env == "dev" {
		//开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		//只打印错误的sql
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	return mysqlLogger
}

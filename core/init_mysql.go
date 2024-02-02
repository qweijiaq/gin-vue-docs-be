package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvd_server/global"
	"time"
)

func InitMysql() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warn("未配置 MySQL，取消 Gorm 连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface

	var logLevel = logger.Error
	switch global.Config.Mysql.LogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	}

	mysqlLogger = logger.Default.LogMode(logLevel)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   mysqlLogger,
		DisableForeignKeyConstraintWhenMigrating: true, // 数据迁移时不生成外键约束
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] MySQL 连接失败, error: %s", dsn, err.Error()))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间, 不能超过 MySQL 的 wait_timeout
	return db
}

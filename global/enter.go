package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvd_server/config"
)

// 使用指针是保证全局唯一性
var (
	Config *config.Config // 配置
	Log    *logrus.Logger // 日志
	DB     *gorm.DB       // MySQL 数据库
	Redis  *redis.Client  // Redis 数据库
)

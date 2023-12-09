package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvd_server/config"
)

var (
	Config *config.Config // 使用指针是保证全局唯一性
	Log    *logrus.Logger
	DB     *gorm.DB
	Redis  *redis.Client
)

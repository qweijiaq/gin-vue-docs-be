package global

import (
	"github.com/sirupsen/logrus"
	"gvd_server/config"
)

var (
	Config *config.Config // 使用指针是保证全局唯一性
	Log    *logrus.Logger
)

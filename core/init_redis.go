package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"time"
)

func InitRedis(db int) *redis.Client {

	redisConf := global.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := client.Ping().Result()
	if err != nil {
		logrus.Fatalf("%s Redis 连接失败 error: %s", redisConf.Addr(), err.Error())
	}
	return client
}

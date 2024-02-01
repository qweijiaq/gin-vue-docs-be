package config

import "fmt"

type Redis struct {
	IP       string `yaml:"ip"`       // IP
	Port     int    `yaml:"port"`     // 端口
	Password string `yaml:"password"` // 密码
	PoolSize int    `yaml:"poolSize"` // 连接池大小
}

// Addr 返回 Redis 数据库的地址
func (redis Redis) Addr() string {
	return fmt.Sprintf("%s:%d", redis.IP, redis.Port)
}

package config

import "fmt"

type Redis struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	PoolSize int    `yaml:"poolSize"` // 连接池大小
}

func (redis Redis) Addr() string {
	return fmt.Sprintf("%s:%d", redis.IP, redis.Port)
}

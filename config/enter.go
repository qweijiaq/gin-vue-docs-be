package config

type Config struct {
	System System `yaml:"system"` // 系统配置
	Mysql  Mysql  `yaml:"mysql"`  // MySQL 数据库配置
	Redis  Redis  `yaml:"redis"`  // Redis 缓存配置
	Jwt    Jwt    `yaml:"jwt"`    // JWT 配置（登录验证）
}

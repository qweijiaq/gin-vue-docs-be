package config

import "fmt"

type System struct {
	IP   string `yaml:"ip"`   // IP
	Port int    `yaml:"port"` // 端口
	Env  string `yaml:"env"`  // 环境
}

// Addr 返回项目后端的访问地址（IP + 端口）
func (system System) Addr() string {
	return fmt.Sprintf("%s:%d", system.IP, system.Port)
}

package config

type Jwt struct {
	Expires int    `yaml:"expires"` // 过期时间 单位: 天
	Issuer  string `yaml:"issuer"`  // 颁发人
	Secret  string `yaml:"secret"`  // 秘钥
}

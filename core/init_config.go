package core

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"gvd_server/config"
	"os"
)

const yamlPath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath) // 读取 YAML 文件
	if err != nil {
		logrus.Fatalln("read yaml error: ", err.Error())
	}
	c = new(config.Config)            // 如果不 new 的话，config.Config 为 nil, 可能引起报错 -- nil 指针错误是 Go 中常见的错误
	err = yaml.Unmarshal(byteData, c) // 反序列化
	if err != nil {
		logrus.Fatalln("parse yaml error: ", err.Error())
	}
	return c
}

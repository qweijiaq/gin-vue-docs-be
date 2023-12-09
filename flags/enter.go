package flags

import "flag"

type Option struct {
	DB   bool   // 初始化数据库
	Port int    // 设置端口号
	Load string // 导入数据集
}

func Parse() (option *Option) {
	option = new(Option)
	flag.BoolVar(&option.DB, "db", false, "初始化数据库")
	flag.IntVar(&option.Port, "port", 0, "程序运行的端口")
	flag.StringVar(&option.Load, "load", "", "导入 sql 数据集")
	flag.Parse()

	return option
}

// Run 根据参数运行不同的脚本
func (option Option) Run() bool {
	if option.DB {
		DB()
		return true
	}
	if option.Port != 0 {
		Port(option.Port)
		return false
	}
	if option.Load != "" {
		Load()
		return true
	}
	return false
}

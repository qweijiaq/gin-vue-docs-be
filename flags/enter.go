package flags

import "flag"

type Option struct {
	DB   bool   // 迁移数据库
	Port int    // 设置端口号
	Load string // 导入数据集
}

func Parse() (option *Option) {
	option = new(Option)
	flag.BoolVar(&option.DB, "db", false, "迁移数据库")
	flag.IntVar(&option.Port, "port", 0, "程序运行的端口")
	flag.StringVar(&option.Load, "load", "", "导入 SQL 数据集")
	flag.Parse()

	return option
}

// Run 根据参数运行不同的脚本
func (option Option) Run() bool {
	if option.DB {
		DB()
		return true // return true 表示中断，即迁移数据表后不会运行程序
	}
	if option.Port != 0 {
		Port(option.Port)
		return false // return true 表示延续，即绑定指定端口后运行程序
	}
	if option.Load != "" {
		Load(option.Load)
		return true // return true 表示中断，即导入 SQL 文件后不会运行程序
	}
	return false
}

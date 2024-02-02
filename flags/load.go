package flags

import (
	"gvd_server/global"
	"os"
	"strings"
)

func Load(sqlPath string) {
	byteData, err := os.ReadFile(sqlPath)
	if err != nil {
		global.Log.Fatalf("%s err: %s", sqlPath, err.Error())
	}

	// 一定要按照 \r\n 分割
	sqlList := strings.Split(string(byteData), ";\r\n")
	for _, sql := range sqlList {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err = global.DB.Exec(sql).Error
		if err != nil {
			global.Log.Errorf("%s err:%s", sql, err.Error())
			continue
		}
	}

	global.Log.Infof("%s SQL 导入成功", sqlPath)
}

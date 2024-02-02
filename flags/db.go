package flags

import (
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
)

func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},           // 用户表
			&models.RoleModel{},           // 角色表
			&models.DocModel{},            // 文档表
			&models.UserCollectDocModel{}, // 用户收藏文档表 -- 如果有自定义连接表，一定要放在被连接的两张表的后面
			&models.RoleDocModel{},        // 角色文档表
			&models.ImageModel{},          // 图片表
			&models.UserPwdDocModel{},     // 用户密码文档表（针对用户需要密码解锁的文档）
			&models.LoginModel{},          // 登录表
			&models.DocDataModel{},        // 文档数据表
		)
	if err != nil {
		logrus.Fatalf("数据库迁移失败 error: %s", err.Error())
	}
	logrus.Infof("数据库迁移成功！")
}

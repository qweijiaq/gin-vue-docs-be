package flags

import (
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
)

func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.RoleModel{},
			&models.DocModel{},
			&models.UserCollectDocModel{}, // 如果有自定义连接表，一定要放在被连接的两张表的后面
			&models.RoleDocModel{},
			&models.ImageModel{},
			&models.UserPwdDocModel{},
			&models.LoginModel{},
			&models.DocDataModel{},
		)
	if err != nil {
		logrus.Fatalf("数据库迁移失败 error: %s", err.Error())
	}
	logrus.Infof("数据库迁移成功！")
}

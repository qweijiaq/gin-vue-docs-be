package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/response"
)

// UserRemoveView 删除用户
// @Tags 用户管理
// @Summary 删除用户
// @Description 从数据库中删除用户
// @Param token header string true "token"
// @Router /api/users [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.IDListRequest
	log := log_stash.NewAction(c)
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	var userList []models.UserModel
	global.DB.Find(&userList, cr.IDList)

	if len(cr.IDList) != len(userList) {
		log.Error("用户删除失败")
		log.SetItemErr("失败原因", "数据一致性校验不通过")
		response.FailWithMsg("数据一致性校验不通过", c)
		return
	}
	for _, model := range userList {
		err = UserRemoveService(model)
		if err != nil {
			log.Error("用户删除失败")
			log.SetItemInfo("userName", model.UserName)
			log.SetItemErr("失败原因", err.Error())
			logrus.Errorf("删除用户 %s 失败 err: %s", model.UserName, err.Error())
		} else {
			log.Info("用户删除成功")
			log.SetItemInfo("userName", model.UserName)
			logrus.Infof("删除用户 %s 成功", model.UserName)
		}
	}
	response.OKWithMsg(fmt.Sprintf("批量删除成功，共删除%d个用户", len(cr.IDList)), c)
	return

}

func UserRemoveService(user models.UserModel) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// imageModel 连带删除
		var imageList []models.ImageModel
		tx.Find(&imageList, "userID = ?", user.ID)
		if len(imageList) > 0 {
			err = tx.Delete(&imageList).Error
			if err != nil {
				return err
			}
		}
		// loginModel 不用连带删除

		// UserCollDocModel 连带删除
		var userCollList []models.UserCollectDocModel
		tx.Find(&userCollList, "user_id = ?", user.ID)
		if len(userCollList) > 0 {
			err = tx.Delete(&userCollList).Error
			if err != nil {
				return err
			}
		}
		// UserPwdDocModel 连带删除
		var userPwdList []models.UserPwdDocModel
		tx.Find(&userPwdList, "user_id = ?", user.ID)
		if len(userPwdList) > 0 {
			err = tx.Delete(&userPwdList).Error
			if err != nil {
				return err
			}

		}
		err = tx.Delete(&user).Error
		return err
	})
	return err
}

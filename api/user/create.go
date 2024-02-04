package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/response"
	"gvd_server/utils/encryption"
	"time"
)

// UserCreateRequest 创建用户时的请求参数
type UserCreateRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required" label:"密码"`  // 密码
	Nickname string `json:"nickname" label:"昵称"`                     // 昵称
	RoleID   uint   `json:"roleID" binding:"required" label:"角色"`    // 角色 ID
}

// UserCreateView 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户，只能管理员创建
// @Param data body UserCreateRequest true "参数"
// @Param token header string true "token"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest     // cr -> controller request
	err := c.ShouldBindJSON(&cr) // ShouldBindJSON 参数通过 JSON 传入（通常是 POST、PUT 请求）
	if err != nil {
		response.FailWithInValidError(err, &cr, c)
		return
	}

	log := log_stash.NewAction(c)
	byteData, _ := json.Marshal(cr)
	log.SetItemInfo("创建用户传入的参数", string(byteData))

	err = CreateUser(models.UserModel{
		Username:  cr.Username,
		Password:  cr.Password,
		Nickname:  cr.Nickname,
		IP:        c.RemoteIP(),
		RoleID:    cr.RoleID,
		LastLogin: time.Now(),
	}, &log)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OKWithMsg("用户创建成功", c)
	return
}

func CreateUser(user models.UserModel, log *log_stash.Action) (err error) {
	err = global.DB.Take(&user, "userName = ?", user.Username).Error
	if err == nil {
		log.Error("创建用户失败")
		log.SetItemInfo("username", user.Username)
		log.SetItemErr("失败原因", "该用户名已存在")
		return errors.New("该用户名已存在")
	}
	log.SetItemInfo("是否自动生成昵称", user.Nickname == "")
	if user.Nickname == "" {
		// 昵称不存在, 按规律生成
		var maxID uint
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxID)
		user.Nickname = fmt.Sprintf("用户_%d", maxID+1)
		log.SetItemInfo("自动生成昵称", user.Nickname)
	}
	var role models.RoleModel
	err = global.DB.Take(&role, user.RoleID).Error
	if err != nil {
		log.Error("创建用户失败")
		log.SetItemInfo("角色ID", user.RoleID)
		log.SetItemErr("失败原因", "该角色不存在")
		return errors.New("该角色不存在")
	}
	user.Password = encryption.HashPwd(user.Password) // 加密密码
	err = global.DB.Create(&user).Error
	if err != nil {
		global.Log.Error(err)
		log.Error("用户创建失败")
		log.SetItemErr("失败原因", err.Error())
		return errors.New("用户创建失败")
	}
	log.Info("用户创建成功")
	log.SetItemInfo("username", user.Username)
	return nil
}

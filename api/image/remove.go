package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/response"
	"os"
)

// ImageRemoveView 删除图片
// @Tags 图片管理
// @Summary 删除图片
// @Description 删除图片, 支持批量删除
// @Param token header string true "token"
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} response.Response{}
func (ImageApi) ImageRemoveView(c *gin.Context) {
	var cr models.IDListRequest
	log := log_stash.NewAction(c)
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithMsg("参数错误", c)
		return
	}

	var imageList []models.ImageModel
	global.DB.Find(&imageList, cr.IDList)

	if len(cr.IDList) != len(imageList) {
		log.Error("图片删除失败")
		log.SetItemErr("失败原因", "数据一致性校验不通过")
		response.FailWithMsg("数据一致性校验不通过", c)
		return
	}

	for _, model := range imageList {
		imageRemove(model, c)
	}
	response.OKWithMsg(fmt.Sprintf("批量删除成功，共删除%d张图片", len(cr.IDList)), c)
}

// 删除图片的时候，发现有多个相同的 hash，那就只删除记录
func imageRemove(image models.ImageModel, c *gin.Context) {
	log := log_stash.NewAction(c)

	var count int64
	global.DB.Model(models.ImageModel{}).
		Where("hash = ?", image.Hash).Count(&count)
	// count 的值肯定是大于等于 1 的
	// 大于等于 2 那就只删记录
	// 等于 1 那就删记录，并且删文件
	if count == 1 {
		err := os.Remove(image.Path)
		if err != nil {
			log.Error("图片删除失败")
			log.SetItemInfo("imagePath", image.Path)
			log.SetItemErr("失败原因", err.Error())
			global.Log.Errorf("删除文件 %s 错误 %s", image.Path, err.Error())
		} else {
			log.Info("图片删除成功")
			log.SetItemInfo("imagePath", image.Path)
			global.Log.Infof("删除文件 %s 成功", image.Path)
		}
	}
	global.DB.Delete(&image)
}

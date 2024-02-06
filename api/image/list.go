package image

import (
	"github.com/gin-gonic/gin"
	"gvd_server/models"
	"gvd_server/service/common/list"
	"gvd_server/service/common/response"
)

// ImageListResponse 请求图片列表时的响应
type ImageListResponse struct {
	models.ImageModel
	WebPath  string `json:"webPath"`  // 图片线上路径
	Nickname string `json:"nickname"` // 上传图片的用户昵称
}

// ImageListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 获取图片列表
// @Param data query models.Pagination true "参数"
// @Param token header string true "token"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[ImageListResponse]}
func (ImageApi) ImageListView(c *gin.Context) {
	var cr models.Pagination
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithInValidError(err, &cr, c)
		return
	}
	_list, count, _ := list.QueryList(models.ImageModel{}, list.Option{
		Pagination: cr,                    // 分页查询
		Likes:      []string{"fileName"},  // 根据 fileName 模糊查询
		Preload:    []string{"UserModel"}, // 预加载用户表
	})
	var imageList = make([]ImageListResponse, 0)
	for _, model := range _list {
		imageList = append(imageList, ImageListResponse{
			ImageModel: model,
			WebPath:    model.WebPath(),
			Nickname:   model.UserModel.Nickname,
		})
	}
	response.OKWithList(imageList, count, c)
}

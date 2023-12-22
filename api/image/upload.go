package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/response"
	"gvd_server/utils/hash"
	"gvd_server/utils/jwts"
	"path"
	"strings"
	"time"
)

var ImageWhiteList = []string{
	"jpg",
	"png",
	"jpeg",
	"gif",
	"svg",
	"webp",
}

// ImageUploadView 上传图片
// @Tags 图片管理
// @Summary 上传图片
// @Description 上传图片到本地文件夹
// @Param token header string true "token"
// @Param image formData file true "文件上传"
// @Router /api/imags [post]
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} response.Response{}
func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		response.FailWithMsg("图片参数错误", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)
	savePath := path.Join("uploads", claims.NickName, fileHeader.Filename)

	// 白名单判断
	if !InImageWhiteList(fileHeader.Filename, ImageWhiteList) {
		response.FailWithMsg("文件非法", c)
		return
	}

	// 文件大小判断  2MB
	if fileHeader.Size > int64(2*1024*1024) {
		response.FailWithMsg("文件过大", c)
		return
	}

	// 计算文件hash
	file, _ := fileHeader.Open()
	fileHash := hash.FileMd5(file)
	// 重复文件判断
	var imageModel models.ImageModel
	err = global.DB.Take(&imageModel, "hash = ?", fileHash).Error
	// 如果当前 hash 值不存在，则要上传图片并入库
	// 如果当前 hash 值存在，只需要入库，入库的 path 和已存在的相同
	if err != nil {
		// 没有
		// 判断一下，数据库里面有没有这个路径的图片
		var count int64
		global.DB.Model(models.ImageModel{}).
			Where("path = ?", savePath).Count(&count)
		if count > 0 {
			// 存在重名的情况，那么这个时候就需要改一下文件名
			// 123.png   ->  123_1688054761.png
			// 12.png.png  ->  12.png_1688054761.png
			fileHeader.Filename = ReplaceFileName(fileHeader.Filename)
			savePath = path.Join("uploads", claims.NickName, fileHeader.Filename)
		}

		err = c.SaveUploadedFile(fileHeader, savePath)
		if err != nil {
			global.Log.Errorf("%s 文件保存错误 %s", savePath, err)
			response.FailWithMsg("上传图片错误", c)
			return
		}
	} else {
		// 有，修改入库的path
		savePath = imageModel.Path
	}
	// 使用这个hash对数据库里面记录的图片进行查询
	// 用户删除图片的时候，发现有多个相同的hash，那就只删除记录
	imageModel = models.ImageModel{
		UserID:   claims.UserID,
		FileName: fileHeader.Filename,
		Size:     fileHeader.Size,
		Path:     savePath,
		Hash:     fileHash,
	}
	// 针对上传成功的图片写库
	err = global.DB.Create(&imageModel).Error
	if err != nil {
		global.Log.Errorln(err)
		response.FailWithMsg("文件上传失败", c)
		return
	}

	err = c.SaveUploadedFile(fileHeader, savePath)
	if err != nil {
		global.Log.Errorf("%s 文件保存错误 %s", savePath, err)
		response.FailWithMsg("上传图片失败", c)
		return
	}
	response.OK(imageModel.WebPath(), "图片上传成功", c)

}

// InImageWhiteList 判断一个图片是否在白名单中
func InImageWhiteList(fileName string, whiteList []string) bool {
	// 截取文件后缀
	_list := strings.Split(fileName, ".")
	if len(_list) < 2 {
		return false
	}
	suffix := strings.ToLower(_list[len(_list)-1])
	for _, s := range whiteList {
		if suffix == s {
			return true
		}
	}
	return false
}

// ReplaceFileName 修改文件名，加上时间戳
// 例如, tupian.png -> tupian_1688054761.png
func ReplaceFileName(oldFileName string) string {
	// 123.png
	_list := strings.Split(oldFileName, ".")
	// [123   png] -> [123 _1688054761  png]
	lastIndex := len(_list) - 2
	var newList []string
	for i, s := range _list {
		if i == lastIndex {
			newList = append(newList, fmt.Sprintf("%s_%d", s, time.Now().Unix()))
			continue
		}
		newList = append(newList, s)
	}
	return strings.Join(newList, ".")
}

package logs

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/list"
	"gvd_server/service/common/response"
	"time"
)

type LogListRequest struct {
	models.Pagination
	Level    log_stash.Level   `json:"level" form:"level"`       // 日志查询的等级
	Type     log_stash.LogType `json:"type" form:"type"`         // 日志的类型   1 登录日志  2 操作日志  3 运行日志
	IP       string            `json:"ip" form:"ip"`             // 根据ip查询
	UserID   uint              `json:"userID" form:"userID"`     // 根据用户id查询
	Addr     string            `json:"addr" form:"addr"`         // 感觉地址查询
	Date     string            `json:"date" form:"date"`         // 查某一天的，格式是年月日
	Status   int               `json:"status" form:"status"`     // 登录状态查询  1  成功  2 失败
	UserName string            `json:"userName" form:"userName"` // 查用户名
}

// LogListView 日志列表
// @Tags 日志管理
// @Summary 日志列表
// @Description 日志列表
// @Param data query LogListRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[log_stash.LogModel]}
func (LogApi) LogListView(c *gin.Context) {
	var cr LogListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	var query = global.DB.Where("")
	if cr.Date != "" {
		_, dateTimeErr := time.Parse("2006-01-02", cr.Date)
		if dateTimeErr != nil {
			response.FailWithMsg("时间格式错误", c)
			return
		}
		query.Where("date(createdAt) = ?", cr.Date)
	}
	if cr.Status == 1 {
		query.Where("status = ?", true)
	}
	if cr.Status == 2 {
		query.Where("status = ?", false)
	}

	_list, count, _ := list.QueryList(log_stash.LogModel{
		Type:     cr.Type,
		Level:    cr.Level,
		IP:       cr.IP,
		Addr:     cr.Addr,
		UserID:   cr.UserID,
		UserName: cr.UserName,
	}, list.Option{
		Pagination: cr.Pagination,
		Where:      query,
		Likes:      []string{"title", "userName"},
	})
	response.OKWithList(_list, count, c)
}

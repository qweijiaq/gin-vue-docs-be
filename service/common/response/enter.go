package response

import (
	"github.com/gin-gonic/gin"
	"gvd_server/utils/valid"
	"net/http"
)

type Code int

// Response 接口响应的结构，内部包含状态码 Code、数据 Data、信息 Msg
type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// ListResponse 列表数据 -- 通常针对一些 GET 请求，如后台管理获取用户列表、文档列表等
type ListResponse[T any] struct {
	List  []T `json:"list"`
	Count int `json:"count"`
}

const (
	SUCCESS     Code = 0 // 成功
	ErrCode     Code = 7 // 系统错误
	InValidCode Code = 9 // 参数校验错误
)

// OK 响应为成功时的数据和消息
func OK(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: SUCCESS, Data: data, Msg: msg})
}

// OKWithMsg 响应为成功时的消息
func OKWithMsg(msg string, c *gin.Context) {
	OK(map[string]any{}, msg, c)
}

// OKWithData 响应为成功时的数据
func OKWithData(data any, c *gin.Context) {
	OK(data, "成功", c)
}

// OKWithList 响应为成功时的列表数据
func OKWithList[T any](list []T, count int, c *gin.Context) {
	if len(list) == 0 {
		list = []T{}
	}
	OK(ListResponse[T]{
		List:  list,  // 数据列表
		Count: count, // 数据总数
	}, "成功", c)
}

// Fail 响应为失败时的数据和消息
func Fail(code Code, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: code, Data: data, Msg: msg})
}

// FailWithMsg 响应为失败时的信息
func FailWithMsg(msg string, c *gin.Context) {
	Fail(ErrCode, map[string]any{}, msg, c)
}

// FailWithData 响应为失败时的数据
func FailWithData(data any, c *gin.Context) {
	Fail(ErrCode, data, "系统错误", c)
}

// FailWithError 响应为校验失败时的错误信息
func FailWithError(err error, c *gin.Context) {
	errorMsg := valid.Error(err)
	Fail(InValidCode, map[string]any{}, errorMsg, c)
}

// FailWithInValidError 响应为参数校验失败时的数据和信息，数据是校验错误对应的字段
func FailWithInValidError(err error, obj any, c *gin.Context) {
	errorMsg, data := valid.InValidError(err, obj)
	Fail(InValidCode, data, errorMsg, c)
}

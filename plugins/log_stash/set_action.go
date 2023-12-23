package log_stash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"io"
	"reflect"
	"strings"
)

type Action struct {
	ip          string
	addr        string
	userName    string
	serviceName string
	userID      uint
	level       Level
	title       string
	itemList    []string
	model       *LogModel // 创建之后赋值给它，用于后期更新
	token       string
	logType     LogType
}

func NewAction(c *gin.Context) Action {
	ip := c.ClientIP()
	addr := "局域网"
	action := Action{
		ip:      ip,
		addr:    addr,
		logType: ActionType,
	}
	token := c.Request.Header.Get("token")
	action.SetToken(token)
	return action
}

// Info Warn Error 三个属于中断式函数
func (action *Action) Info(title string) {
	action.level = Info
	action.title = title
	action.save()
}
func (action *Action) Warn(title string) {
	action.level = Warning
	action.title = title
	action.save()
}
func (action *Action) Error(title string) {
	action.level = Error
	action.title = title
	action.save()
}

// SetToken 设置 token
func (action *Action) SetToken(token string) {
	action.token = token
}

// SetItemInfo 设置 info 详情
func (action *Action) SetItemInfo(label string, value any) {
	action.setItem(label, value, Info)
}

// SetItemWarn 设置 warn 详情
func (action *Action) SetItemWarn(label string, value any) {
	action.setItem(label, value, Warning)
}

// SetItemErr 设置 error 详情
func (action *Action) SetItemErr(label string, value any) {
	action.setItem(label, value, Error)
}

func (action *Action) setItem(label string, value any, level Level) {
	// 判断类型
	_type := reflect.TypeOf(value).Kind()
	switch _type {
	case reflect.Struct, reflect.Map, reflect.Slice:
		// 可以设置关键字，然后有关键字的高亮显示，或者有颜色的字符
		// 颜色有两种方案
		// 1. html字符  <span style="color:red" />
		// 2. 控制字符 \033[31m xxx \033[0m
		byteData, _ := json.Marshal(value)
		action.itemList = append(action.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%s</div></div>", level.String(), label, string(byteData)))
	default:
		action.itemList = append(action.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%v</div></div>", level.String(), label, value))
	}
}

// SetUrl 设置一组url
func (action *Action) SetUrl(title, url string) {
	// 如果要使用html显示，一定要注意xss问题
	action.itemList = append(action.itemList, fmt.Sprintf("<div class=\"log_link\"><a target=\"_blank\" href=\"%s\">%s</a></div>", url, title))
}

// SetRequestHeader 设置请求头
func (action *Action) SetRequestHeader(c *gin.Context) {
	header := c.Request.Header.Clone()
	byteData, _ := json.Marshal(header)
	action.itemList = append(action.itemList, fmt.Sprintf(
		`<div class="log_request_header">
	<div class="log_request_body">
		<pre class="log_json_body">%s</pre>
	</div>
</div>`, string(byteData)))
}

// SetRequest 设置一组入参
func (action *Action) SetRequest(c *gin.Context) {
	// 请求方法
	method := c.Request.Method
	// 请求路径
	path := c.Request.URL.String()
	// 请求体
	// 关于请求体的问题，拿了之后要还回去
	// 一定要在参数绑定之前调用
	byteData, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteData))
	action.itemList = append(action.itemList, fmt.Sprintf(
		`<div class="log_request">
	<div class="log_request_head">
		<span class="log_request_method">%s</span>
		<span class="log_request_path">%s</span>
	</div>
	<div class="log_request_body">
		<pre class="log_json_body">%s</pre>
	</div>
</div>`, method, path, string(byteData)))
}

// SetResponse 设置一组出参
func (action *Action) SetResponse(c *gin.Context) {
	c.Set("action", action)
}
func (action *Action) SetResponseContent(response string) {
	action.itemList = append(action.itemList, fmt.Sprintf(`
<div class="log_response">
	<pre class="log_json_body">%s</pre>
</div>
`, response))
}

// SetFlush 保存level不变，更新
func (action *Action) SetFlush() {
	action.level = action.model.Level
	action.save()
}

// SetUpload 文件上传的函数
func (action *Action) SetUpload(c *gin.Context) {
	forms, err := c.MultipartForm()
	if err != nil {
		action.SetItemErr("form参数错误", err.Error())
		return
	}
	for s, headers := range forms.File {
		action.itemList = append(action.itemList, fmt.Sprintf(
			`<div class="log_upload">
        <div class="log_upload_head">
            <span class="log_upload_file_key">%s</span>
            <span class="log_upload_file_name">%s</span>
            <span class="log_upload_file_size">%s</span>
        </div>
    </div>`, s, headers[0].Filename, FormatBytes(headers[0].Size)))
	}

}

func (action *Action) SetImage(url string) {
	action.itemList = append(action.itemList, fmt.Sprintf("<div class=\"log_image\"/><img src=\"%s\"></div>", url))
}

func (action *Action) save() {
	content := strings.Join(action.itemList, "\n")
	if action.token != "" {
		jwyPayLoad := parseToken(action.token)
		if jwyPayLoad != nil {
			action.userID = jwyPayLoad.UserID
			action.userName = jwyPayLoad.UserName
		}
	}
	if action.model == nil {
		action.model = &LogModel{
			IP:          action.ip,
			Addr:        action.addr,
			Level:       action.level,
			Title:       action.title,
			Content:     content,
			UserID:      action.userID,
			UserName:    action.userName,
			ServiceName: action.serviceName,
			Type:        action.logType,
		}
		global.DB.Create(action.model)
		// 如果不对content进行置空，那么content会重复
		action.itemList = []string{}
		return
	}
	// 更新操作
	global.DB.Model(action.model).Updates(LogModel{
		Level:   action.level,
		Title:   action.title,
		Content: action.model.Content + "\n" + content,
	})

}

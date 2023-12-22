package log_stash

import "encoding/json"

type LogType int

const (
	LoginType   LogType = 1 // 登录日志
	ActionType  LogType = 2 // 操作日志
	RuntimeType LogType = 3 // 运行日志
)

// String 转字符串
func (t LogType) String() string {
	switch t {
	case LoginType:
		return "loginType"
	case ActionType:
		return "actionType"
	case RuntimeType:
		return "runtimeType"
	}
	return ""
}

// MarshalJSON 自定义类型转换为JSON
func (t LogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

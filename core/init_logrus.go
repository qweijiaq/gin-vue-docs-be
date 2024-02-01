package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogRequest struct {
	LogPath  string // 日志的目录
	AppName  string // 应用名称
	NoDate   bool   // 是否需要按照时间分割 -- 默认是
	NoErr    bool   // 是否单独存放 error 日志 -- 默认是
	NoGlobal bool   // 是否替换全局 logrus -- 默认是
}

type LogFormatter struct {
}

// Format 实现 Format(entry *logrus.Entry) ([]byte, error) 方法
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的 level 展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var err error
	if entry.HasCaller() {
		funcVal := entry.Caller.Function                                                 // 日志输出来自哪一个函数
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line) // 日志输出来自哪个文件的哪一行
		// 自定义输出格式
		_, err = fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
		if err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil
}

// DateHook 按照时间分割写入日志文件的 hook
type DateHook struct {
	file     *os.File
	fileDate string // 判断日期切换目录
	logPath  string
	appName  string
}

func (DateHook) Levels() []logrus.Level {
	return logrus.AllLevels // 全部等级都触发钩子函数
}
func (hook DateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		_, err := hook.file.Write([]byte(line))
		if err != nil {
			return err
		}
		return nil
	}
	// 时间不等
	err := hook.file.Close()
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(hook.logPath, timer), os.ModePerm)
	if err != nil {
		return err
	}
	filename := path.Join(hook.logPath, timer, hook.appName+".logs")

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	_, err = hook.file.Write([]byte(line))
	if err != nil {
		return err
	}
	return nil
}

// ErrorHook 将error级别的日志写入到具体文件中
type ErrorHook struct {
	file     *os.File
	fileDate string //判断日期切换目录
	logPath  string
	appName  string
}

func (ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}
func (hook ErrorHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		_, err := hook.file.Write([]byte(line))
		if err != nil {
			return err
		}
		return nil
	}
	// 时间不等
	err := hook.file.Close()
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(hook.logPath, timer), os.ModePerm)
	if err != nil {
		return err
	}
	filename := path.Join(hook.logPath, timer, "err.logs")

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	_, err = hook.file.Write([]byte(line))
	if err != nil {
		return err
	}
	return nil
}

func InitLogger(requestList ...LogRequest) *logrus.Logger {
	var request LogRequest
	if len(requestList) > 0 {
		request = requestList[0]
	}
	if request.LogPath == "" {
		request.LogPath = "logs"
	}
	if request.AppName == "" {
		request.AppName = "gin-vue-docs"
	}
	mLog := logrus.New()               // 新建一个实例
	mLog.SetOutput(os.Stdout)          // 设置输出类型
	mLog.SetReportCaller(true)         // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) // 设置自己定义的 Formatter
	mLog.SetLevel(logrus.DebugLevel)   // 设置最低的 Level

	// DateHook 必须在 ErrorHook 之前 Add, 否则当文件夹还未创建时, 会出现错误
	if !request.NoDate {
		mLog.AddHook(&DateHook{
			logPath: request.LogPath,
			appName: request.AppName,
		})
	}
	if !request.NoErr {
		mLog.AddHook(&ErrorHook{
			logPath: request.LogPath,
			appName: request.AppName,
		})
	}
	if !request.NoGlobal {
		InitDefaultLogger()
	}
	return mLog
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)          // 设置输出类型
	logrus.SetReportCaller(true)         // 开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) // 设置自己定义的 Formatter
	logrus.SetLevel(logrus.DebugLevel)   // 设置最低的 Level
}

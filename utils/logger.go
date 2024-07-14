package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

var logger = &MyLogger{logrus.StandardLogger()}

func GetLogger() *MyLogger {
	return logger
}

func init() {
	// 设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	// 配置日志格式
	logrus.SetFormatter(&CustomFormatter{})
}

// noinspection SpellCheckingInspection(忽略单词检查，去掉波浪线，强迫症行为)
func getCallerLine() (string, string, int) {
	pcs := make([]uintptr, 25)
	depth := runtime.Callers(4, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		if strings.Contains(f.Function, "logrus") || strings.Contains(f.Function, "MyLogger") {
			continue
		}

		file, funcName, line := f.File, f.Function, f.Line
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				file = file[i+1:]
				break
			}

		}
		for i := len(funcName) - 1; i > 0; i-- {
			if funcName[i] == '/' {
				funcName = funcName[i+1:]
				break
			}
		}
		for i := 0; i < len(funcName)-2; i++ {
			if funcName[i] == '.' {
				funcName = funcName[i+1:]
				break
			}
		}
		return file, funcName, line
	}
	return "", "", 0

}

// CustomFormatter 自定义日志格式化器
type CustomFormatter struct{ logrus.TextFormatter }

// noinspection SpellCheckingInspection
// 定制日志序列化器
// 1、定制API服务日志序列化器
// 2、定制定时任务日志序列化器
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取行号
	file, funcName, line := getCallerLine()
	// go1.16以下版本不支持Format("2006-01-02 15:04:05,000")这种写法, 会获取到毫秒数为0
	// 所以还是用.来格式化, 期望打印的是逗号','格式, 替换一下字符串
	currTime := entry.Time.Format("2006-01-02 15:04:05,000")

	// 使用自定义的 FieldMap
	fields := make(map[string]interface{})
	for key, value := range entry.Data {
		fields[key] = value
	}
	uuidVal, ok := entry.Data["uuid"]
	message := ""
	if ok && uuidVal != nil { // 用于 API 服务
		message = fmt.Sprintf("[%s] [%s] [%s-%d-%s] request_id: %s %s\n",
			currTime,    // 时间戳格式
			entry.Level, // 日志级别
			file,        // 文件名
			line,        // 行号
			funcName,    // 函数名
			uuidVal,
			entry.Message, // 日志消息
		)
	} else { // 用于定时任务服务
		message = fmt.Sprintf("[%s] [%s] [%s-%d-%s] %s\n",
			currTime,      // 时间戳格式
			entry.Level,   // 日志级别
			file,          // 文件名
			line,          // 行号
			funcName,      // 函数名
			entry.Message, // 日志消息
		)
	}

	return []byte(message), nil
}

// MyLogger 自定义日志器，扩展 Logger
type MyLogger struct {
	*logrus.Logger
}

// 重写以下最常用的日志类型，添加用户请求ID到日志中，这对于api服务很是关键（尤其对于高并发场景），便于追踪定位问题
// 其他日志级别如果有需要同理添加即可

func (logger *MyLogger) Printf(format string, args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Printf(format, args...)
}

// noinspection SpellCheckingInspection(忽略单词检查，去掉波浪线，强迫症行为)
func (logger *MyLogger) Debugf(format string, args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Debugf(format, args...)
}

func (logger *MyLogger) Debug(args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Debug(args...)
}

// noinspection SpellCheckingInspection(忽略单词检查，去掉波浪线，强迫症行为)
func (logger *MyLogger) Infof(format string, args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Infof(format, args...)
}

func (logger *MyLogger) Info(args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Info(args...)
}

// noinspection SpellCheckingInspection(忽略单词检查，去掉波浪线，强迫症行为)
func (logger *MyLogger) Warnf(format string, args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Warnf(format, args...)
}

func (logger *MyLogger) Warn(args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Warn(args...)
}

// noinspection SpellCheckingInspection(忽略单词检查，去掉波浪线，强迫症行为)
func (logger *MyLogger) Warningf(format string, args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Warningf(format, args...)
}

func (logger *MyLogger) Warning(args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Warning(args...)
}

// noinspection SpellCheckingInspection(忽略单词检查，去掉波浪线，强迫症行为)
func (logger *MyLogger) Errorf(format string, args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Errorf(format, args...)
}

func (logger *MyLogger) Error(args ...interface{}) {
	logger.Logger.WithFields(logrus.Fields{"uuid": GetCurrentUUID()}).Error(args...)
}

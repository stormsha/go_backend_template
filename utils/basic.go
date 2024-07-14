package utils

import (
	uuid "github.com/satori/go.uuid"
	"path"
	"runtime"
	"stormsha.com/gbt/config"
)

var conf = config.Setting

type ResponseData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func GetProjectRootPath() string {
	// 1. 获取当前文件的路径
	_, currFileName, _, _ := runtime.Caller(0)               // 获取当前函数的路径信息
	localProjectRootPath := path.Dir(path.Dir(currFileName)) // 获取当前项目根路径
	return localProjectRootPath
}

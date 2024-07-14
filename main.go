package main

import (
	"flag"
	"fmt"
	"stormsha.com/gbt/controller"
	_ "stormsha.com/gbt/controller/router"
	"stormsha.com/gbt/cron"
	_ "stormsha.com/gbt/docs"
	"stormsha.com/gbt/utils"
)

var (
	deployName string
	helpFlag   bool
)

func init() {
	flag.StringVar(&deployName, "n", "api", "deploy name (api/cron) \n"+
		" api: web api\n"+
		"cron: 定时任务\n")
	flag.BoolVar(&helpFlag, "help", false, "this help")
}

// @title           go backend template
// @version         v1.0
// @description     Go Echo Web Api 服务高效开发模板
// @contact.name   	stormsha
// @contact.url    	https://www.stormsha.com
// @host      		localhost:8080
// @BasePath  		/api
func main() {
	// 通过参数分开部署
	flag.Parse()
	if helpFlag {
		flag.Usage()
		return
	}
	// 初始化日志器
	logger := utils.GetLogger()
	fmt.Println("___________Logger 初始化成功!___________")
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("程序发成异常 : %v ", r)
		}
	}()
	logger.Info("初始化数据库连接开始...")
	utils.InitDataBaseConn() // 初始化数据库连接连接
	logger.Info("初始化数据库连接完成...")

	switch deployName {
	case "api":
		logger.Info("初始化echo...")
		controller.InitEcho() // 初始化echo
		logger.Info("开启echo服务 ...")
		controller.StartEcho("8080") //启动echo服务

	case "cron":
		logger.Info("初始化cron...")
		cron.InitCorn() // 初始化cron
		logger.Info("开启cron服务 ...")
		// 启动任务
		cron.StartCron() // 初始化启动定时任务
		logger.Info("定时任务启动成功")
		select {} // 不让程序退出
	default:
		flag.Usage()
		return
	}
}

package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"stormsha.com/gbt/config/constant"
)

var app = echo.New() // 初始化echo实例

// Fail 设置请求失败消息体
func Fail(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": constant.ServiceStatusFail,
		"msg":  msg,
	})
}

// InitEcho 初始化echo实例
func InitEcho() {
	// 默认中间件-添加用户请求ID，用于高并发场景日志追踪
	app.Use(middleware.RequestID())
	// 默认中间件-允许跨域请求
	app.Use(middleware.CORS())
	// 默认中间件-配置echo的日志Logger
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format:           "[${time_custom}] [info] request_id = ${id} method=${method}, uri=${uri}, status=${status}, bytes_out=${bytes_out}\n",
			CustomTimeFormat: "2006-01-02 15:04:05,000",
		}))

	// 自定义中间件-记录用户请求信息：请求ID、接口、入参等等
	app.Use(MiddleWareLogger)
	// 自定义中间件-捕获未知异常
	app.Use(MiddleWareDefer)

	// 默认开启-echo的Debug模式
	app.Debug = conf.DEBUG
}

// StartEcho 启动echo服务
func StartEcho(port string) {
	// 创建 swagger 接口文档路由
	app.GET("/docs/*", echoSwagger.WrapHandler)
	// 启动服务并输出服务端口
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%v", port)))
}

package router

import (
	"github.com/labstack/echo/v4"
	"stormsha.com/gbt/controller"
)

// Sys 初始化一个组业务对象
type Sys struct{}

// 在init()函数中绑定路由对应的控制器函数
func init() {
	s := &Sys{}
	g := controller.Group("/api/v1/sys") // 根据前缀获取echo路由分组
	//g.GET("/swagger/*", echoSwagger.WrapHandler)
	g.GET("/version", wrapper(s.version))
}

// version 获取 Web Api 版本号
// @Summary      服务版本
// @Description  查看api服务当前版本
// @Tags         sys
// @Produce      json
// @Success      200
// @Router       /v1/sys/version [get]
func (s *Sys) version(ctx echo.Context) (interface{}, error) {
	// 可以把版本信息配置化，随着业务增长，会出现不同版本接口共存的阶段
	/* 例如：{
		"v1":["/api/v1/star", "/api/v1/love", "/api/v1/share"]
		"v2":["/api/v2/star", "/api/v2/love", "/api/v2/share"]
	}
	*/
	data := map[string][]string{
		"v1":  []string{"/api/v1/star", "/api/v1/love", "/api/v1/share"},
		"vv2": []string{"/api/v2/star", "/api/v2/love", "/api/v2/share"},
	}
	return data, nil
}

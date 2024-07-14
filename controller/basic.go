package controller

import (
	"github.com/labstack/echo/v4"
	"stormsha.com/gbt/config"
	"stormsha.com/gbt/utils"
)

var conf = config.Setting

var logger = utils.GetLogger()

// Group 获取路由分组（有则直接返回，无则创建并返回）
func Group(prefix string, middleware ...echo.MiddlewareFunc) *echo.Group {
	return app.Group(prefix, middleware...)
}

package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stormsha.com/gbt/config/constant"
	"stormsha.com/gbt/utils"
)

var logger = utils.GetLogger()

type myFunc func(echo.Context) (interface{}, error)

// 捕获接口未知异常，防止因某个接口异常导致服务挂掉
func wrapper(h myFunc) echo.HandlerFunc {

	return func(context echo.Context) error {
		r := utils.ResponseData{}
		data, err := h(context)
		if err != nil { // 未知异常统一返回状态
			r.Code = constant.ServiceStatusFail
			r.Msg = "服务异常"
		} else { // 请求成功统一返回状态
			r.Msg = ""
			r.Code = constant.ServiceStatusOk
			r.Data = data
		}

		return context.JSON(http.StatusOK, r)
	}
}

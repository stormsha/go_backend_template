package controller

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"stormsha.com/gbt/config/constant"
	"stormsha.com/gbt/utils"
	"strings"
	"sync"
)

// API锁中间件, 同一时间只执行一次POST,PUT,DELETE请求
var lock sync.Mutex

func MiddleWareApiLock(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		if method := context.Request().Method; method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
			lock.Lock()
			defer lock.Unlock()
			logger.Infof("%v请求加锁，url: %v", method, context.Request().URL)
		}
		return next(context)
	}
}

func MiddleWarePermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		header := context.Request().Header
		token := header.Get("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")
		user, err := utils.CheckUserPermission(token)
		if err != nil {
			re := utils.ResponseData{}
			re.Code = constant.ServiceStatusFail
			re.Msg = err.Error()
			return context.JSON(http.StatusOK, &re)
		} else {
			logger.Infof("token：%v, username：%v", token, user.UserName)
			// 记录用户信息到session
			utils.SetSessionUser(user)
			return next(context)
		}
	}
}

// MiddleWareLogger 记录请求日志
func MiddleWareLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		defer func() {
			utils.ClearCurrentSession()
		}()
		utils.CacheSession()
		req := context.Request()
		header := context.Request().Header
		requestID := context.Response().Header().Get(echo.HeaderXRequestID)
		// 复制出body
		body := ""
		cType := header.Get(echo.HeaderContentType)
		if strings.HasPrefix(cType, echo.MIMEApplicationJSON) {
			bodyB, err := io.ReadAll(req.Body)
			if err != nil {
				logger.Errorf("链路uuid: %v, 读取请求体错误 %v", requestID, err)
				re := utils.ResponseData{}
				re.Code = constant.ServiceStatusFail
				logger.Errorf("服务异常 %v", err.Error())
				re.Msg = "服务异常"
				return context.JSON(http.StatusOK, &re)
			}
			req.Body = ioutil.NopCloser(bytes.NewReader(bodyB))
			body = string(bodyB)
		}
		utils.SetSessionUUid(requestID)
		logger.Infof("url: %v, method: %v, content-type: %v, body: %v", req.URL, req.Method, cType, body)
		return next(context)
	}
}

// MiddleWareDefer 使用echo框架的中间件机制注册一个自定义异常处理中间件
func MiddleWareDefer(next echo.HandlerFunc) echo.HandlerFunc {
	// 返回一个新的HandlerFunc，该HandlerFunc将在中间件链中执行
	return func(ctx echo.Context) error {
		// 使用defer关键字定义一个延迟执行的匿名函数，用于捕获panic异常
		defer func() {
			// 捕获panic并赋值给变量r
			if r := recover(); r != nil {
				// 使用fmt.Errorf创建一个新的错误，包含panic的值
				err := fmt.Errorf("%v", r)

				// 获取当前 goroutine 的堆栈信息，用于调试和记录
				stack := debug.Stack()

				// 格式化错误信息和堆栈信息，记录到日志中
				logger.Errorf(fmt.Sprintf("[PANIC RECOVER] %v %v\n", err, stack))

				// 检查是否处于debug模式，如果是，则输出详细的堆栈信息
				if conf.DEBUG {
					logger.Errorf("%v", stack)
				}

				// 调用自定义的Fail函数，将错误信息返回给客户端
				Fail(ctx, err.Error())
			}
		}()

		// 调用下一个中间件或路由处理器，并捕获返回的错误
		err := next(ctx)

		// 如果存在未知错误，则记录错误日志
		if err != nil {
			// 记录错误信息到日志中
			logger.Error(err.Error())
			// 调用自定义的Fail函数，将错误信息返回给客户端
			return Fail(ctx, err.Error())
		}

		// 如果没有错误发生，则正常返回nil
		return nil
	}
}

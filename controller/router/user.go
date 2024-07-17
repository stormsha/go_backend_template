package router

import (
	"github.com/labstack/echo/v4"
	"stormsha.com/gbt/controller"
	"stormsha.com/gbt/model"
	"stormsha.com/gbt/view"
)

type User struct{}

// 在init()函数中绑定路由对应的控制器函数
func init() {
	u := &User{}
	// 根据前缀获取echo路由分组
	g1 := controller.Group("/api/v1", controller.MiddleWareApiLock)
	g1.POST("/register", wrapper(u.register))
	g1.POST("/login", wrapper(u.login))

	g2 := controller.Group("/api/v1/user", controller.MiddleWarePermission)
	g2.GET("/detail", wrapper(u.detail))
}

type Account struct {
	UserAccount  string `json:"user_account" example:"admin"`  // 账号
	UserPassword string `json:"user_password" example:"admin"` // 密码
}

// @Summary      注册
// @Description  用户注册
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 account body Account true "注册信息"
// @Success      200
// @Router       /v1/register [post]
// version 获取 Web Api 版本号
func (u *User) register(ctx echo.Context) (interface{}, error) {
	args := new(model.User)
	err := ctx.Bind(args)
	if err != nil {
		return nil, err
	}
	result, err := view.Register(args)
	return result, err
}

// @Summary      登录
// @Description  用户登录
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 account body Account true "用户信息"
// @Success      200
// @Router       /v1/login [post]
// version 获取 Web Api 版本号
func (u *User) login(ctx echo.Context) (interface{}, error) {
	args := new(model.User)
	err := ctx.Bind(args)
	if err != nil {
		return nil, err
	}
	result, err := view.Login(args)
	return result, err
}

// @Summary      用户信息
// @Description  获取用户信息
// @Tags         user
// @SWG.Path("/api/v1/user", "API v1 User")
// @Accept       json
// @Produce      json
// @Param 		 id query int true "用户ID"
// @Param Authorization header string true "Authorization token"
// @Success      200
// @Router       /v1/user/detail [get]
func (u *User) detail(ctx echo.Context) (interface{}, error) {
	args := new(model.User)
	err := ctx.Bind(args)
	if err != nil {
		return nil, err
	}
	result, err := view.Detail(args)
	return result, err
}

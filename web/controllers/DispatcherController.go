package controllers

import (
	"RBAC_GO/src/models"
	"RBAC_GO/src/service"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/v12/mvc"
)

type DispatcherController struct {
	//每个请求都由Iris自动绑定上下文，
	//记住，每次传入请求时，iris每次都会创建一个新的UserController，
	//所以所有字段都是默认的请求范围，只能设置依赖注入
	//自定义字段，如服务，对所有请求都是相同的（静态绑定）
	//和依赖于当前上下文的会话（动态绑定）。
	Ctx iris.Context
	//我们的UserService，它是一个接口
	//从主应用程序绑定。
	Service           service.UserService
}
func (c *DispatcherController) BeforeActivation(b mvc.BeforeActivation) {


	b.Handle("Post", "/login", "login")
	b.Handle("GET", "/error", "error")
	b.Handle("GET", "/main", "main")

}


// 设置主页为登录界面
func (c *DispatcherController) Get() error {
	return c.Ctx.View("login.html")
}
// 设置主页为登录界面
func (c *DispatcherController) PostLoginBy(loginAcct string, userPsWd string) models.AJAXResult {
	loginer := models.User{
		LoginAcct: loginAcct,
		UserPsWd: userPsWd,
	}
	login, err := c.Service.QueryLogin(&loginer)
	if err != nil {
		return models.AJAXResult{
			Success: false,
		}
	}
	fmt.Println(login)
	if login.Id != -1 {

	}
	return models.AJAXResult{
		Success: true,
	}
}
// 设置错误界面
func (c *DispatcherController) GetError() error {
	return c.Ctx.View("error.html")
}
// 设置登录界面
func (c *DispatcherController) login() error {
	return c.Ctx.View("login.html")
}


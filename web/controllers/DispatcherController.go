package controllers

import (
	"RBAC_GO/src/models"
	"RBAC_GO/src/service"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
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
	// session
	Session *sessions.Session
}
func (c *DispatcherController) BeforeActivation(b mvc.BeforeActivation) {

	b.Handle("POST", "/doAJAXLogin", "doAJAXLogin") // 验证登录，并且获取此登录用户的权限角色，权限
	b.Handle("GET", "/login", "login")
	b.Handle("GET", "/error", "error")
	b.Handle("GET", "/logout", "logout") // 登出，进行重定向，删除session
	b.Handle("GET", "/main", "main")
	b.Handle("GET", "/doLogin", "doLogin")
}


// 设置主页为登录界面
func (c *DispatcherController) Get() error {
	return c.Ctx.View("login.html")
}
// 设置错误界面
func (c *DispatcherController) GetError() error {
	return c.Ctx.View("error.html")
}
// 设置登录界面
func (c *DispatcherController) GetLogin() error {
	return c.Ctx.View("login.html")
}
// doAJAXLogin ajax登录验证
func (d *DispatcherController) doAJAXLogin() models.AJAXResult {
	result := models.AJAXResult{}
	var user = d.Ctx.FormValue("loginAcct")
	var password = d.Ctx.FormValue("userPsWd")
	users := models.User{
		LoginAcct: user,
		UserPsWd:  password,
	}
	fmt.Println(users)

	login, err := d.Service.QueryLogin(&users)
	if err != nil {
		result.Success = false
		return result
	} else {
		// 验证成功，存储session，并且获取用户权限信息
		fmt.Println(login)

		if login != (models.User{}) {

		}
		return result
	}

}

package controllers

import (
	"RBAC_GO/src/models"
	"RBAC_GO/src/service"
	"fmt"
	"github.com/kataras/iris"
)

type DispatcherController struct {
	//每个请求都由Iris自动绑定上下文，
	//记住，每次传入请求时，iris每次都会创建一个新的UserController，
	//所以所有字段都是默认的请求范围，只能设置依赖注入
	//自定义字段，如服务，对所有请求都是相同的（静态绑定）
	//和依赖于当前上下文的会话（动态绑定）。
	Ctx iris.Context
	//我们的UserService，它是一个接口

}

func Index(index iris.Party)  {
	// user路由组
	index.PartyFunc("/user",User)

	// Dispatcher 写的是/ 分组下的方法
	index.Get( "/", indexPage)
	index.Post("/doAJAXLogin",doAJAXLogin)
}
// 返回登录界面
func indexPage(ctx iris.Context) {
	ctx.View("login.html")
}
// ajax 传递数据，验证登录
func doAJAXLogin(ctx iris.Context) {
	var (
		user = ctx.FormValue("loginAcct")
		password = ctx.FormValue("userPsWd")

	)
	fmt.Println(user,password)
	users := models.User{
		LoginAcct: user,
		UserPsWd: password,
	}
	service.UserService.QueryLogin(&users)
	ctx.JSON(models.AJAXResult{
		Success: true,
	})

}
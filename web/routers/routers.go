/*
@Time : 2020/5/26 22:19
@Author : Firewine
@File : routers
@Software: GoLand
@Description:
*/
package routers

import (
	"RBAC_GO/web/controllers"
	"RBAC_GO/web/middleware"
	//"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)
// 路由方法
var mySecret = []byte("HS2JDFKhu7Y1av7b")
// 进行路由分组
func Routers(api *iris.Application){
	app := api.Party("/", middleware.CrsAuth()).AllowMethods(iris.MethodOptions)
	{
		app.Get("/", func(ctx iris.Context) { // 首页模块
			_ = ctx.View("login.html")
		})

		back := app.Party("/back")
		{


			back.Post("/doAJAXLogin",controllers.DoAJAXLogin).Name = "登录"
			//back.Use(irisyaag.New())
			middleware.InitCasbin()
			casbinMiddleware := middleware.New(middleware.Enforcer)               //casbin for xorm                                                   // <- IMPORTANT, register the middleware.
			app.Use(middleware.JWTMiddleware().Serve,casbinMiddleware.ServeHTTP) //登录验证
			back.Get("/logout",controllers.Logout).Name = "退出"
			back.PartyFunc("/user",controllers.UserControllers)
		}
	}
}
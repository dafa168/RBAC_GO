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
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
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
			back.HandleDir("/", "./web/static")
			back.Post("/doAJAXLogin",controllers.DoAJAXLogin).Name = "登录"
			//back.Use(irisyaag.New())
			casbin := middleware.InitCasbin()
			casbinMiddleware := middleware.New(casbin)               //casbin for xorm
			back.Use(middleware.JWTMiddleware().Serve,casbinMiddleware.ServeHTTP) //登录验证
			back.Get("/logout",controllers.Logout).Name = "退出"
			back.Get("/main",controllers.Main).Name = "back 主页"
			back.PartyFunc("/user", func(p router.Party) {
				//后台之 user管理
				p.Post("/deletes",controllers.Deletes).Name = "user deletes"
				p.Post("/delete",controllers.Delete).Name = "user delete"

			})
		}
	}
}
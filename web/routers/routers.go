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
	"github.com/kataras/iris"
)
// 路由方法
// 进行路由分组
func Routers(app *iris.Application) *iris.Application {
	// / 下 url
	app.PartyFunc("/", controllers.Dispatcher)
	return app
}
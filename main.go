/*
@Time : 2020/5/23 10:18
@Author : Firewine
@File : main
@Software: GoLand
@Description:
*/
package main

import (
	"RBAC_GO/configs"
	"RBAC_GO/web/middleware"
	"RBAC_GO/web/routers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	// 初始化中间件，xorm 引擎
	configs.MysqlEngine()
	//请求日志记录
	middleware.Logger(app)
	//注册api文档生成
	//middleware.Yaag(app)
	//加载 casbin 权限控制


	//加载静态网页路径
	middleware.PathLoad(app)
	// iris 的mvc
	mvc.Configure(app.Party("/"),routers.Dispatcher)
	app.Run(
		//开启web服务
		iris.Addr(":8080"),
		//实现更快的json序列化和更多优化
		iris.WithOptimizations,
		iris.WithCharset("utf-8"),
	)
}

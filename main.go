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
	"RBAC_GO/web/controllers"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
)

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	//初始化中间件,API文档
	yaag.Init(&yaag.Config{
		On:       true, //是否开启自动生成API文档功能
		DocTitle: "Iris",
		DocPath:  "./api/apidoc.html", //生成API文档名称存放路径
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	// 初始化中间件，xorm 引擎
	configs.MysqlEngine()

	//注册中间件
	//app.Use(irisyaag.New())
	//加载模板文件
	app.RegisterView(iris.HTML("./web/views/", ".html"))
	//加载静态资源
	app.StaticWeb("/","./web/static/")
	// 添加路由
	app.PartyFunc("/", controllers.Index)


	app.Run(
		//开启web服务
		iris.Addr(":8080"),
		//实现更快的json序列化和更多优化
		iris.WithOptimizations,
		iris.WithCharset("utf-8"),
	)
}

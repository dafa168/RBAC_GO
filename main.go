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
	"RBAC_GO/web/routers"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
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
	// 请求日志记录
	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(customLogger)
	//加载模板文件
	app.RegisterView(iris.HTML("./web/views/", ".html"))
	//加载静态资源
	app.HandleDir("/", "./web/static")
	// 添加路由
	routers.Routers(app)

	app.Run(
		//开启web服务
		iris.Addr("localhost:8080"),
		//实现更快的json序列化和更多优化
		iris.WithOptimizations,
		iris.WithCharset("utf-8"),
	)
}

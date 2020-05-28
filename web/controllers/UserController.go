/*
@Time : 2020/5/24 11:12
@Author : Firewine
@File : UserController
@Software: GoLand
@Description:
*/
package controllers

import (
	"RBAC_GO/src/service"
	"github.com/kataras/iris"
)

type UserController struct {
	//每个请求都由Iris自动绑定上下文，
	//记住，每次传入请求时，iris每次都会创建一个新的UserController，
	//所以所有字段都是默认的请求范围，只能设置依赖注入
	//自定义字段，如服务，对所有请求都是相同的（静态绑定）
	//和依赖于当前上下文的会话（动态绑定）。
	Ctx iris.Context
	//我们的UserService，它是一个接口
	//从主应用程序绑定。
	Service service.UserService
}
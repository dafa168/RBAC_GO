/*
@Time : 2020/5/26 22:27
@Author : Firewine
@File : DispatcherController
@Software: GoLand
@Description:
*/
package controllers

import (
	"github.com/kataras/iris"
)

func Dispatcher(index iris.Party)  {
	index.Get("/login",login)
}

func login(ctx iris.Context){
	ctx.View("login.html")
}
func error(ctx iris.Context){
	ctx.View("error.html")
}
func logout(ctx iris.Context){

	ctx.View("login.html")
}
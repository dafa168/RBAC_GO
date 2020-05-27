/*
@Time : 2020/5/26 22:27
@Author : Firewine
@File : DispatcherController
@Software: GoLand
@Description:
*/
package controllers

import (
	"RBAC_GO/src/dao"
	"RBAC_GO/src/models"
	"fmt"
	"github.com/kataras/iris"
)

func DoAJAXLogin(ctx iris.Context) {
	var (
		loginAcct = ctx.FormValue("loginAcct")
		password = ctx.FormValue("userPsWd")
	)
	loginers:= models.User{
		LoginAcct: loginAcct,
		UserPsWd: password,
	}
	login, err := dao.QueryLogin(&loginers)
	if err != nil{
		ctx.JSON(models.AJAXResult{
			Success: false,
			Msg: string(iris.StatusNonAuthoritativeInfo),
		})
	}
	fmt.Println(login)
}

func Logout(ctx iris.Context){


}
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
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
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
			Msg: "error",
		})
		return
	}
	if login.Id == 0 {
		ctx.JSON( models.AJAXResult{
			Success: false,
		})
		return
	} else {
		// 账号密码正确，得到用户
		token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			"iat": time.Now().Unix(),
		})
		tokenString, _ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))

		oauthToken := new(dao.OauthToken)
		oauthToken.Token = tokenString
		oauthToken.UserId = login.Id
		oauthToken.Secret = "secret"
		oauthToken.Revoked = false
		oauthToken.ExpressIn = time.Now().Add(time.Hour * time.Duration(1)).Unix()

		response := oauthToken.OauthTokenCreate()

		ctx.JSON(models.AJAXResult{
			Success: true,
			Data: response,
		})
		return
	}
}

func Logout(ctx iris.Context){
	userId, _ := strconv.Atoi(ctx.FormValue("userId"))
	ot := dao.OauthToken{}
	ot.UpdateOauthTokenByUserId(userId)

	ctx.JSON(models.AJAXResult{
		Success: true,
		Msg: "user exit",
	})

}

func Main(ctx iris.Context){
	_ = ctx.View("main.html")
}
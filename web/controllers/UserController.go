/*
@Time : 2020/5/24 11:12
@Author : Firewine
@File : UserController
@Software: GoLand
@Description:
*/
package controllers

import (
	"RBAC_GO/src/dao"
	"RBAC_GO/src/models"
	"github.com/kataras/iris/v12"
	"strconv"
)
// 批量删除users
func Deletes(ctx iris.Context)  {
	userId := ctx.FormValue("userid")
	maps := make(map[string]string,0)
	maps["userids"] = userId

	rows := dao.DeleteUsers(maps)
	if rows >0 {
		ctx.JSON(models.AJAXResult{
			Success: true,
		})
	}else {
		ctx.JSON(models.AJAXResult{
			Success: false,
		})
	}
}
// 通过id 删除单个user
func Delete(ctx iris.Context)  {
	id := ctx.FormValue("id")

	rows := dao.DeleteUserById(id)
	if rows >0 {
		ctx.JSON(models.AJAXResult{
			Success: true,
		})
	}else {
		ctx.JSON(models.AJAXResult{
			Success: false,
		})
	}
}
// 更新user 信息
func Update(ctx iris.Context)  {
	user := models.User{}
	_, _ = ctx.JSON(&user)
	updateUser := dao.UpdateUser(&user)
	if updateUser >0 {
		ctx.JSON(models.AJAXResult{
			Success: true,
		})
	}else {
		ctx.JSON(models.AJAXResult{
			Success: false,
		})
	}
}
func Edit(ctx iris.Context)  {
	ctx.FormValue("")
}
func Assign(ctx iris.Context)  {
	ctx.FormValue("")
}
func DoAssign(ctx iris.Context)  {
	ctx.FormValue("")
}
func DounAssign(ctx iris.Context)  {
	ctx.FormValue("")
}
func Insert(ctx iris.Context)  {
	ctx.FormValue("")
}
func Add(ctx iris.Context)  {

}
func Index(ctx iris.Context)  {

}
func Index1(ctx iris.Context)  {
	ctx.FormValue("")
}

func PageQuery(ctx iris.Context){

	queryText  := ctx.FormValue("queryText")
	pageno:= ctx.FormValue("pageno")
	pagesize := ctx.FormValue("pagesize")

	maps := make(map[string]interface{},0)

	pageno1, _ := strconv.ParseInt(pageno, 10, 64)
	pagesize1, _ := strconv.ParseInt(pagesize, 10, 64)
	maps["start"]  = (pageno1-1) * pagesize1
	maps["size"] = pagesize
	maps["queryText"] = queryText
	data, err := dao.PageQueryUserData(maps)
	if err != nil{
		ctx.JSON(models.AJAXResult{
			Success: false,
			Msg: "出错",
		})
		return
	}
	totalsize,err := dao.PageQueryUserCount(maps)
	if err != nil{
		ctx.JSON(models.AJAXResult{
			Success: false,
			Msg: "出错",
		})
		return
	}
	var totalno int64
	if totalsize % pagesize1 == 0{
		totalno = totalsize / pagesize1
	}else{
		totalno  = totalsize / pagesize1 +1
	}
	userPage := models.Page{
		Totalno: totalno,
		TotalSize: pagesize1,
	}
	userPage.SetDatas(data)
	ctx.JSON(models.AJAXResult{
		Data: userPage,
		Success: true,
		Msg: "分页获取user",
	})
}
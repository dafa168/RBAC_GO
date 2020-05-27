package middleware

import (
	"RBAC_GO/configs"
	"fmt"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/kataras/iris"
)


func InitCasbin(app *iris.Application) {

	engine := xormadapter.NewAdapterByEngine(configs.Engine)
	Enforcer, err := casbin.NewEnforcer("./rbac_model.conf", engine)
	if err != nil{
		fmt.Println(err)
		return
	}
	Enforcer.LoadPolicy()






}

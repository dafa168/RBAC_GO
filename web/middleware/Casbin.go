/*
@Time : 2020/5/26 22:05
@Author : Firewine
@File : Casbin
@Software: GoLand
@Description:
*/
package middleware

import (
	"RBAC_GO/configs"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	//sub   "alice"// 想要访问资源的用户.
	//obj  "data1" // 要访问的资源.
	//act  "read"  // 用户对资源执行的操作.
	engine := xormadapter.NewAdapterByEngine(configs.Engine)

	Enforcer, _ := casbin.NewEnforcer("rbac_model.conf", engine)

}
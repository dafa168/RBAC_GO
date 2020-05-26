package middleware

import (
	"RBAC_GO/configs"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {

	engine := xormadapter.NewAdapterByEngine(configs.Engine)

	Enforcer, _ := casbin.NewEnforcer("rbac_model.conf", engine)

}

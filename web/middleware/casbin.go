package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/kataras/iris"
	"net/http"
)
var Enforcer *casbin.Enforcer

func InitCasbin() {
	//sub   "alice"// 想要访问资源的用户.
	//obj  "data1" // 要访问的资源.
	//act  "read"  // 用户对资源执行的操作.
	//engine := xormadapter.NewAdapterByEngine(configs.Engine)

	Enforcer, _ := casbin.NewEnforcer("E:\\GO\\RBAC_GO\\rbac_model.conf", "E:\\GO\\RBAC_GO\\rbac_policy.csv")

	_ = Enforcer.LoadPolicy()

}
func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}

func (c *Casbin) ServeHTTP(ctx iris.Context) {
	// 账户信息验证

	fmt.Println(ctx.Request().Method)

	ctx.Next()
}

// Casbin is the auth services which contains the casbin enforcer.
type Casbin struct {
	enforcer *casbin.Enforcer
}

// Check checks the username, request's method and path and
// returns true if permission grandted otherwise false.
func (c *Casbin) Check(r *http.Request, userId string) bool {
	method := r.Method
	path := r.URL.Path
	ok, _ := c.enforcer.Enforce(userId, path, method)
	return ok
}

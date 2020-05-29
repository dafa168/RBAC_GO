/*
@Time : 2020/5/26 22:05
@Author : Firewine
@File : Casbin
@Software: GoLand
@Description:
*/
package middleware

import (
	"RBAC_GO/src/dao"
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
	"net/http"
	"strconv"
	"time"
)



func InitCasbin() *casbin.Enforcer {
	//sub   "alice"// 想要访问资源的用户.
	//obj  "data1" // 要访问的资源.
	//act  "read"  // 用户对资源执行的操作.
	//adapter := xormadapter.NewAdapterByEngine(configs.Engine) // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	a := xormadapter.NewAdapter("mysql", "root:root@(192.168.20.235:3306)/")
	// Load the policy from DB.

	e, _ := casbin.NewEnforcer("./rbac_model.conf", a)
	e.LoadPolicy()


	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()

	return e
}

func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}
// 判断token
func (c *Casbin) ServeHTTP(ctx context.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	token := dao.OauthToken{}
	token.GetOauthTokenByToken(value.Raw) //获取 access_token 信息
	if token.Revoked || token.ExpressIn < time.Now().Unix() {
		//_, _ = ctx.Writef("token 失效，请重新登录") // 输出到前端
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.StopExecution()
		return
	} else if !c.Check(ctx.Request(), strconv.FormatUint(uint64(token.UserId), 10)) {
		ctx.StatusCode(http.StatusForbidden) // Status Forbidden
		ctx.StopExecution()
		return
	} else {
		ctx.Values().Set("auth_user_id", token.UserId)
	}

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
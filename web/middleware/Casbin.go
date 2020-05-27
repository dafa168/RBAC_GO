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
	"RBAC_GO/src/models"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/xorm-adapter"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"net/http"
	"strconv"
	"time"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	//sub   "alice"// 想要访问资源的用户.
	//obj  "data1" // 要访问的资源.
	//act  "read"  // 用户对资源执行的操作.
	engine := xormadapter.NewAdapterByEngine(configs.Engine)

	Enforcer, _ := casbin.NewEnforcer("rbac_model.conf", engine)

	_ = Enforcer.LoadPolicy()

}

func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}
// 判断token
func (c *Casbin) ServeHTTP(ctx iris.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	token := models.OauthToken{}
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
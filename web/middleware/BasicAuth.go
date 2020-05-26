package middleware

//import (
//	"github.com/kataras/iris"
//	"github.com/kataras/iris/middleware/basicauth"
//	"time"
//)

//func newApp() *iris.Application {
//
//	//authConfig := basicauth.Config{
//	//	Users:       map[string]string{"myusername": "mypassword", "mySecondusername": "mySecondpassword"},
//	//	Realm:       "Authorization Required", // 默认表示域 "Authorization Required"
//	//	Expires:     time.Duration(30) * time.Minute,
//	//
//	//}
//	//authentication := basicauth.New(authConfig)
//	//作用范围 全局 app.Use(authentication) 或者 (app.UseGlobal 在Run之前)
//	//作用范围 单个路由 app.Get("/mysecret", authentication, h)
//
//
//	////作用范围  Party
//	//needAuth := app.Party("/admin", authentication)
//	//{
//	//	//http://localhost:8080/admin
//	//	needAuth.Get("/", h)
//	//	// http://localhost:8080/admin/profile
//	//	needAuth.Get("/profile", h)
//	//	// http://localhost:8080/admin/settings
//	//	needAuth.Get("/settings", h)
//	//}
//}

package routers

import (
	"RBAC_GO/src/dao"
	"RBAC_GO/src/service"
	"RBAC_GO/web/controllers"
	"github.com/kataras/iris/mvc"

)

func Dispatcher(mvc *mvc.Application) {
	//添加基本身份验证（admin：password）中间件
	// 创建dao
	//创建服务
	userService := service.NewUserService(dao.NewUserDao(), dao.NewUserRoleDao())
	mvc.Register(userService)
	//为我们的电影控制器服务
	//请注意，您可以为多个控制器提供服务
	// 如果你想。
	mvc.Handle(new(controllers.DispatcherController))

}

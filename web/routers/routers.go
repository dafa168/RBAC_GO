package routers

import (
	"RBAC_GO/src/dao"
	"RBAC_GO/src/service"
	"RBAC_GO/web/controllers"
	"github.com/kataras/iris/mvc"
)
func TotalRouters(mvc *mvc.Application){
	Dispatcher(mvc)
}
func Dispatcher(mvc *mvc.Application) {
	//添加基本身份验证（admin：password）中间件
	// 创建dao
	//创建服务
	userService := service.NewUserService(dao.NewUserDao())
	roleService := service.NewRoleService(dao.NewRoleDao())
	permissionService := service.NewPermissionService(dao.NewPermissionDao())
	mvc.Register(userService)
	mvc.Register(roleService)
	mvc.Register(permissionService)
	//为我们的电影控制器服务
	//请注意，您可以为多个控制器提供服务
	mvc.Handle(new(controllers.DispatcherController))
	mvc.Handle(new(controllers.UserController))
}

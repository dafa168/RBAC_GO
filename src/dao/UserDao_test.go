package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"fmt"
	"testing"
)

func test(r *testing.M)  {
	configs.MysqlEngine()
	r.Run()
}
func Test_userDao_DeleteUsers1(t *testing.T) {
	dao := userDao{}
	cc := make(map[string]interface{}, 0)
	a := []int{1, 2, 3}
	cc["usersId"] = a
	users := dao.DeleteUsers(cc)
	fmt.Println(users)
}
func Test_userDao_Querylogin(t *testing.T) {
	dao := userDao{}
	login := models.User{
		UserPsWd: "admin",
		LoginAcct: "admin",
	}
	users, _ := dao.QueryLogin(&login)

	fmt.Println(users)
}

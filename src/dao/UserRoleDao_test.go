package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试userdao中的方法")
	configs.MysqlEngine()
	m.Run()
}



func TestQueryRoleByUserid(t *testing.T) {
	results := QueryRoleIdsByUserId(1)

	fmt.Printf("%v\n",results)


}

func TestInsertUserRoles(t *testing.T) {
	var lists []models.UserRole
	userRole := models.UserRole{
		RoleId: 1,
		UserId: 2,
	}
	userRole1 := models.UserRole{
		RoleId: 1,
		UserId: 2,
	}
	lists = append(lists, userRole)
	lists = append(lists, userRole1)
	roles := InsertUserRoles(lists)
	fmt.Println(roles)
}
/*
@Time : 2020/5/23 21:49
@Author : Firewine
@File : RolePermissionDao_test.go
@Software: GoLand
@Description:
*/
package dao

import (
	"RBAC_GO/src/models"
	"fmt"
	"testing"
)

func TestInsertRolePermission(t *testing.T) {

	lists := make([]models.RolePermission,0)
	a := models.RolePermission{
		RoleId: 1,
		PermissionId: 2,
	}
	a1 := models.RolePermission{
		RoleId: 1,
		PermissionId: 2,
	}
	lists = append(lists,a,a1)
	InsertRolePermission(lists)
}

func TestDeleteRolePermissions(t *testing.T) {

}

func TestQueryPermissionIdsByRoleId(t *testing.T) {
	id, _ := QueryPermissionIdsByRoleId(1)
	fmt.Println(id)
}

func TestQueryPermissionsByUser(t *testing.T) {
	userRole := models.UserRole{
		UserId: 1,
	}
	user, _ := QueryPermissionsByUser(userRole)
	fmt.Println(user)
}
/*
@Time : 2020/5/24 0:34
@Author : Firewine
@File : PermissionDao_test.go
@Software: GoLand
@Description:
*/
package dao

import (
	"fmt"
	"testing"
)


func TestQueryAllPermission(t *testing.T) {
	permission, _ := QueryAllPermission()
	fmt.Println(permission)
}

func TestQueryChildPermissions(t *testing.T) {
	permissions, _ := QueryChildPermissions(1)
	fmt.Println(permissions)
}


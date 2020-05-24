/*
@Time : 2020/5/23 20:35
@Author : Firewine
@File : RoleDao_test.go
@Software: GoLand
@Description:
*/
package dao

import (
	"fmt"
	"testing"
)

func TestPageQueryDataRole(t *testing.T) {
	maps := make(map[string]string)
	maps["start"] = "0"
	maps["size"] = "2"
	//maps["queryText"] = "11111"
	data, err := PageQueryRoleData(maps)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(data)
}

func TestPageQueryCountRole(t *testing.T) {
	maps := make(map[string]string)
	maps["start"] = "0"
	maps["size"] = "2"
	//maps["queryText"] = "11111"
	data, err := PageQueryRoleCount(maps)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(data)

}

func TestQueryAllRole(t *testing.T) {
	roles, err := QueryAllRole()
	fmt.Println(err)
	fmt.Println(roles)
}
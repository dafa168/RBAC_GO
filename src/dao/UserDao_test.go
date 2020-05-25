package dao

import (
	"fmt"
	"testing"
)




func Test_userDao_DeleteUsers1(t *testing.T) {
	dao := userDao{}
	cc := make(map[string]interface{},0)
	a := []int{1,2,3}
	cc["usersId"] = a
	users := dao.DeleteUsers(cc)
	fmt.Println(users)
}
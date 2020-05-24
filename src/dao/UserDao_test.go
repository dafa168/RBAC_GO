package dao

import (
	"RBAC_GO/src/models"
	"fmt"
	"testing"
	"time"
)




func TestQueryAll(t *testing.T) {
	results, _ := QueryAll()

	fmt.Println(results)


}


func TestUpdateUser(t *testing.T) {
	user := models.User{
		Id:       1,
		UserPsWd: "admin",
		LoginAcct: "admin",
		Username: "lizi",
		Email: "123@qq.com",
	}
   results := UpdateUser(&user)
   fmt.Println(results)
}

func TestDeleteUsers(t *testing.T) {
	users := make(map[string][]string,0)
	users["usersid"] = []string{"1","2"}
	deleteUsers := DeleteUsers(users)
	fmt.Println(deleteUsers)
}

func TestInsertUser(t *testing.T) {
	user  := models.User{
		Username:   "lll",
		LoginAcct:  "admin",
		UserPsWd:   "admin",
		Email:      "111@111.com",
		CreateTime: time.Now().Unix(),

	}

	fmt.Println(user)
	insertUser, err := InsertUser(&user)
	fmt.Println(user)
	if err != nil{
		fmt.Println("cuwou :",err)
	}
	fmt.Println(insertUser)
}
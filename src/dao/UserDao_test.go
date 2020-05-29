package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"fmt"
	"testing"
	"time"
)


func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	configs.MysqlEngine()
	m.Run()
	// 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作

}

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
	users := make(map[string][]int,0)
	users["userids"] = []int{1,2}
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

func TestQueryLogin(t *testing.T) {
	user  := models.User{
		LoginAcct:  "root",
		UserPsWd:   "admin",
	}
	login, _ := QueryLogin(&user)
	fmt.Println(login)
}

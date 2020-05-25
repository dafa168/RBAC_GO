package dao

import (
	"RBAC_GO/configs"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试userdao中的方法")
	configs.MysqlEngine()
	m.Run()
}
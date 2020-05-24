package configs

import (
	"RBAC_GO/src/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func MysqlEngine() {
	// 读取配置文件
	myEnv := loadEnvText()
	var err error

	Engine, err = xorm.NewEngine(myEnv["DBName"], myEnv["CONNECT_URL"])
	if err != nil {
		fmt.Printf("错误信息：%e", err)
	}
	Engine.ShowSQL(true)
	Engine.Logger()
	err = Engine.Sync2(new(models.User), new(models.Role), new(models.Permission), new(models.RolePermission), new(models.UserRole))
	if err != nil {
		fmt.Printf("同步结构错误：%v", err)
	}
}
// 项目从main方法启动，需要在db上面改动
func loadEnv() map[string]string {
	var myEnv map[string]string

	myEnv, err := godotenv.Read("test_db.env")
	if err != nil{
		fmt.Println(err)
	}
	return myEnv
}
// 后台的方法测试，使用 的绝对路径方法
func loadEnvText() map[string]string {
	var myEnv map[string]string

	myEnv, err := godotenv.Read("E:\\GO\\RBAC_GO\\test_db.env")
	if err != nil{
		fmt.Println(err)
	}

	return myEnv
}
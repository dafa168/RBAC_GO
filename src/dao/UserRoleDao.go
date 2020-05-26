package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"bytes"
	"strconv"
)
type UserRoleDao interface {
	QueryRoleIdsByUserId(id int) interface{}
	InsertUserRoles(list []models.UserRole) error
	DeleteUserRoles(list []models.UserRole) error
}
// sql采用工具里， 这里不需要在结构体中重新定义
type userRoleDao struct {}

func NewUserRoleDao() *userRoleDao {
	return &userRoleDao{}
}
func (Dao *userRoleDao)QueryRoleIdsByUserId(id int) interface{} {
	sqlStr := "select role_id from user_role where user_id = ?"
	results, err := configs.Engine.QueryInterface(sqlStr, id)
	if err != nil {
		return err
	}
	userId := results[0]["role_id"]
	return userId
}
// 向关系表，添加关联数据
func (Dao *userRoleDao)InsertUserRoles(list []models.UserRole) error {
	sqlStr := "insert into user_role ( user_id, role_id ) values ("
	var bt bytes.Buffer
	//向bt中写入字符串
	bt.WriteString(sqlStr)
	for num,userRole := range list{
		// v是models.userrole，这就可以拼接sql
		bt.WriteString(strconv.Itoa(userRole.UserId))
		bt.WriteString(",")
		bt.WriteString(strconv.Itoa(userRole.RoleId))
		if (num) == len(list)-1{
			bt.WriteString(")")
		}else {
			bt.WriteString("),(")
		}
	}
	//fmt.Println(bt.String())
	//	下面进行遍历，得到id，进行添加
	_, err := configs.Engine.Exec(bt.String())
	if err != nil{
		return err
	}
	return nil

}
// 删除关系表数据
func (Dao *userRoleDao)DeleteUserRolesV1(list []models.UserRole) error {
	sqlStr := "delete from user_role where user_id = ? and role_id in ("
	var bt bytes.Buffer
	//向bt中写入字符串
	bt.WriteString(sqlStr)
	for _,userRole := range list{
		_, err := configs.Engine.Where("user_id=", userRole.UserId).Where("role_id=", userRole.RoleId).Delete(new(models.UserRole))
		if err != nil{
			return err
		}
	}
	return nil
}
// 删除关系表数据 单用秃，user_id 是不变的，所以这里UserROle 结构体，里面
// user_id 是一样的
func (Dao *userRoleDao)DeleteUserRoles(list []models.UserRole) error{
	sqlStr := "delete from user_role where user_id = ? and role_id in ("
	var bt bytes.Buffer
	//向bt中写入字符串
	bt.WriteString(sqlStr)
	var userId int
	for _,userRole := range list{
		userId = userRole.UserId
		bt.WriteString(strconv.Itoa(userRole.RoleId))
		bt.WriteString(",")
	}
	bt.WriteString(")")
	_, err := configs.Engine.Exec(bt.String(), userId)
	if err != nil{
		return err
	}
	return nil
}

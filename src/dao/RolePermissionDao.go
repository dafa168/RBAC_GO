/*
@Time : 2020/5/23 21:48
@Author : Firewine
@File : RolePermissionDao
@Software: GoLand
@Description:
*/
package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"bytes"
	"strconv"
)

// 插入role-permission 权限信息
func InsertRolePermission(list []models.RolePermission) (int64, error) {
	//"roleid", roleid
	//"permissionids", permissionids
	sqlStr := "insert into role_permission (role_id, permission_id) values ("
	var bt bytes.Buffer
	//向bt中写入字符串
	bt.WriteString(sqlStr)
	for num, rolePermission := range list {
		// v是models.userrole，这就可以拼接sql
		bt.WriteString(strconv.Itoa(rolePermission.RoleId))
		bt.WriteString(",")
		bt.WriteString(strconv.Itoa(rolePermission.PermissionId))
		if (num) == len(list)-1 {
			bt.WriteString(")")
		} else {
			bt.WriteString("),(")
		}
	}
	//fmt.Println(bt.String())
	exec, err := configs.Engine.Exec(bt.String())
	if err != nil{
		return -1, err
	}
	affected, _ := exec.RowsAffected()
	return affected,nil
}
// 删除role_permission
func DeleteRolePermissions(list []models.RolePermission)(int64,error){
	sqlStr := "delete from role_permission where role_id = ?"

	exec, err := configs.Engine.Exec(sqlStr,list[0].RoleId)
	if err != nil{
		return -1, err
	}
	affected, _ := exec.RowsAffected()
	return affected,nil
}

// 根据角色id，查找权限id
func QueryPermissionIdsByRoleId(roleId int) ([]int,error){
	sqlStr := "select permission_id from role_permission where role_id = ?"
	permissionIds := make([]int,0)
	err := configs.Engine.SQL(sqlStr, roleId).Find(&permissionIds)
	if err != nil{
		return nil, err
	}
	return permissionIds,nil
}
// 根据角色，查找权限
func QueryPermissionsByUser(userRole models.UserRole)([]models.Permission,error){

	sqlStr := "  select * from permission where id in ( select permission_id from role_permission  where role_id in ( select  role_id    from user_role  where user_id = ?))"
	rows, err := configs.Engine.SQL(sqlStr, userRole.UserId).Rows(new(models.Permission))
	if err != nil{
		return nil, err
	}
	permissions := make([]models.Permission,0)
	defer rows.Close()
	for rows.Next(){
		permission := models.Permission{}
		rows.Scan(&permission)
		permissions = append(permissions,permission)

	}
	return permissions,nil
}
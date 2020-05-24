package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
)
// 查询root 权限
func QueryRootPermission()(models.Permission,error){
	sqlStr := "select * from permission where pid is null"
	permission := models.Permission{}
	rows, err := configs.Engine.SQL(sqlStr).Rows(new(models.RolePermission))
	if err != nil{
		return permission,err
	}
	defer  rows.Close()
	for rows.Next(){

		err = rows.Scan(&permission)

	}
	return permission,nil
}

// 查询选择的权限
func QueryChildPermissions(pid int) ([]models.Permission,error) {
	sqlStr := "select * from permission where pid = ?"
	permissions  := make([]models.Permission,0)
	rows, err := configs.Engine.SQL(sqlStr,pid).Rows(new(models.Permission))
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		permission := models.Permission{}
		rows.Scan(&permission)
		permissions = append(permissions,permission)
	}
	return permissions,nil
}

// 查询所有权限
func QueryAllPermission()([]models.Permission,error){
	sqlStr := "select * from permission "
	permissions  := make([]models.Permission,0)
	rows, err := configs.Engine.SQL(sqlStr).Rows(new(models.Permission))
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		permission := models.Permission{}
		rows.Scan(&permission)
		permissions = append(permissions,permission)
	}
	return permissions,nil
}

// 插入权限
func InsertPermission(permission models.Permission)(int64,error){
	sqlStr := "insert into t_permission (name, url, pid) values (?,?,?)"
	exec, err := configs.Engine.Exec(sqlStr, permission.Name, permission.Url, permission.Pid)
	if err != nil{
		return -1 ,err
	}
	rowsAffected, _ := exec.RowsAffected()
	return rowsAffected,nil
}

//根据id查找权限
func QueryPermissionById(id int)(models.Permission,error){
	sqlStr := "select * from t_permission where id =  ?"
	permission  := models.Permission{}
	rows, err := configs.Engine.SQL(sqlStr,id).Rows(new(models.Permission))
	if err != nil{
		return permission, err
	}
	defer rows.Close()
	for rows.Next(){
		permission := models.Permission{}
		rows.Scan(&permission)

	}
	return permission,nil
}

// 更新权限
func UpdatePermission(permission models.Permission)(int64,error){
	sqlStr :=  "update permission set name = ?, url = ?  where id = ?"
	exec, err := configs.Engine.Exec(sqlStr, permission.Name, permission.Url, permission.Id)
	if err != nil {
		return -1,err
	}
	affected, _ := exec.RowsAffected()
	return affected,nil
}

//删除权限
func DeletePermission(permission models.Permission)(int64,error){

	sqlStr := "delete from permission where id = ?"
	exec, err := configs.Engine.Exec(sqlStr, permission.Id)
	if err != nil {
		return -1,err
	}
	affected, _ := exec.RowsAffected()
	return affected,nil
}


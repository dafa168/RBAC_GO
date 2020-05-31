package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
)
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

// 根据角色查找权限
func QueryRolePermission(){

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


package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
)

// 查询选择的权限
func QueryChildPermissions(pid int) ([]models.CasbinRule,error) {
	sqlStr := "select * from permission where pid = ?"
	permissions  := make([]models.CasbinRule,0)
	rows, err := configs.Engine.SQL(sqlStr,pid).Rows(new(models.CasbinRule))
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		permission := models.CasbinRule{}
		rows.Scan(&permission)
		permissions = append(permissions,permission)
	}
	return permissions,nil
}


// 查询所有权限
func QueryAllPermission()([]models.CasbinRule,error){
	sqlStr := "select * from casbin_rule where p_type = 'p'"
	permissions  := make([]models.CasbinRule,0)
	rows, err := configs.Engine.SQL(sqlStr).Rows(new(models.CasbinRule))
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		permission := models.CasbinRule{}
		rows.Scan(&permission)
		permissions = append(permissions,permission)
	}
	return permissions,nil
}

// 根据角色查找权限
func QueryRolePermission(){

}
// 插入权限
func InsertPermission(permission models.CasbinRule)(int64,error){
	sqlStr := "insert into casbin_rule (p_type, v0, v1,v2) values (p,?,?,?)"
	exec, err := configs.Engine.Exec(sqlStr,permission.V0,permission.V1,permission.V2)
	if err != nil{
		return -1 ,err
	}
	rowsAffected, _ := exec.RowsAffected()
	return rowsAffected,nil
}


//查找角色下管理的权限
func QueryPermissionByRole(role string)([]models.CasbinRule,error){
	sqlStr := "select * from casbin_rule where p_type = 'p' and v0 = ?"
	rows, err := configs.Engine.SQL(sqlStr,role).Rows(new(models.CasbinRule))
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	permissions := make([]models.CasbinRule,0)
	for rows.Next(){
		permission := models.CasbinRule{}
		rows.Scan(&permission)
		permissions = append(permissions, permission)

	}
	return permissions,nil
}



// 更新权限
func UpdatePermission(permission models.CasbinRule)(int64,error){
	sqlStr :=  "update casbin_rule set v1 = ?,v2 = ?  where p_type= ? AND v0 = ?"
	exec, err := configs.Engine.Exec(sqlStr, permission.V1,permission.V2,permission.PType,permission.V0)
	if err != nil {
		return -1,err
	}
	affected, _ := exec.RowsAffected()
	return affected,nil
}

//删除权限
func DeletePermission(permission models.CasbinRule)(int64,error){

	sqlStr := "delete from casbin_rule where p_type = 'p' and v0 = ? and  v1= ? and v2 = ?"
	exec, err := configs.Engine.Exec(sqlStr, permission.PType,permission.V0,permission.V1,permission.V2)
	if err != nil {
		return -1,err
	}
	affected, _ := exec.RowsAffected()
	return affected,nil
}


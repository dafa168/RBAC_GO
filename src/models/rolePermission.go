package models

// role_permission 结构
type RolePermission struct {
	Id           int `xorm:" pk  int notnull autoincr unique"`
	RoleId       int `xorm:"int "`
	PermissionId int `xorm:"int "`
}

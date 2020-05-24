package models

type UserRole struct {
	Id int `xorm:" pk  int notnull autoincr unique"`
	UserId int `xorm:"int"`
	RoleId int`xorm:"int "`
}

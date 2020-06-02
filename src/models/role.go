package models

type Role struct {
	Id   int    `xorm:" pk  int notnull autoincr unique"`
	Name string `xorm:"varchar(255)  "`
	Ename  string `xorm:"varchar(255)"`
}

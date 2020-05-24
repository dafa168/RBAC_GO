package models

type Permission struct {
	Id int `xorm:" pk  int notnull autoincr unique"`
	Name string `xorm:"varchar(255) "`
	Pid int`xorm:"int"`
	Url string`xorm:"varchar(255) "`
	Icon string`xorm:"varchar(255) "`

}

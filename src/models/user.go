package models

type User struct {
	Id         int    `xorm:" pk  int notnull autoincr unique"`
	Username   string `xorm:"varchar(255) "`
	LoginAcct  string `xorm:"varchar(255)  "`
	UserPsWd   string `xorm:"varchar(255)  "`
	Email      string `xorm:"varchar(255)  "`
	CreateTime int64 `xorm:"timestamp "`
}

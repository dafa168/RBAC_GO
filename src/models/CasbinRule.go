package models


type CasbinRule struct {
	PType string `xorm:"varchar(100) "`
	V0    string `xorm:"varchar(100) "`
	V1    string `xorm:"varchar(100) "`
	V2    string `xorm:"varchar(100) "`
	V3    string `xorm:"varchar(100) "`
	V4    string `xorm:"varchar(100) "`
	V5    string `xorm:"varchar(100) "`
}

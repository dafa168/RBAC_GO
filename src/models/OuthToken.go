/*
@Time : 2020/5/27 22:00
@Author : Firewine
@File : OuthToken
@Software: GoLand
@Description:
*/
package models

type OauthToken struct {
	Token     string `xorm:"not null default '' comment('Token') VARCHAR(191)"`
	UserId    uint   `xorm:"not null default '' comment('UserId') VARCHAR(191)"`
	Secret    string `xorm:"not null default '' comment('Secret') VARCHAR(191)"`
	ExpressIn int64  `xorm:"not null default 0 comment('是否是标准库') BIGINT(20)"`
	Revoked   bool
}

type Token struct {
	Token string `json:"access_token"`
}


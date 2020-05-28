/*
@Time : 2020/5/27 23:58
@Author : Firewine
@File : OuthTokenDao
@Software: GoLand
@Description:
*/
package dao

import (
	"RBAC_GO/configs"
	"RBAC_GO/src/models"
	"strconv"
)

type OauthToken struct {
	models.OauthToken
}
type Token struct {
	Token string `json:"access_token"`
}
func  (ot *OauthToken)OauthTokenCreate() *Token {
	sqlStr := "INSERT INTO oauth_token ( token, user_id, secret, express_in, revoked) VALUES (?,?,?,?,?)"
	configs.Engine.Exec(sqlStr,&ot.Token,&ot.UserId,&ot.Secret,&ot.ExpressIn,&ot.Revoked)

	return &Token{ot.Token}
}

/**
* 通过 token 获取 access_token 记录
* @method GetOauthTokenByToken
* @param  {[type]}       token string [description]
 */
func  (ot *OauthToken)GetOauthTokenByToken(token string) {
	sqlStr := "select * from oauth_token where token = ?"
	get, _ := configs.Engine.QueryString(sqlStr,token)

	ot.UserId,_ = strconv.Atoi(get[0]["user_id"])
	ot.Token = get[0]["token"]
	ot.Secret = get[0]["secret"]
	ot.ExpressIn,_ = strconv.ParseInt(get[0]["express_in"],10,64)
	if get[0]["revoked"] == "0"{
		ot.Revoked = false
	}else {
		ot.Revoked = true
	}

	//fmt.Println(get[0]["token"])
	//fmt.Println(get[0]["express_in"])
	//fmt.Println(get[0]["secret"])
	//fmt.Println(get[0]["revoked"])
	//fmt.Println(get[0]["user_id"])

	//fmt.Println(get)
	//fmt.Println(ot)

}

//通过 user_id 更新 oauth_token 记录
//@method UpdateOauthTokenByUserId
func  (ot *OauthToken)UpdateOauthTokenByUserId(userId int) {
	sqlStr := "UPDATE `oauth_token` SET `revoked` = true WHERE revoked = false AND user_id = ?"
	_, _ = configs.Engine.Update(sqlStr,userId)
}



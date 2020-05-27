/*
@Time : 2020/5/27 23:58
@Author : Firewine
@File : OuthTokenDao
@Software: GoLand
@Description:
*/
package dao

/**
 * oauth_token
 * @method OauthTokenCreate
 */
//func  OauthTokenCreate() *Token {
//	configs.Engine.Insert(&ot)
//	return &Token{ot.Token}
//}
//
///**
//* 通过 token 获取 access_token 记录
//* @method GetOauthTokenByToken
//* @param  {[type]}       token string [description]
// */
//func  GetOauthTokenByToken(token string) {
//	configs.Engine.Where("token =  ?", token).Get(&ot)
//}
//
////通过 user_id 更新 oauth_token 记录
////@method UpdateOauthTokenByUserId
//func  UpdateOauthTokenByUserId(userId uint) {
//	configs.Engine.Table(ot).Where("revoked = ?", false).
//		Where("user_id = ?", userId).Update(map[string]interface{}{"revoked": true})
//}

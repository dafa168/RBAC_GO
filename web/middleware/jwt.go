/*
@Time : 2020/5/27 21:41
@Author : Firewine
@File : jwt
@Software: GoLand
@Description:
*/
package middleware

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

//func JwtHandler() *jwt.Middleware {
//	var mySecret = []byte("HS2JDFKhu7Y1av7b")
//	return jwt.New(jwt.Config{
//		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//			return mySecret, nil
//		},
//		Extractor: jwt.FromParameter("token"),
//		SigningMethod: jwt.SigningMethodHS256,
//	})
//}

func JWTMiddleware() *jwtmiddleware.Middleware {
	var mySecret = []byte("HS2JDFKhu7Y1av7b")
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return mySecret, nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)

		//debug 模式
		//Debug: bool
	})
	return jwtHandler
}
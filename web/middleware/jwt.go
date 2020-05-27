package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
)

// 验证jwt
func JwtHandler() *jwt.Middleware {
	var mySecret = []byte("yosemite11111")
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}


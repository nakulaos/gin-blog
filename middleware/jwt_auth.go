package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
)

/*
	这样写也可以，就是一个包含 *gin.Context的函数

	func JwtAuth(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			res.FailWithMessage("请输入token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
*/
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			res.FailWithMessage("请输入token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

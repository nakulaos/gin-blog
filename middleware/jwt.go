package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
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
		// 判断用户注销的token在不在redis里面
		/* 也可以这样写，感觉写法不少的
		count, err := global.Redis.Exists(context.Background(), fmt.Sprintf("logout_%s", token)).Result()
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		if count == 1 {
			res.OkWithMessage("用户已注销，token已失效", c)
			c.Abort()
			return
		}
		*/

		result := global.Redis.Exists(context.Background(), fmt.Sprintf("logout_%s", token))
		if result.Err() != nil {
			res.FailWithMessage(result.Err().Error(), c)
			c.Abort()
			return
		}
		if result.Val() == 1 {
			res.OkWithMessage("用户已注销，token已失效", c)
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

func JwtAdmin() gin.HandlerFunc {
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
		c.Set("role", claims.Role)
	}
}

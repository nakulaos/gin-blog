package user_api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
	"time"
)

func (userApi UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	my_time := time.Unix(claims.ExpiresAt, 0)
	ExpireTime := my_time.Sub(time.Now())
	err := global.Redis.Set(context.Background(), fmt.Sprintf("logout_%s", c.GetHeader("token")), "dzc", ExpireTime).Err()
	if err != nil {
		res.FailWithMessage("注销失败，请稍后再试", c)
		return
	}
	res.OkWithMessage("注销成功", c)

}

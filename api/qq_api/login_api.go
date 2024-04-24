package qq_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/utils"
	"gvb_server/utils/jwt"
)

func (qqLoginApi QQLoginApi) QQLoginApiView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	fmt.Println(code)
	qqinfo, err := NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		//	fmt.Println(err)
		return
	}
	fmt.Println(qqinfo)
	res.OkWithData(qqinfo, c)
	openID := qqinfo.OpenID
	var usr models.UserModel
	// 根据openID判断用户是否存在
	err = global.DB.Take(&usr, "token = ?", openID).Error

	if err == nil { //用户已存在,只给个jwt
		token, err := jwt.GenToken(jwt.JwtPayLoad{
			Nickname: usr.NickName,
			Role:     int(usr.Role),
			UserID:   usr.ID,
		})
		if err != nil {
			global.Log.Error(err)
			return
		}
		res.OkWithData(token, c)
		return

	}
	user := models.UserModel{
		NickName:   qqinfo.Nickname,
		UserName:   qqinfo.OpenID,
		Password:   utils.RandStr(16), //随机生成
		Avatar:     qqinfo.FigureurlQQ,
		Addr:       "内网", //可以使用第三方包根据ip来算地理位置
		Token:      qqinfo.OpenID,
		IP:         c.ClientIP(),
		Role:       ctype.PermissionUser,
		SignStatus: ctype.SignQQ,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		res.FailWithMessage("使用QQ注册失败", c)
		return
	}
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		Nickname: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithData(token, c)

}

package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
)

type LoginEmailRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (userApi UserApi) UserLoginEmailView(c *gin.Context) {
	var cr LoginEmailRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	var userLogin models.UserModel
	if err := global.DB.Take(&userLogin, "user_name = ?", cr.UserName).Error; err != nil {
		global.Log.Error("该用户不存在！")
		res.FailWithMessage("用户名或密码错误！", c)
		return
	}
	if !pwd.ValidPassword(userLogin.Password, cr.Password) {
		res.FailWithMessage("用户名或密码错误！", c)
		return
	}

	//登陆成功，分发token
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		Nickname: userLogin.NickName,
		Role:     int(userLogin.Role),
		UserID:   userLogin.ID,
	})
	if err != nil {
		res.FailWithMessage("登陆失败！", c)
		return
	}
	res.OkWithData(token, c)
}

package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/utils/pwd"
)

type UserCreateRequest struct {
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"`
	NickName string     `json:"nick_name" binding:"required" msg:"请输入用户昵称"`
	Password string     `json:"password" binding:"required" msg:"请输入用户密码"`
	Role     ctype.Role `json:"role" binding:"required" msg:"请输入用户权限"`
}

func (userApi UserApi) UserCreateView(c *gin.Context) {
	_role, _ := c.Get("role")
	role := _role.(int)
	if role != 1 {
		res.FailWithMessage("您不是管理员，权限受限!", c)
		return
	}
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}

	if count := global.DB.Take(&models.UserModel{}, "user_name = ?", cr.UserName).RowsAffected; count != 0 {
		global.Log.Error("该用户名已存在！")
		return
	}
	userModel := models.UserModel{
		UserName:   cr.UserName,
		NickName:   cr.NickName,
		Password:   pwd.HashPassword(cr.Password),
		Addr:       "内网",
		Avatar:     "/uploads/avatar/default.jpg",
		IP:         "127.0.0.1",
		Role:       cr.Role,
		SignStatus: ctype.SignEmail,
	}
	if err := global.DB.Create(&userModel).Error; err != nil {
		global.Log.Error("用户创建失败!")
		return
	}
	res.OkWithMessage(fmt.Sprintf("用户 %s 创建成功！", cr.UserName), c)
	global.Log.Info(fmt.Sprintf("用户 %s 创建成功！", cr.UserName))

}

package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
)

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" msg:"请输入旧密码！"`
	Password    string `json:"password" binding:"required" msg:"请输入新密码！"`
}

func (userApi UserApi) UserUpdatePasswordView(c *gin.Context) {
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var user models.UserModel
	if err := global.DB.Take(&user, claims.UserID).Error; err != nil {
		res.FailWithMessage("没有此用户！", c)
		return
	}
	//验证旧密码
	if !pwd.ValidPassword(user.Password, cr.OldPassword) {
		res.FailWithMessage("旧密码不正确！", c)
		return
	}
	//更新新密码
	hash := pwd.HashPassword(cr.Password)
	if err := global.DB.Model(&user).Update("password", hash).Error; err != nil {
		res.FailWithMessage("修改密码失败！", c)
		return
	}
	res.OkWithMessage("修改密码成功！", c)

}

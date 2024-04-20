package user_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type UpdateRequest struct {
	ID       uint       `json:"id" binding:"required" msg:"请输入id"`
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"`
}

func (userApi UserApi) UserUpdateUserinfoView(c *gin.Context) {
	_role, _ := c.Get("role")
	role := _role.(int)
	if role != 1 {
		res.FailWithMessage("您不是管理员，权限受限!", c)
		return
	}
	var cr UpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	var user models.UserModel
	if err := global.DB.Take(&user, "id = ?", cr.ID).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	var maps map[string]any
	if cr.NickName != "" {
		maps = structs.Map(&cr)
	} else {
		maps = structs.Map(&UpdateRequest{
			ID:       cr.ID,
			Role:     cr.Role,
			NickName: user.NickName,
		})
	}
	if err := global.DB.Model(&user).Updates(maps).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("修改用户信息成功!", c)

}

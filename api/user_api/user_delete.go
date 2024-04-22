package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (userApi UserApi) UserDeleteView(c *gin.Context) {
	_role, _ := c.Get("role")
	role := _role.(int)
	if role != 1 {
		res.FailWithMessage("您不是管理员，不能删除用户！", c)
		return
	}
	var cr models.RemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	var users []models.UserModel
	global.DB.Model(&models.UserModel{}).Find(&users, cr.IDList)

	if err := global.DB.Delete(&users).Error; err != nil {
		res.FailWithMessage("删除用户失败", c)
		return
	}
	res.OkWithMessage("删除用户成功！", c)
}

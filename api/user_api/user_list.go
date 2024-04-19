package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils"
	"gvb_server/utils/jwt"
)

func (userApi UserApi) UserListView(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		res.FailWithMessage("请输入token", c)
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	list, count, _ := common.CommonList(models.UserModel{}, common.Option{
		PageInfo: page,
		Debug:    true,
	})
	users := []models.UserModel{}
	for _, user := range list {
		if claims.Role != int(ctype.PermissionAdmin) {
			user.Password = ""
		}
		user.Email = utils.DesensitizationEmail(user.Email)
		user.Tel = utils.DesensitizationTel(user.Tel)
		users = append(users, user)
	}
	res.OkWithList(users, count, c)
}

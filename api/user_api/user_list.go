package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (userApi UserApi) UserListView(c *gin.Context) {
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	list, count, _ := common.CommonList(&models.UserModel{}, common.Option{
		PageInfo: page,
		Debug:    true,
	})
	res.OkWithList(list, count, c)
}

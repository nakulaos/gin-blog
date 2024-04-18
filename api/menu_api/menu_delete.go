package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (menuApi MenuApi) MenuDeleteView(c *gin.Context) {
	var removeRequest models.RemoveRequest
	err := c.ShouldBindJSON(&removeRequest)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var deleteList []models.MenuModel
	count := global.DB.Find(&deleteList, removeRequest.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("没有指定的菜单！", c)
		return
	}
	err = global.DB.Model(&deleteList).Association("Banners").Clear()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	err = global.DB.Delete(&deleteList).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("删除菜单成功！", c)
}

package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (advertiseApi AdvertiseApi) AdvertiseDeleteApi(c *gin.Context) {
	var removeRequest models.RemoveRequest
	err := c.ShouldBindJSON(&removeRequest)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var deleteList []models.AdvertModel
	err = global.DB.Find(&deleteList, removeRequest.IDList).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	err = global.DB.Delete(&deleteList).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithMessage("删除成功！", c)
}

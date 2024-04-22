package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (tagApi TagApi) TagDeleteView(c *gin.Context) {
	var removeRequest models.RemoveRequest
	err := c.ShouldBindJSON(&removeRequest)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var deleteList []models.TagModel
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

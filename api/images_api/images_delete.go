package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (imagesApi ImagesApi) ImagesDeleteView(c *gin.Context) {
	var removeRequest models.RemoveRequest
	var DeleteList []models.BannerModel
	c.ShouldBindJSON(&removeRequest)
	count := global.DB.Find(&DeleteList, removeRequest.IDList).RowsAffected
	if count == 0 {
		res.OkWithMessage("没有要删除的图片", c)
		return
	}
	if err := global.DB.Delete(&DeleteList).Error; err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("删除%d张图片成功！", count), c)
}

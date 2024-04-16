package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (advertiseApi AdvertiseApi) AdvertiseUpdateView(c *gin.Context) {
	var cr AdvertiseRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}
	id := c.Param("id")
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("没有这个id的广告", c)
		return
	}
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("更新广告失败！", c)
		return
	}
	res.OkWithMessage("更新广告成功！", c)
}

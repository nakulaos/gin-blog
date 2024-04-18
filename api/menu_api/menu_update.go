package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (menuApi MenuApi) MenuUpdateView(c *gin.Context) {

	var cr MenuRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	id := c.Param("id")
	var menuModel models.MenuModel
	if err := global.DB.Take(&menuModel, id).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if err := global.DB.Model(&menuModel).Association("Banners").Clear(); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	//如果更新时选择了banner，那就插入banner，操作第三张表
	if len(cr.ImageSortList) > 0 {
		var bannerList []models.MenuBannerModel
		for _, val := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: val.ImageID,
				Sort:     val.Sort,
			})
		}
		if err := global.DB.Create(&bannerList).Error; err != nil {
			global.Log.Error(err.Error())
			return
		}
	}

	//普通更新menu
	maps := structs.Map(&cr)
	if err := global.DB.Model(&models.MenuModel{}).Updates(maps); err != nil {
		res.FailWithMessage("更新广告失败！", c)
		return
	}
	res.OkWithMessage("修改广告成功!", c)
}

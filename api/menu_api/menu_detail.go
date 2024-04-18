package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (menuApi MenuApi) MenuDetailView(c *gin.Context) {
	var menu models.MenuModel
	id := c.Param("id")
	if count := global.DB.Take(&menu, id).RowsAffected; count == 0 {
		res.FailWithMessage("没有找到这个菜单!", c)
		return
	}
	//res.OkWithData(menuIDList, c)

	var menubanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menubanners)
	//res.OkWithData(menubanners, c)

	banner_response := []Banner{}
	for _, menubanner := range menubanners {
		if menu.ID != menubanner.MenuID {
			continue
		}
		banner_response = append(banner_response, Banner{
			ID:   menubanner.BannerID,
			Path: menubanner.BannerModel.Path,
		})
	}

	menu_response := MenuResponse{
		MenuModel: menu,
		Banners:   banner_response,
	} //是一个包含menu，和对应相联的banner的切片
	res.OkWithData(menu_response, c)
}

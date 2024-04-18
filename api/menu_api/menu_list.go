package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}
type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (menuApi MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	//res.OkWithData(menuIDList, c)

	var menubanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menubanners)
	//res.OkWithData(menubanners, c)

	var menu_response []MenuResponse //是一个包含menu，和对应相联的banner的切片
	for _, menu := range menuList {
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
		menu_response = append(menu_response, MenuResponse{
			MenuModel: menu,
			Banners:   banner_response,
		})
	}
	res.OkWithData(menu_response, c)

}

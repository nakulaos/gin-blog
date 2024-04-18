package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (menuApi MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameResponse []MenuNameResponse
	global.DB.Model(&models.MenuModel{}).Select("id", "title", "path").Scan(&menuNameResponse)
	res.OkWithData(menuNameResponse, c)
}

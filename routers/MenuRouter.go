package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func MenuRouter(enter *gin.RouterGroup) {
	MenuApi := api.ApiGroupApp.MenuApi
	enter.POST("/menu", MenuApi.MenuCreateView)
	enter.GET("/menu", MenuApi.MenuListView)
}

package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func AdvertiseRouter(enter *gin.RouterGroup) {
	advertiseApi := api.ApiGroupApp.AdvertiseApi
	enter.POST("/advert", advertiseApi.AdvertiseCreateView)
	enter.GET("/advert", advertiseApi.AdvertiseListView)
	enter.PUT("/advert/:id", advertiseApi.AdvertiseUpdateView)
}

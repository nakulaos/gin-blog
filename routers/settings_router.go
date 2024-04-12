package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func SettingsRouter(enter *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
	enter.GET("/settings", settingsApi.SettingsInfoView)
}

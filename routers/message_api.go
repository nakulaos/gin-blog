package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func MessageRouter(enter *gin.RouterGroup) {
	messageApi := api.ApiGroupApp.MessageApi
	enter = enter.Group("")
	enter.Use(middleware.JwtAuth())
	enter.POST("/message", messageApi.MessageSendView)
	enter.GET("/message_all", messageApi.MessageAllView)
	enter.GET("/message_mine", messageApi.MessageMineView)
}

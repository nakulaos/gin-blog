package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func TagRouter(enter *gin.RouterGroup) {
	tagApi := api.ApiGroupApp.TagApi
	enter = enter.Group("")
	enter.Use(middleware.JwtAuth())
	enter.POST("/tag", tagApi.TagCreateView)
	enter.GET("/tag", tagApi.TagListView)
	enter.PUT("/tag/:id", tagApi.TagUpdateView)
	enter.DELETE("/tag", tagApi.TagDeleteView)
}

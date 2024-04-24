package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func QQLoginRouter(enter *gin.Engine) {
	qqloginApi := api.ApiGroupApp.QQLoginApi
	enter.GET("/login/callback/qq", qqloginApi.QQLoginApiView)
}

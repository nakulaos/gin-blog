package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func UserRouter(enter *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UserApi
	enter.POST("/userlogin_email", userApi.UserLoginEmailView)
}

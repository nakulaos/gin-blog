package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func UserRouter(enter *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UserApi
	enter.POST("/user_login_email", userApi.UserLoginEmailView)
	enter.GET("/user_list", userApi.UserListView)
}

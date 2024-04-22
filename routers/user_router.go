package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func UserRouter(enter *gin.RouterGroup) {
	userApi := api.ApiGroupApp.UserApi
	enter.POST("/user_login_email", userApi.UserLoginEmailView)
	enter.GET("/user_list", middleware.JwtAuth(), userApi.UserListView)
	enter.PUT("/admin_update_userinfo", middleware.JwtAdmin(), userApi.UserUpdateUserinfoView)
	enter.PUT("/user_update_password", middleware.JwtAuth(), userApi.UserUpdatePasswordView)
	enter.POST("/user_logout", middleware.JwtAuth(), userApi.UserLogoutView)
	enter.DELETE("/user_delete", middleware.JwtAdmin(), userApi.UserDeleteView)
	enter.POST("/user_bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)
}

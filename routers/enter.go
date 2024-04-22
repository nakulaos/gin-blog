package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gvb_server/global"
	"gvb_server/middleware"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Enter := gin.Default()
	Enter.Use(middleware.Cors())
	Enter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	EnterApi := Enter.Group("/api")
	SettingsRouter(EnterApi)
	ImageRouter(EnterApi)
	AdvertiseRouter(EnterApi)
	MenuRouter(EnterApi)
	UserRouter(EnterApi)
	TagRouter(EnterApi)

	return Enter
}

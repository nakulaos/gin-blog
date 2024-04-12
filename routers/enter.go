package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Enter := gin.Default()

	EnterApi := Enter.Group("/api")
	SettingsRouter(EnterApi)

	return Enter
}

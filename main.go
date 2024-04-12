package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化logrus
	core.InitLogger()
	//连接数据库，初始化gorm
	core.InitGorm()
	//初始化路由
	enter := routers.InitRouter()
	global.Log.Infof("项目运行在:%s", global.Config.System.Addr())
	enter.Run(global.Config.System.Addr())

}

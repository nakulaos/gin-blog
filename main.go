package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化logrus
	core.InitLogger()
	global.Log.Warnln("aa")
	global.Log.Error("aa")
	global.Log.Infof("aa")

	//连接数据库，初始化gorm
	core.InitGorm()

}

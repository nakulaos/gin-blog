package main

import (
	"gvb_server/core"
	"gvb_server/flag"
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

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//初始化路由
	enter := routers.InitRouter()
	global.Log.Infof("项目运行在:%s", global.Config.System.Addr())
	err := enter.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Error(err.Error())
	}

}

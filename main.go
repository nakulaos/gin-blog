package main

import (
	"gvb_server/core"
)

func main() {
	//读取配置文件
	core.InitConf()
	//连接数据库
	core.InitGorm()

}

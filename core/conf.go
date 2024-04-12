package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"gvb_server/models/res"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConfig error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init unmarshal:%v", err)
	}
	log.Println("config init yamlFile success!")
	global.Config = c
}
func UpdateYaml(c *gin.Context) {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	global.Log.Infof("修改yaml成功！")
	res.OkWithMessage("修改yaml成功！", c)
}

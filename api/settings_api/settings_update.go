package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (s *SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	name := c.Param("name")

	switch name {
	case "siteinfo":
		choice := config.SiteInfo{}
		if err := c.ShouldBindJSON(&choice); err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.SiteInfo = choice
		core.UpdateYaml(c)
	case "email":
		choice := config.Email{}
		if err := c.ShouldBindJSON(&choice); err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.Email = choice
		core.UpdateYaml(c)
	case "qq":
		choice := config.QQ{}
		if err := c.ShouldBindJSON(&choice); err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.QQ = choice
		core.UpdateYaml(c)
	case "qiniu":
		choice := config.QiNiu{}
		if err := c.ShouldBindJSON(&choice); err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.QiNiu = choice
		core.UpdateYaml(c)
	case "jwt":
		choice := config.Jwt{}
		if err := c.ShouldBindJSON(&choice); err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.Jwt = choice
		core.UpdateYaml(c)
	default:
		return
	}

}

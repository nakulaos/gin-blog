package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (s *SettingsApi) SettingsInfoView(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "siteinfo":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		return
	}
}

package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (s *SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var siteinfo config.SiteInfo
	if err := c.ShouldBindJSON(&siteinfo); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	global.Config.SiteInfo = siteinfo
	core.UpdateYaml(c)

}

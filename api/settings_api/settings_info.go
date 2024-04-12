package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
)

func (s *SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(3344, c)
}

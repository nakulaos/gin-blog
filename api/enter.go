package api

import (
	"gvb_server/api/images_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)

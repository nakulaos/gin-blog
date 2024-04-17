package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi  settings_api.SettingsApi
	ImageApi     images_api.ImagesApi
	AdvertiseApi advert_api.AdvertiseApi
	MenuApi      menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)

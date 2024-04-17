package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/ctype"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}
type MenuRequest struct {
	MenuTitle     string      `json:"menu_title"`      //
	MenuTitleEn   string      `json:"menu_title_en"`   //
	Slogan        string      `json:"slogan"`          //
	Abstract      ctype.Array `json:"abstract"`        //
	AbstractTime  int         `json:"abstract_time"`   //切换的时间,单位 秒
	BannerTime    int         `json:"banner_time"`     //切换的时间，单位 秒
	Sort          int         `json:"sort"`            //菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list"` //具体图片的顺序
}

func (menuApi MenuApi) MenuCreateView(c *gin.Context) {

}

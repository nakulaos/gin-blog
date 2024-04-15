package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type AdvertiseRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`           //
	Href   string `json:"href" binding:"required,url" msg:"请输入合法链接"`      //
	Images string `json:"images" binding:"required,url" msg:"请输入合法图片url"` //
	IsShow *bool  `json:"is_show" binding:"required" msg:"请选择是否展示"`       //
}

func (advertiseApi *AdvertiseApi) AdvertiseCreateView(c *gin.Context) {
	var cr AdvertiseRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	//判断有没有重复广告，按照title查

	err = global.DB.Take(&models.AdvertModel{}, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该广告标题已存在", c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: *cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("添加广告成功！", c)
}

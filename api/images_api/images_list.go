package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type Page struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
	Sort  string `form:"sort"`
}

func (img ImagesApi) ImageListView(c *gin.Context) {
	var cr Page
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	var imagesList []models.BannerModel

	count := global.DB.Find(&imagesList).RowsAffected
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}
	global.DB.Limit(cr.Limit).Offset(offset).Find(&imagesList)
	res.OkWithData(gin.H{
		"count": count,
		"list":  imagesList,
	}, c)
}

package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (imagesApi ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	imagesList, count, err := common.CommonList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		return
	}
	res.OkWithList(imagesList, count, c)

}

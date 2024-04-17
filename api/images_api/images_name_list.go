package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImagesResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"` //图片路径
	Name string `json:"name"` //图片名称
}

// @ImagesNameListView 查询图片名称
// @Tags 图片名称查询
// @Summary 图片名称
// @Description 查询图片名称列表
// @Router /api/images_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[[]ImagesResponse]}
func (imagesApi ImagesApi) ImagesNameListView(c *gin.Context) {
	var imagesResponse []ImagesResponse
	global.DB.Model(&models.BannerModel{}).Select("id", "path", "name").Scan(&imagesResponse)
	res.OkWithData(imagesResponse, c)
}

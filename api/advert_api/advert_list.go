package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"strings"
)

// @AdvertiseListView 查询广告列表
// @Tags 广告查询
// @Summary 查询广告
// @Description 查询广告列表
// @Param data query models.PageInfo  true "表示多个参数"
// @Router /api/advert [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (advertiseApi AdvertiseApi) AdvertiseListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	//如果是admin，就返回全部，否则只返回is_show = true的一些广告
	referer := c.GetHeader("refer")
	if strings.Contains(referer, "admin") {

		list, count, err := common.CommonList(models.AdvertModel{}, common.Option{
			PageInfo: cr,
			Debug:    true,
		})
		if err != nil {
			res.FailWithMessage("查找广告列表失败！", c)
			return
		}
		res.OkWithList(list, count, c)
		return
	}

	list, count, err := common.CommonList(models.AdvertModel{IsShow: true}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMessage("查找广告列表失败！", c)
		return
	}
	res.OkWithList(list, count, c)

}

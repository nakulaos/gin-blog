package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

func (tagApi TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}

	list, count, err := common.CommonList(models.TagModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMessage("查找tag列表失败！", c)
		return
	}
	res.OkWithList(list, count, c)
}

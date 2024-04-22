package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (tagApi TagApi) TagUpdateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	id := c.Param("id")
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("没有这个id的tag", c)
		return
	}
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("更新tag失败！", c)
		return
	}
	res.OkWithMessage("更新tag成功！", c)
}

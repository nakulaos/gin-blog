package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标签标题" structs:"title"`
}

func (tagApi TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	//判断有没有重复tag，按照title查

	err = global.DB.Take(&models.TagModel{}, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该tag标题已存在", c)
		return
	}
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("添加tag成功！", c)
}

package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type UpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"id不能为空!"`
	Name string `json:"name" binding:"required" msg:"name不能为空!"`
}

func (imagesApi ImagesApi) ImagesUpdateView(c *gin.Context) {
	var cr UpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	var image models.BannerModel
	if err := global.DB.Take(&image, cr.ID).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMessage("没有找到此id对应的图片", c)
		return
	}
	if err := global.DB.Model(&image).Update("name", cr.Name).Error; err != nil {
		res.FailWithMessage(fmt.Sprintf("更改图片名称失败:", err.Error()), c)
		return
	}
	res.OkWithMessage("更改图片名称成功！", c)

}

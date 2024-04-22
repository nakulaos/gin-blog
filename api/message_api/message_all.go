package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils/jwt"
)

func (messageApi MessageApi) MessageAllView(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if claims.Role != 1 {
		res.FailWithMessage("对不起，您没有权限查看所有消息", c)
		return
	}
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	list, count, err := common.CommonList(models.MessageModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMessage("查找消息列表失败！", c)
		return
	}
	res.OkWithList(list, count, c)
}

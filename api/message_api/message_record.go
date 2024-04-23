package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
)

// 查找自己与某用户的聊天记录
type MessageRecordRequest struct {
	AnotherUserID uint `json:"another_user_id" binding:"required" msg:"接收者id不能为空"`
}

func (messageApi MessageApi) MessageRecordView(c *gin.Context) {
	var cr MessageRecordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var Messages []models.MessageModel
	if err := global.DB.Order("created_at asc").Find(&Messages, "(send_user_id = ? or rev_user_id = ?) and (send_user_id = ? or rev_user_id = ?)", claims.UserID, claims.UserID, cr.AnotherUserID, cr.AnotherUserID).Error; err != nil {
		res.FailWithMessage("查找失败", c)
		return
	}
	res.OkWithData(Messages, c)
}

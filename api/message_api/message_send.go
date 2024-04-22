package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
)

type MessageSendRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required" msg:"请输入发送人id" `
	RevUserID  uint   `json:"rev_user_id" binding:"required" msg:"请输入接收者id"`
	Content    string `json:"content" binding:"required" msg:"请输入发送内容"`
}

func (messageApi MessageApi) MessageSendView(c *gin.Context) {
	var cr MessageSendRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if claims.UserID != cr.SendUserID {
		res.FailWithMessage("发送者不一致", c)
		return
	}

	var sendUser, revUser models.UserModel
	if err := global.DB.Take(&sendUser, cr.SendUserID).Error; err != nil {
		res.FailWithMessage("发件人id不存在", c)
		return
	}
	if err := global.DB.Take(&revUser, cr.RevUserID).Error; err != nil {
		res.FailWithMessage("收件人id不存在", c)
		return
	}

	message := models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  revUser.NickName,
		RevUserAvatar:    revUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}
	if err := global.DB.Create(&message).Error; err != nil {
		res.FailWithMessage("发送消息失败", c)
		return
	}
	res.OkWithMessage("发送消息成功!", c)
}

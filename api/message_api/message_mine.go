package message_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwt"
	"time"
)

type Message struct {
	SendUserID       uint      `json:"send_user_id"`
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"`
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	CreateAt         time.Time `json:"create_at"` //创建时间，默认将会展示最新的时间，跟qq微信等及时通信软件一样
	Content          string    `json:"content"`
	Count            int       `json:"count"`
}
type MessageGroup map[uint]*Message

type MessageResponse []Message

func (messageApi MessageApi) MessageMineView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var messages []models.MessageModel
	/*这里要注意，不能写成
	var messageMap MessageGroup
	要不然只是定义，没有为map分配空间，会报nil map错误
	正确写法：
	messageMap := make(MessageGroup)
	或
	messageMap := MessageGroup{}
	*/
	messageMap := MessageGroup{}
	if err := global.DB.Order("created_at asc").Find(&messages, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID).Error; err != nil {
		res.FailWithMessage("查找消息失败！", c)
		return
	}
	for _, message := range messages {
		Msg := Message{
			SendUserID:       message.SendUserID,
			SendUserNickName: message.SendUserNickName,
			SendUserAvatar:   message.SendUserAvatar,
			RevUserID:        message.RevUserID,
			RevUserNickName:  message.RevUserNickName,
			RevUserAvatar:    message.RevUserAvatar,
			CreateAt:         message.CreatedAt,
			Content:          message.Content,
			Count:            1,
		}
		idSum := message.SendUserID + message.RevUserID
		val, ok := messageMap[idSum]
		if !ok {
			//新的分组
			fmt.Println(val, ok)
			fmt.Println(Msg)
			messageMap[idSum] = &Msg
		} else {
			//旧的分组，只保留最新的，然后count++就行了
			Msg.Count = messageMap[idSum].Count + 1
			messageMap[idSum] = &Msg
		}
	}
	var MsgResponse MessageResponse
	for _, group := range messageMap {
		MsgResponse = append(MsgResponse, *group)
	}
	res.OkWithData(MsgResponse, c)
}

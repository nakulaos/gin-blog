package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/plugins/sendmail"
	"math/rand"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"请输入正确的邮箱！"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (userApi UserApi) UserBindEmailView(c *gin.Context) {
	var cr BindEmailRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	//第一次给邮箱发一个验证码
	if cr.Code == nil {
		code := rand.Int31n(10000)
		if err := sendmail.NewCode().Send(cr.Email, fmt.Sprintf("您的验证码是：%04d，如非本人操作，请忽略本邮件。", code)); err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		//存入session

	} else { /*第二次可以完成绑定*/

	}

}

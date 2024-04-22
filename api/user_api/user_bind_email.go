package user_api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/sendmail"
	"gvb_server/utils/jwt"
	"gvb_server/utils/pwd"
	"math/rand"
	"strconv"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"请输入正确的邮箱！"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (userApi UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr BindEmailRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}
	//第一次给邮箱发一个验证码
	//预备好session
	session := sessions.Default(c)

	if cr.Code == nil {
		code := rand.Int31n(10000)
		if err := sendmail.NewCode().Send(cr.Email, fmt.Sprintf("您的验证码是：%04d，如非本人操作，请忽略本邮件。", code)); err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		//发送成功，将验证码存入session
		session.Set("valid_code", code)
		//发送成功，将邮箱存入session
		session.Set("email", cr.Email)
		err := session.Save()
		if err != nil {
			res.FailWithMessage("session保存失败！", c)
			return
		}
		res.OkWithMessage("验证码已发送！", c)
		return

	}
	/*第二次可以完成绑定*/
	if len(cr.Password) < 4 {
		res.FailWithMessage("密码长度不能小于4位!", c)
		return
	}
	code := session.Get("valid_code")
	email := session.Get("email")

	//	fmt.Printf("%t %t", code, *cr.Code)
	Code, _ := strconv.Atoi(*cr.Code)
	//fmt.Printf("%t %t", code, Code)
	if code != int32(Code) {
		res.FailWithMessage("验证码错误!", c)
		return
	}
	if email != cr.Email {
		res.FailWithMessage("收到验证码的邮箱和输入邮箱不一致，无法绑定！", c)
		return
	}
	//没问题,操作数据库
	hashPwd := pwd.HashPassword(cr.Password)
	var user models.UserModel
	if err := global.DB.Take(&user, claims.UserID).Error; err != nil {
		global.Log.Error(err.Error(), c)
		res.FailWithMessage("绑定失败！", c)
		return
	}
	if err := global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error; err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("绑定失败！", c)
		return
	}

	res.OkWithMessage("绑定邮箱成功，下次登陆请填写用户名和对应的密码！", c)

}

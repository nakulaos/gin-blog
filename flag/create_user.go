package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/pwd"
)

var (
	UserName         string
	NickName         string
	Password         string
	Confirm_Password string
	Email            string
)

func CreateUser(Level string) {
	fmt.Print("请输入用户名:")
	fmt.Scan(&UserName)
	fmt.Print("请输入用户的Nickname:")
	fmt.Scan(&NickName)
	fmt.Print("请输入密码：")
	fmt.Scan(&Password)
	fmt.Print("请确认密码：")
	fmt.Scan(&Confirm_Password)
	fmt.Print("请输入邮箱：")
	fmt.Scan(&Email)
	if Password != Confirm_Password {
		global.Log.Error("两次密码输入不一致！")
		return
	}

	if count := global.DB.Take(&models.UserModel{}, "user_name = ?", UserName).RowsAffected; count != 0 {
		global.Log.Error("该用户名已存在！")
		return
	}
	permission := ctype.PermissionUser
	if Level == "admin" {
		permission = ctype.PermissionAdmin
	}

	userModel := models.UserModel{
		UserName:   UserName,
		NikeName:   NickName,
		Password:   pwd.HashPassword(Password),
		Email:      Email,
		Addr:       "内网",
		Avatar:     "/uploads/avatar/default.jpg",
		IP:         "127.0.0.1",
		Role:       permission,
		SignStatus: ctype.SignEmail,
	}
	if err := global.DB.Create(&userModel).Error; err != nil {
		global.Log.Error("用户创建失败!")
		return
	}
	global.Log.Info(fmt.Sprintf("用户 %s 创建成功！", UserName))

}

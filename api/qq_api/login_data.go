package qq_api

import (
	"gvb_server/global"
	"gvb_server/models"
)

// LoginDataInsert 插入一条登陆数据
func LoginDataInsert(userID uint, nickName, token, ip, device, addr string) {
	global.DB.Create(&models.LoginDataModel{
		UserID:   userID,
		IP:       ip,
		NickName: nickName,
		Token:    token,
		Device:   device,
		Addr:     addr,
	})
}

package models

import "gvb_server/models/ctype"

// 用户表
type UserModel struct {
	MODEL
	NikeName       string           `gorm:"size:36" json:"nike_name"`                                                         //对外昵称
	UserName       string           `gorm:"size:36" json:"user_name"`                                                         //用户名
	PassWord       string           `gorm:"size:128" json:"password"`                                                         //密码
	Avatar         string           `gorm:"size:256" json:"avatar_id"`                                                        //头像id
	Email          string           `gorm:"size:128" json:"email"`                                                            //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                               //手机号
	Addr           string           `gorm:"size:64" json:"addr"`                                                              //地址
	Token          string           `gorm:"size:64" json:"token"`                                                             //其他平台的唯一id
	IP             string           `gorm:"size:128" json:"ip"`                                                               //ip地址
	Role           ctype.Role       `gorm:"size:12;default:1" json:"role"`                                                    //权限
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`                                              //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:AuthID" json:"-"`                                                       //发布的文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:auth2_collects;joinForeignKey:AuthID;JoinReferences:ArticleID" json:"-"` //收藏的文章列表
}

//func ParseRole(role ctype.Role) string {
//
//	switch role {
//	case ctype.PermissionAdmin:
//		return "管理员"
//	case ctype.PermissionUser:
//		return "普通用户"
//	case ctype.PermissionVisitor:
//		return "游客"
//	case ctype.PermissionDisableUser:
//		return "禁言用户"
//	}
//}

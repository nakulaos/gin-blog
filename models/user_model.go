package models

import "gvb_server/models/ctype"

// 用户表
type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:36" json:"nick_name"`                                                              //对外昵称
	UserName       string           `gorm:"size:36" json:"user_name"`                                                              //用户名
	Password       string           `gorm:"size:128" json:"-"`                                                                     //这样可以不展示给前端，至于登录我可以另定义一个request                                                                    //密码
	Avatar         string           `gorm:"size:256" json:"avatar_id"`                                                             //头像id
	Email          string           `gorm:"size:128" json:"email"`                                                                 //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                                    //手机号
	Addr           string           `gorm:"size:64" json:"addr"`                                                                   //地址
	Token          string           `gorm:"size:64" json:"token"`                                                                  //其他平台的唯一id
	IP             string           `gorm:"size:128" json:"ip"`                                                                    //ip地址
	Role           ctype.Role       `gorm:"size:12;default:2" json:"role"`                                                         //权限
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`                                                   //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`                                                            //发布的文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"` //收藏的文章列表
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

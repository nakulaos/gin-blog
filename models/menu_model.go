package models

import "gvb_server/models/ctype"

// 菜单表
type MenuModel struct {
	MODEL
	Title        string        `gorm:"size:32" json:"itle"`                                                                       //
	Path         string        `gorm:"size:32" json:"path"`                                                                       //
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                     //口号
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                               //简介
	AbstractTime int           `json:"abstract_time"`                                                                             //简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"` //菜单的图片列表
	BannerTime   int           `json:"banner_time"`                                                                               //菜单图片的切换时间 为0表示不切换
	Sort         int           `gorm:"size:1-" json:"sort"`                                                                       //菜单的顺序
}

package models

import "time"

type MODEL struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"-"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
type PageInfo struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Key   string `form:"key"`
	Sort  string `form:"sort"`
}

package models

import "time"

type MODEL struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `gorm:"default:current_timestamp" json:"-"`
}

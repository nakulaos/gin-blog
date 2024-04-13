package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo      //调用查询list时，需要绑定的相应模型
	Debug           bool //调用查询list时，是否开启gorm的logger打印
}

func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	count = DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return
}

package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	MODEL
	Path             string                 `json:"path"`                //图片路径
	Hash             string                 `json:"hash"`                //图片的hash值
	Name             string                 `gorm:"size:38" json:"name"` //图片名称
	ImagesStorgeType ctype.ImagesStorgeType `gorm:"default:1" json:"images_storge"`
}

func (b *BannerModel) BeforeDelete(db *gorm.DB) error {
	if b.ImagesStorgeType == ctype.Local {
		//在执行删除数据库数据时，先删除本地文件目录下的图片
		if err := os.Remove(b.Path); err != nil {
			global.Log.Error(err.Error())
			return err
		}
	}
	return nil
}

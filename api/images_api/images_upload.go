package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
)

var FileMap = map[string]string{
	".jpg":  "",
	".jpeg": "",
	".png":  "",
	".icon": "",
	".svg":  "",
	".gif":  "",
	".webp": "",
}

type ResUploadLocal struct {
	Msg       string `json:"msg"`
	IsSuccess bool   `json:"is_success"`
	FileName  string `json:"file_name"`
}

func (imagesApi *ImagesApi) ImagesUploadView(c *gin.Context) {
	var multipartFile *multipart.Form
	var err error
	if multipartFile, err = c.MultipartForm(); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	imagesList, ok := multipartFile.File["images"]
	if !ok {
		res.FailWithMessage("没有获取到这个参数的文件", c)
		return
	}
	var resSlice []ResUploadLocal
	for index, image := range imagesList {
		//判断是否为图片
		ext := filepath.Ext(image.Filename)
		ext = strings.ToLower(ext)
		if _, ok := FileMap[ext]; !ok {
			resSlice = append(resSlice, ResUploadLocal{
				Msg:       fmt.Sprintf("第 %d 个文件不是一个图片类型", index+1),
				IsSuccess: false,
				FileName:  image.Filename,
			})
			continue
		}
		//判断图片是否超过大小限制
		size := float64(image.Size) / float64(1024*1024)
		if size > float64(global.Config.LocalUpload.Size) {
			resSlice = append(resSlice, ResUploadLocal{
				Msg:       fmt.Sprintf("上传第 %d 张图片失败，图片大小为 %.2f MB,本地上传图片大小不得超过：%d MB", index+1, size, global.Config.LocalUpload.Size),
				IsSuccess: false,
				FileName:  image.Filename,
			})
			continue
		}
		fileobj, err := image.Open()
		if err != nil {
			global.Log.Error(err)
		}

		//去数据库查这个图片是否存在
		byteData, err := io.ReadAll(fileobj)
		ImageHash := utils.MD5(byteData)
		var bannerModel models.BannerModel
		var count int64
		if global.DB.Take(&bannerModel, "hash = ?", ImageHash).Count(&count); count == 1 {
			//找到了,直接返回
			resSlice = append(resSlice, ResUploadLocal{
				Msg:       "数据库中有这张图片了",
				IsSuccess: true,
				FileName:  bannerModel.Path,
			})
			continue
		}
		//数据库中没这张图片
		msg := "上传成功"
		success := true
		TargetFilePath := filepath.Join(global.Config.LocalUpload.UploadFilePath, image.Filename)
		if err := c.SaveUploadedFile(image, TargetFilePath); err != nil {
			global.Log.Error(fmt.Sprintf("上传第%d张图片失败"), index+1)
			msg = "上传失败"
			success = false
		}
		resSlice = append(resSlice, ResUploadLocal{
			Msg:       msg,
			IsSuccess: success,
			FileName:  TargetFilePath,
		})

		//图片入库
		if success {
			global.DB.Create(&models.BannerModel{
				Path: TargetFilePath,
				Hash: ImageHash,
				Name: image.Filename,
			})
		}

	}
	res.OkWithData(resSlice, c)
}

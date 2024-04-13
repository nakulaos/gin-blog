package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"mime/multipart"
	"path/filepath"
)

type ResUploadLocal struct {
	Msg       string `json:"msg"`
	IsSuccess bool   `json:"is_success"`
	FileName  string `json:"file_name"`
}

func (sr *ImagesApi) ImagesUploadView(c *gin.Context) {
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
		size := float64(image.Size) / float64(1024*1024)
		if size > float64(global.Config.LocalUpload.Size) {
			resSlice = append(resSlice, ResUploadLocal{
				Msg:       fmt.Sprintf("上传失败，本地上传图片大小不超过：%d MB", global.Config.LocalUpload.Size),
				IsSuccess: false,
				FileName:  image.Filename,
			})
			continue
		}
		msg := "上传成功"
		success := true
		TargetFilePath := filepath.Join(global.Config.LocalUpload.UploadFilePath, image.Filename)
		if err := c.SaveUploadedFile(image, TargetFilePath); err != nil {
			global.Log.Error(fmt.Sprintf("上传第%d张图片失败"), index)
			msg = "上传失败"
			success = false
		}
		resSlice = append(resSlice, ResUploadLocal{
			Msg:       msg,
			IsSuccess: success,
			FileName:  TargetFilePath,
		})
	}
	res.OkWithData(resSlice, c)
}

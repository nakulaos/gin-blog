package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func ImageRouter(enter *gin.RouterGroup) {
	imagesApi := api.ApiGroupApp.ImageApi
	enter.POST("/images_upload", imagesApi.ImagesUploadView)

}

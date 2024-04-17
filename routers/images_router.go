package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func ImageRouter(enter *gin.RouterGroup) {
	imagesApi := api.ApiGroupApp.ImageApi
	enter.GET("/images", imagesApi.ImageListView)            //分页查询
	enter.GET("/images_names", imagesApi.ImagesNameListView) //图片名称列表
	enter.POST("/images", imagesApi.ImagesUploadView)        //上传图片
	enter.PUT("/images", imagesApi.ImagesUpdateView)         //修改图片名称
	enter.DELETE("/images", imagesApi.ImagesDeleteView)      //批量删除图片
}

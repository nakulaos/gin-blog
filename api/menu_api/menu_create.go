package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}
type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请输入菜单标题" structs:"title"` //
	Path          string      `json:"path" binding:"required" msg:"请输入菜单路径" structs:"path"`   //
	Slogan        string      `json:"slogan" structs:"slogan"`                                //
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`                            //
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`                  //切换的时间,单位 秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`                      //切换的时间，单位 秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单的序号" structs:"sort"`  //菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`                            //具体图片的顺序
}
type ResMenuBanner struct {
	Msg    string `json:"msg"`
	Status bool   `json:"status"`
}

func (menuApi MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMyMessage(&cr, err, c)
		return
	}

	//先查有没有这个菜单
	err := global.DB.Take(&models.MenuModel{}, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("已经有这个菜单标题了!", c)
		return
	}
	//sort不能重复
	count := global.DB.Select("id").Take(&models.MenuModel{}, "sort = ?", cr.Sort).RowsAffected
	if count != 0 {
		res.FailWithMessage("sort号重复，换一个!", c)
		return
	}
	/*这种方法不会给某些默认值，比如create_at,update_at。你传了就insert，不传就不给。
	maps := structs.Map(&menuRequest)
	if err := global.DB.Model(&models.MenuModel{}).Create(&maps).Error; err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}*/
	menu := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menu).Error
	if err != nil {
		res.FailWithMessage("创建菜单失败！", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.FailWithMessage("创建菜单成功！", c)
		return
	}

	//若cr.ImageSortList！=0，还要去menu_banner数据表中插入数据
	var menubannerlist []models.MenuBannerModel
	var response []ResMenuBanner
	for _, val := range cr.ImageSortList {
		//这里其实还要判断图片是否有这个序号
		if count := global.DB.Select("id").Take(&models.BannerModel{}, "id = ?", val.ImageID).RowsAffected; count == 0 {
			response = append(response, ResMenuBanner{
				Msg:    fmt.Sprintf("没有id为%d的图片,此图片将不会关联成功", val.ImageID),
				Status: false,
			})
			continue
		}
		response = append(response, ResMenuBanner{
			Msg:    fmt.Sprintf("有此id为%d的图片", val.ImageID),
			Status: true,
		})
		menubannerlist = append(menubannerlist, models.MenuBannerModel{
			MenuID:   menu.ID,
			BannerID: val.ImageID,
			Sort:     val.Sort,
		})
	}

	if err := global.DB.Create(&menubannerlist).Error; err != nil {
		res.OkWithData(response, c)
		res.FailWithMessage("menu关联banner失败，但menu创建成功!", c)
		return
	}
	res.OkWithData(map[string]any{
		"prosess": response,
		"final":   "创建菜单成功！",
	}, c)

}

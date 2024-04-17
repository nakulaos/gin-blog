package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type AdvertiseRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`       //
	Href   string `json:"href" binding:"required,url" msg:"请输入合法链接"`                  //
	Images string `json:"images" binding:"required,url" msg:"请输入合法图片url"`             //
	IsShow *bool  `json:"is_show" binding:"required" msg:"请选择是否展示" structs:"is_show"` //
}

func main() {
	ch := false
	u1 := AdvertiseRequest{
		Title:  "xxx",
		Href:   "xx",
		Images: "xx",
		IsShow: &ch,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}

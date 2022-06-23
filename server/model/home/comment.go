package home

import "server/global"

type Comment struct {
	global.GVA_MODEL
	BlogId uint `json:"blog_id" gorm:"comment:博客ID"`
	Name string `json:"name" gorm:"comment:评论人"`
	Phone  string `json:"phone" gorm:"comment:评论人联系"`
	Message string `json:"message" gorm:"comment:评论内容"`
	Status int `json:"status" gorm:"default:1;comment:1显示，0隐藏"`
}

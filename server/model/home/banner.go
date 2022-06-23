package home

import "server/global"

type Banner struct {
	global.GVA_MODEL
	Title    string `json:"title" form:"title" gorm:"comment:banner标题"`          // banner标题
	Cover    string `json:"cover" form:"cover" gorm:"comment:banner图片地址"`        // banner图片
	Browses  int    `json:"browses" form:"browses" gorm:"default:0;comment:浏览次数"`          //浏览次数
	Comments int    `json:"comments"  form:"comments"  gorm:"default:0;comment:评论次数"`      //评论次数
	Content  string `json:"content" form:"content"  gorm:"comment:内容;type:text"` //banner内容
}

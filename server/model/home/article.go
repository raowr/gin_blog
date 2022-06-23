package home

import "server/global"

type Article struct {
	global.GVA_MODEL
	Title    string `json:"title" form:"title" gorm:"comment:文章标题"`          // 文章标题
	Cover    string `json:"cover" form:"cover" gorm:"comment:文章图片地址"`        // 文章图片
	Browses  int    `json:"browses" form:"browses" gorm:"default:0;comment:浏览次数"`          //浏览次数
	Comments int    `json:"comments"  form:"comments"  gorm:"default:0;comment:评论次数"`      //评论次数
	Content  string `json:"content" form:"content"  gorm:"comment:内容;type:text"` //文章内容
}

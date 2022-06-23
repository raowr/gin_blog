package request

import (
	"server/model/common/request"
	"server/model/home"
	"time"
)

type SearchArticleParams struct {
	home.Article
	request.PageInfo
	CreatedAtStart time.Time `json:"created_at_start"`
	CreatedAtEnd time.Time `json:"created_at_end"`
}

type EditArticleParams struct {
	Id 		 uint `json:"id"`//自增ID
	Title    string `json:"title" form:"title" gorm:"comment:文章标题"`          // 文章标题
	Cover    string `json:"cover" form:"cover" gorm:"comment:文章图片地址"`       // 文章图片
	Content  string `json:"content" form:"content"  gorm:"comment:内容;type:text"` //文章内容
}

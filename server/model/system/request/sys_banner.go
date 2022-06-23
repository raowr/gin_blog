package request

import (
	"server/model/common/request"
	"server/model/home"
	"time"
)

type SearchBannerParams struct {
	home.Banner
	request.PageInfo
	CreatedAtStart time.Time `json:"created_at_start"`
	CreatedAtEnd time.Time `json:"created_at_end"`
}

type EditBannerParams struct {
	Id 		 uint `json:"id"`//自增ID
	Title    string `json:"title" form:"title" gorm:"comment:banner标题"`          // banner标题
	Cover    string `json:"cover" form:"cover" gorm:"comment:banner图片地址"`       // banner图片
	Content  string `json:"content" form:"content"  gorm:"comment:内容;type:text"` //banner内容
}

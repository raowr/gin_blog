package msg

import (
	"server/model/common/request"
	"server/model/home"
)

type SearchListBlogParams struct {
	request.PageInfo
	home.Article
}

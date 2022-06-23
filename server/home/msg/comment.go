package msg

import (
	"server/model/common/request"
	"server/model/home"
)

type SearchListCommentParams struct {
	request.PageInfo
	home.Comment
}


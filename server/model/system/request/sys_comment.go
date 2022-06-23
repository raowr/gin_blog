package request

import (
	"server/model/common/request"
	"server/model/home"
	"time"
)

type SearchCommentParams struct {
	home.Comment
	request.PageInfo
	CreatedAtStart time.Time `json:"created_at_start"`
	CreatedAtEnd time.Time `json:"created_at_end"`
}

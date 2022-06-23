package msg

import (
	"server/model/common/request"
	"server/model/home"
)

type SearchListBannerParams struct {
	request.PageInfo
	home.Banner
}

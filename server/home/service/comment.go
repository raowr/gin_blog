package service

import (
	"fmt"
	"server/global"
	"server/home/msg"
	"server/model/home"
)

type ServiceComment struct {}

func (this *ServiceComment) List(Params msg.SearchListCommentParams) (CommentList []home.Comment,total int64,err error) {
	limit := Params.PageSize
	offset := Params.PageSize * (Params.Page -1)
	db := global.GVA_DB.Model(&home.Comment{})

	db = db.Where("blog_id = ? AND status = 1",Params.BlogId)
	err = db.Count(&total).Error
	if err != nil {
		fmt.Errorf("banner列表查询错误: %v", err)
	}else {
		err = db.Order("id desc").Limit(limit).Offset(offset).Find(&CommentList).Error
	}
	return CommentList, total, err
}

func (this *ServiceComment) Create(commentInfo home.Comment) error {
	return global.GVA_DB.Create(&commentInfo).Error
}

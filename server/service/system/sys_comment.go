package system

import (
	"server/global"
	commonRequest "server/model/common/request"
	"server/model/home"
	"server/model/system/request"
	"time"
)

type CommentService struct {}

func (this *CommentService) ListComment(SearchCommentParams request.SearchCommentParams) (list []home.Comment, total int64, err error) {
	limit := SearchCommentParams.PageSize
	offset := SearchCommentParams.PageSize * (SearchCommentParams.Page - 1)
	db := global.GVA_DB.Model(&home.Comment{})
	if SearchCommentParams.Name != ""{
		db = db.Where("name LIKE ?","%"+SearchCommentParams.Name+"%")
	}
	if SearchCommentParams.Message != ""{
		db = db.Where("message LIKE ?","%"+SearchCommentParams.Message+"%")
	}
	if SearchCommentParams.Status > -1 {
		db = db.Where("status = ?",SearchCommentParams.Status)
	}
	nilTime := time.Time{}
	if SearchCommentParams.CreatedAtStart != nilTime{
		db = db.Where("created_at > ?",SearchCommentParams.CreatedAtStart)
	}

	if SearchCommentParams.CreatedAtEnd != nilTime{
		db = db.Where("created_at < ?",SearchCommentParams.CreatedAtEnd)
	}
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (this *CommentService) ChangCommentStatusIdsByIds(ids commonRequest.IdsReq,status int) (err error) {
	return global.GVA_DB.Model(&home.Comment{}).Where("id in ?", ids.Ids).Update("status", status).Error
}

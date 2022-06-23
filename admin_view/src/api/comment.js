import service from '@/utils/request'

// @Tags SysComment
// @Summary 列表comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "列表comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /comment/listComment [post]
// {
//  page     int
//	pageSize int
// }
export const listComment = (data) => {
  return service({
    url: '/comment/listComment',
    method: 'post',
    data: data
  })
}

// @Tags SysComment
// @Summary 批量显示comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "批量显示comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /comment/showCommentIdsByIds [post]
export const showCommentIdsByIds = (data) => {
  return service({
    url: '/comment/showCommentIdsByIds',
    method: 'post',
    data: data
  })
}

// @Tags SysComment
// @Summary 批量隐藏comment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "批量隐藏comment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"隐藏成功"}"
// @Router /comment/hideCommentIdsByIds [post]
export const hideCommentIdsByIds = (data) => {
  return service({
    url: '/comment/hideCommentIdsByIds',
    method: 'post',
    data: data
  })
}

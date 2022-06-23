import service from '@/utils/request'

// @Tags SysArticle
// @Summary 添加article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "添加article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /article/addArticle [post]
export const addArticle = (data) => {
  return service({
    url: '/article/addArticle',
    method: 'post',
    data: data
  })
}

// @Tags SysArticle
// @Summary 列表article
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "列表article"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /article/listArticle [post]
// {
//  page     int
//	pageSize int
// }
export const listArticle = (data) => {
  return service({
    url: '/article/listArticle',
    method: 'post',
    data: data
  })
}

// @Tags SysArticle
// @Summary article信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "article信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /article/infoArticle [post]
// {
//  ID     int
// }
export const infoArticle = (data) => {
  return service({
    url: '/article/infoArticle',
    method: 'post',
    data: data
  })
}

// @Tags SysArticle
// @Summary article编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "article编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"编辑成功"}"
// @Router /article/editArticle [post]
// {
//  ID      int
//  title   string
//  content string
//  cover   string
// }
export const editArticle = (data) => {
  return service({
    url: '/article/editArticle',
    method: 'post',
    data: data
  })
}

// @Tags SysArticle
// @Summary article删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "article删除"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /article/deleteArticle [post]
// {
//  ID     int
// }
export const deleteArticle = (data) => {
  return service({
    url: '/article/deleteArticle',
    method: 'post',
    data: data
  })
}

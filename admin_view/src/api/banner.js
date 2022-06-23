import service from '@/utils/request'

// @Tags SysBanner
// @Summary 添加banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "添加banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /banner/addBanner [post]
export const addBanner = (data) => {
  return service({
    url: '/banner/addBanner',
    method: 'post',
    data: data
  })
}

// @Tags SysBanner
// @Summary 列表banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "列表banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"添加成功"}"
// @Router /banner/listBanner [post]
// {
//  page     int
//	pageSize int
// }
export const listBanner = (data) => {
  return service({
    url: '/banner/listBanner',
    method: 'post',
    data: data
  })
}

// @Tags SysBanner
// @Summary banner信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "banner信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /banner/infoBanner [post]
// {
//  ID     int
// }
export const infoBanner = (data) => {
  return service({
    url: '/banner/infoBanner',
    method: 'post',
    data: data
  })
}

// @Tags SysBanner
// @Summary banner编辑
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "banner编辑"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"编辑成功"}"
// @Router /banner/editBanner [post]
// {
//  ID      int
//  title   string
//  content string
//  cover   string
// }
export const editBanner = (data) => {
  return service({
    url: '/banner/editBanner',
    method: 'post',
    data: data
  })
}

// @Tags SysBanner
// @Summary banner删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "banner删除"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banner/deleteBanner [post]
// {
//  ID     int
// }
export const deleteBanner = (data) => {
  return service({
    url: '/banner/deleteBanner',
    method: 'post',
    data: data
  })
}

package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"server/global"
	"server/home/msg"
	"server/model/common/response"
	"server/utils"
	"strconv"
)

type HandleBlog struct {}

func (this *HandleBlog) List(c *gin.Context)  {
	var pageInfo msg.SearchListBlogParams
	pageInfo.Page,_ = strconv.Atoi( c.Param("page"))
	pageInfo.PageSize,_ = strconv.Atoi(c.Param("pageSize"))
	if err := utils.Verify(pageInfo.PageInfo, utils.ListBlog); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := serviceBlog.List(pageInfo.Article,pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"title": "博客列表",//
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

func (this *HandleBlog) ListPost(c *gin.Context)  {
	var pageInfo msg.SearchListBlogParams
	_= c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.ListBlog); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := serviceBlog.List(pageInfo.Article,pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (this *HandleBlog) Detail(c *gin.Context)  {
	var blogInfo msg.SearchListBlogParams
	id,_ := strconv.Atoi( c.Param("id"))
	blogInfo.ID = uint(id)
	if err := utils.Verify(blogInfo, utils.InfoArticle); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//增加浏览数
	_ = serviceBlog.BrowsesArticle(blogInfo.ID)

	if blogDetail, err := serviceBlog.InfoArticle(blogInfo.ID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {

		//获取评论
		var commentParams msg.SearchListCommentParams
		commentParams.BlogId = blogInfo.ID
		commentParams.Page = 1
		commentParams.PageSize = 4
		if listComment, total, err := serviceComment.List(commentParams); err != nil {
			global.GVA_LOG.Error("获取评论失败!", zap.Error(err))
			response.FailWithMessage("获取评论失败", c)
		}else {
			c.HTML(http.StatusOK, "detail.html", gin.H{
				"title": "博客信息-花无百日红，人无百日好",//
				"info":     blogDetail,
				"listComment" : listComment,
				"totalComment" :total,
				"page" :commentParams.Page,
				"pageSize" :commentParams.PageSize,
			})
		}

	}
}

func (this *HandleBlog) ToPre(c *gin.Context)  {
	id, _ := serviceBlog.FindID(c,"pre")
	response.OkWithDetailed(response.PageResultId{
		Id:     id,
	}, "获取成功", c)
}

func (this *HandleBlog) ToNext(c *gin.Context)  {
	id, _ := serviceBlog.FindID(c,"next")
	response.OkWithDetailed(response.PageResultId{
		Id:     id,
	}, "获取成功", c)
}



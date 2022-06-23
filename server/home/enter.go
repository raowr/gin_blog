package home

import "server/home/handler"

type HomeGroup struct {
	IndexGroup handler.HandlerIndex
	AboutGroup handler.HandlerAbout
	HandleBlog handler.HandleBlog
	HandleContact handler.HandleContact
	HandleComment handler.HandleComment

}

var HomeGroupApp = new(HomeGroup)

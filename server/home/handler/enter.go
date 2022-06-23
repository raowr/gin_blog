package handler

import "server/home/service"

type HandleGroup struct {
	HandlerIndex
	HandlerAbout
	HandleBlog
	HandleContact
}

var (
	serviceIndex   = new(service.ServiceIndex)
	serviceBlog    = new(service.ServiceBlog)
	serviceComment = new(service.ServiceComment)
)

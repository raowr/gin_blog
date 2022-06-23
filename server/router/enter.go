package router

import (
	"server/router/example"
	"server/router/home"
	"server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Home  	 home.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

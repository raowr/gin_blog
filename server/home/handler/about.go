package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerAbout struct {}

func (this *HandlerAbout)About(c *gin.Context)  {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

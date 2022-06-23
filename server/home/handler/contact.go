package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandleContact struct {}

func (this *HandleContact) Index(c *gin.Context)  {
	c.HTML(http.StatusOK, "contact.html", gin.H{})
}

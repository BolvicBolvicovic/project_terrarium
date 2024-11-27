package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.HTML(http.StatusOK, "root.tmpl", gin.H{})
}

func NewWorld(c *gin.Context) {
	c.HTML(http.StatusOK, "terrarium.tmpl", gin.H{})
}

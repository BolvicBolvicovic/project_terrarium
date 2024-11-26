package server

import (
	"github.com/BolvicBolvicovic/project_terrarium/api"
	"github.com/gin-gonic/gin"
)

func buildRouter() *gin.Engine {
	router := gin.Default()

	router.StaticFile("/terrarium.js", "./templates/components/terrarium.js")
	router.LoadHTMLGlob("templates/*/**")

	router.GET("/", api.Root)

	return router
}

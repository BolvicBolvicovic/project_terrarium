package server

import (
	"github.com/BolvicBolvicovic/project_terrarium/api"
	"github.com/gin-gonic/gin"
)

func buildRouter() *gin.Engine {
	router := gin.Default()

	router.StaticFile("/terrarium.js", "./templates/components/terrarium.js")
	router.StaticFile("/root.js", "./templates/pages/root.js")
	router.StaticFile("/root.css", "./templates/pages/root.css")

	router.LoadHTMLGlob("templates/*/**")

	router.GET("/", api.Root)
	router.GET("/newWorld", api.NewWorld)

	return router
}

package routes

import (
	"github.com/gin-gonic/gin"
)

var apiBasePath string = "/api"

func AddRoutes(router *gin.Engine) {
	api := router.Group(apiBasePath)
	{
		// Define a route that handles GET requests to the root URL
		api.GET("", indexHandler)
	}

	AddUsersRoutes(router, api)
}

func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, world!",
	})
}

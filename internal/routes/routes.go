package routes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	// Define a route that handles GET requests to the root URL
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})
}

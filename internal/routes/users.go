package routes

import (
	"github.com/gin-gonic/gin"
)

var usersPath string = "/users"

func AddUsersRoutes(router *gin.Engine, group *gin.RouterGroup) {
	users := group.Group(usersPath)
	{
		users.GET("", getUsers)
	}
}

func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Users API",
	})
}

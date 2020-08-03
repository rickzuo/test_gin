package router

import "github.com/gin-gonic/gin"
import . "test_gin/api"

func InitROuter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/hello",Hello)
	router.GET("/users",GetUserList)
	router.POST("/users/add",AddUser)

	return router
}

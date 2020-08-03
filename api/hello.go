package api

import "github.com/gin-gonic/gin"
import "net/http"

func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"msg": "hello", "code": 0})
}


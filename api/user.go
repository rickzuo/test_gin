package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"test_gin/models"
)
import . "test_gin/utils"
func AddUser(c *gin.Context) {
	var user models.User
	var name = c.Query("name")
	var age int
	ageStr := c.Query("age")

	age, _ = strconv.Atoi(ageStr)

	user.Name = name
	user.Age = age
	user.Add()

	data := make(map[string]interface{})

	data["name"] = user.Name
	data["age"] = user.Age

	c.JSON(http.StatusOK, data)

}

func GetUserList(c *gin.Context) {
	var user models.User
	userList, _ := user.GetUserList()
	rs := ReturnData(0, "ok", userList)
	c.JSON(http.StatusOK, rs)
	//users, _ := user.GetUserList()
	//c.JSON(200, res(0, "ok", users))
}

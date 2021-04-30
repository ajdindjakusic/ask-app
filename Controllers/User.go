package Controllers

import (
	"fmt"
	"go/first-api/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
 	var user []Models.User
 	err := Models.GetAllUsers(&user)
 	if err != nil {
  	 c.AbortWithStatus(http.StatusNotFound)
 	} else {
  	 c.JSON(http.StatusOK, user)
 	}
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
 	var user Models.User
 	c.BindJSON(&user)
 	err := Models.CreateUser(&user)
 	if err != nil {
  	 fmt.Println(err.Error())
  	c.AbortWithStatus(http.StatusNotFound)
 	} else {
  	 c.JSON(http.StatusOK, user)
 	}
}

func UpdateUser(c *gin.Context) {
 	var user Models.User
 	id := c.Params.ByName("id")
 	err := Models.GetUserByID(&user, id)
 	if err != nil {
  	 c.JSON(http.StatusNotFound, user)
 	}
 	c.BindJSON(&user)
 	err = Models.UpdateUser(&user, id)
 	if err != nil {
  	 c.AbortWithStatus(http.StatusNotFound)
 	} else {
  	 c.JSON(http.StatusOK, user)
 	}
}

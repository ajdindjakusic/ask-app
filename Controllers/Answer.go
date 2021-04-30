package Controllers

import (
	"fmt"
	"go/first-api/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAnswerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var answer Models.Answer
	err := Models.GetAnswerByID(&answer, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, answer)
	}
}

func GetAnswerByQuestionID(c *gin.Context) {
 	id := c.Params.ByName("id")
 	var answer []Models.Answer
 	err := Models.GetAnswerByQuestionID(&answer, id)
 	if err != nil {
 	c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, answer)
	}
}

func CreateAnswer(c *gin.Context) {
	var answer Models.Answer
 	c.BindJSON(&answer)
 	err := Models.CreateAnswer(&answer)
 	if err != nil {
 	 fmt.Println(err.Error())
  	 c.AbortWithStatus(http.StatusNotFound)
 	} else {
  	 c.JSON(http.StatusOK, answer)
 	}
}

func UpdateAnswer(c *gin.Context) {
	var answer Models.Answer
	id := c.Params.ByName("id")
	err := Models.GetAnswerByID(&answer, id)
	if err != nil {
	 c.JSON(http.StatusNotFound, answer)
	}
	c.BindJSON(&answer)
	err = Models.UpdateAnswer(&answer, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, answer)
	}
}

func DeleteAnswer(c *gin.Context) {
	var answer Models.Answer
	id := c.Params.ByName("id")
	err := Models.DeleteAnswer(&answer, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func GetUsersWithMostAnswers(c *gin.Context) {
	var user []Models.User
	err := Models.GetUsersWithMostAnswers(&user)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, user)
	}
}
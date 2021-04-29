package Controllers

import (
	"fmt"
	"go/first-api/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

//GetAnswerByID ... Used to get answer for editing
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

//GetAnswersByQuestionID ... Get answers when on Question page
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

//CreateAnswer - for creating new questions
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

//UpdateAnswer ... Update the user information - for profile editing
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

   //DeleteAnswer ... Delete the answer on specific question
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
	var answer []Models.Answer
	err := Models.GetUsersWithMostAnswers(&answer)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, answer)
	}
   }
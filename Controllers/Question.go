package Controllers

import (
	"fmt"
	"go/first-api/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

//GetQuestions - gets all questions for home page
func GetQuestions(c *gin.Context) {
 var question []Models.Question
 err := Models.GetAllQuestions(&question)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, question)
 }
}

func LoadMoreQuestions(c *gin.Context) {
	limit := c.Params.ByName("limit")
	offset := c.Params.ByName("offset")
	var question []Models.Question
	err := Models.LoadMoreQuestions(&question, limit, offset)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, question)
	}
   }

//GetQuestionByUserID ... Get questions per specific user
func GetQuestionByUserID(c *gin.Context) {
	id := c.Params.ByName("id")
	var question []Models.Question
	err := Models.GetQuestionByUserID(&question, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, question)
	}
   }

//CreateQuestion - for creating new questions
func CreateQuestion(c *gin.Context) {
 var question Models.Question
 c.BindJSON(&question)
 err := Models.CreateQuestion(&question)
 if err != nil {
  fmt.Println(err.Error())
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, question)
 }
}


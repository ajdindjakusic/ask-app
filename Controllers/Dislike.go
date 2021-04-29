package Controllers

import (
	"fmt"
	"go/first-api/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

   func GetQuestionDislikes(c *gin.Context) {
	id := c.Params.ByName("id")
	var dislike Models.Dislike
	err := Models.GetQuestionDislikes(&dislike, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, dislike)
	}
   }

   func GetAnswerDislikes(c *gin.Context) {
	id := c.Params.ByName("id")
	var dislike Models.Dislike
	err := Models.GetAnswerDislikes(&dislike, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, dislike)
	}
   }

func PostQuestionDislike(c *gin.Context) {
	var dislike Models.Dislike
	c.BindJSON(&dislike)
	err := Models.PostQuestionDislike(&dislike)
	if err != nil {
	 fmt.Println(err.Error())
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, dislike)
	}
   }

func PostAnswerDislike(c *gin.Context) {
	var dislike Models.Dislike
	c.BindJSON(&dislike)
	err := Models.PostAnswerDislike(&dislike)
	if err != nil {
	 fmt.Println(err.Error())
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, dislike)
	}
   }
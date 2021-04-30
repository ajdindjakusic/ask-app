package Controllers

import (
	"fmt"
	"go/first-api/Models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HotQuestions(c *gin.Context) {
	var question []Models.Question
	err := Models.HotQuestions(&question)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, question)
	}
}

func GetQuestionLikes(c *gin.Context) {
	id := c.Params.ByName("id")
	var like Models.Like
	err := Models.GetQuestionLikes(&like, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, like)
	}
}

func GetAnswerLikes(c *gin.Context) {
	id := c.Params.ByName("id")
	var like Models.Like
	err := Models.GetAnswerLikes(&like, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, like)
	}
}

func PostQuestionLike(c *gin.Context) {
	var like Models.Like
	c.BindJSON(&like)
	err := Models.PostQuestionLike(&like)
	if err != nil {
	 fmt.Println(err.Error())
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, like)
	}
}

func PostAnswerLike(c *gin.Context) {
	var like Models.Like
	c.BindJSON(&like)
	err := Models.PostAnswerLike(&like)
	if err != nil {
	 fmt.Println(err.Error())
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, like)
	}
}
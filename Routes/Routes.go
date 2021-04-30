package Routes

import (
	"go/first-api/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

 r := gin.New()
 ask := r.Group("/ask-api")
 {
  ask.GET("user", Controllers.GetUsers)
  ask.GET("user/:id", Controllers.GetUserByID)
  ask.GET("user/most", Controllers.GetUsersWithMostAnswers)
  ask.POST("user", Controllers.CreateUser)
  ask.PUT("user/:id", Controllers.UpdateUser)

  ask.GET("question", Controllers.GetQuestions)
  ask.GET("question/limit/:limit/offset/:offset", Controllers.LoadMoreQuestions)
  ask.GET("question/:id", Controllers.GetQuestionByUserID)
  ask.GET("question/hot", Controllers.HotQuestions)
  ask.POST("question", Controllers.CreateQuestion)

  ask.GET("answer/:id", Controllers.GetAnswerByQuestionID)
  ask.POST("answer", Controllers.CreateAnswer)
  ask.PUT("answer/:id", Controllers.UpdateAnswer)
  ask.DELETE("answer/:id", Controllers.DeleteAnswer)

  ask.GET("question/like/:id", Controllers.GetQuestionLikes)
  ask.GET("answer/like/:id", Controllers.GetAnswerLikes)
  ask.POST("question/like", Controllers.PostQuestionLike)
  ask.POST("answer/like", Controllers.PostAnswerLike)

  ask.GET("question/dislike/:id", Controllers.GetQuestionDislikes)
  ask.GET("answer/dislike/:id", Controllers.GetAnswerDislikes)
  ask.POST("question/dislike", Controllers.PostQuestionDislike)
  ask.POST("answer/dislike", Controllers.PostAnswerDislike)
  
  return r
 }
}

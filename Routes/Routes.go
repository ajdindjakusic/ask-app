package Routes

import (
	"go/first-api/Controllers"
	// "go/first-api/Service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

 r := gin.New()
 ask := r.Group("/ask-api")
 {
  ask.GET("user", Controllers.GetUsers)
  ask.GET("user/:id", Controllers.GetUserByID)
  ask.POST("user", Controllers.CreateUser)
  ask.PUT("user/:id", Controllers.UpdateUser)

  ask.GET("question", Controllers.GetQuestions)
  ask.GET("question/limit/:limit/offset/:offset", Controllers.LoadMoreQuestions)
  ask.GET("question/:id", Controllers.GetQuestionByUserID)
  ask.POST("question", Controllers.CreateQuestion)

  ask.GET("answer/:id", Controllers.GetAnswerByQuestionID)
  ask.GET("answer/most", Controllers.GetUsersWithMostAnswers)
  ask.POST("answer", Controllers.CreateAnswer)
  ask.PUT("answer/:id", Controllers.UpdateAnswer)
  ask.DELETE("answer/:id", Controllers.DeleteAnswer)

  ask.GET("like/hot", Controllers.HotQuestions)
  ask.GET("like/question/:id", Controllers.GetQuestionLikes)
  ask.GET("like/answer/:id", Controllers.GetAnswerLikes)
  ask.POST("like/question", Controllers.PostQuestionLike)
  ask.POST("like/answer", Controllers.PostAnswerLike)

  ask.GET("dislike/question/:id", Controllers.GetQuestionDislikes)
  ask.GET("dislike/answer/:id", Controllers.GetAnswerDislikes)
  ask.POST("dislike/question", Controllers.PostQuestionDislike)
  ask.POST("dislike/answer", Controllers.PostAnswerDislike)
  
  return r
 }
}

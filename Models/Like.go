package Models
import (
 "go/first-api/Config"
//  "fmt"
_ "github.com/go-sql-driver/mysql"
)

func HotQuestions(like *[]Like) (err error) {
	if err = Config.DB.Raw("select question_id, count(*) as total from likes where question_id is not null group by question_id order by total desc limit 10").Scan(&like).Error; err != nil {
	 return err
	}
	return nil
   }

func GetQuestionLikes(like *Like, id string) (err error) {
	if err = Config.DB.Raw("select question_id, count(*) as total from likes where question_id = ?", id).Scan(&like).Error; err != nil {
		return err
	}
	return nil
}

func GetAnswerLikes(like *Like, id string) (err error) {
		if err = Config.DB.Raw("select answer_id, count(*) as total from likes where answer_id = ?", id).Scan(&like).Error; err != nil {
		 return err
		}
		return nil
	   }

func PostQuestionLike(like *Like) (err error) {
		if err = Config.DB.Exec("insert into likes (user_id, question_id) values (?, ?)", like.User_id, like.Question_id).Error; err != nil {
	
	  return err
	 }
	 return nil
	}

func PostAnswerLike(like *Like) (err error) {
	if err = Config.DB.Exec("insert into likes (user_id, answer_id) values (?, ?)", like.User_id, like.Answer_id).Error; err != nil {

  return err
 }
 return nil
}
	
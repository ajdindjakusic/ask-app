package Models
import (
 "go/first-api/Config"
//  "fmt"
_ "github.com/go-sql-driver/mysql"
)
//GetAllQuestions - Show all questions on home page with 20 limit
func GetAllQuestions(question *[]Question) (err error) {
 if err = Config.DB.Raw("select * from question order by created desc limit 20").Scan(&question).Error; err != nil {
  return err
 }
 return nil
}

//LoadMoreQuestions - Loads another 20 questions on button press
func LoadMoreQuestions(question *[]Question, limit, offset string ) (err error) {
	if err = Config.DB.Raw("select * from question order by created desc limit ? offset ?", limit, offset).Scan(&question).Error; err != nil {
	 return err
	}
	return nil
   }

//GetQuestionByUserID - Shows only questions by specific users
func GetQuestionByUserID(question *[]Question, id string) (err error) {
	if err = Config.DB.Raw("select * from question where user_id = ?", id).Scan(&question).Error; err != nil {
	 return err
	}
	return nil
   }

//CreateQuestion - Create new question
func CreateQuestion(question *Question) (err error) {
	//  if err = Config.DB.Create(question).Error; err != nil {
		if err = Config.DB.Exec("insert into question (text, user_id) values (?, ?)", question.Text, 1 ).Error; err != nil {
	
	  return err
	 }
	 return nil
	}

  
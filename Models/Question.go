package Models
import (
 	"go/first-api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllQuestions(question *[]Question) (err error) {
 	if err = Config.DB.Raw("select * from question order by created desc limit 20").Scan(&question).Error; err != nil {
 	 return err
 	}
 	return nil
}

func LoadMoreQuestions(question *[]Question, limit, offset string ) (err error) {
	if err = Config.DB.Raw("select * from question order by created desc limit ? offset ?", limit, offset).Scan(&question).Error; err != nil {
	 return err
	}
	return nil
}

func GetQuestionByUserID(question *[]Question, id string) (err error) {
	if err = Config.DB.Raw("select * from question where user_id = ?", id).Scan(&question).Error; err != nil {
	 return err
	}
	return nil
}

func CreateQuestion(question *Question) (err error) {
	if err = Config.DB.Exec("insert into question (text, user_id) values (?, ?)", question.Text, 1 ).Error; err != nil {
	  return err
	}
	return nil
}

func HotQuestions(question *[]Question) (err error) {
	if err = Config.DB.Raw("select likes.question_id, question.text, count(*) as total from likes inner join question on likes.question_id=question.question_id group by question_id order by total desc limit 10").Scan(&question).Error; err != nil {
	 return err
	}
	return nil
}


  
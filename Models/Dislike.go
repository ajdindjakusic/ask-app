package Models
import (
 	"go/first-api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetQuestionDislikes(dislike *Dislike, id string) (err error) {
	if err = Config.DB.Raw("select question_id, count(*) as total from likes where question_id = ?", id).Scan(&dislike).Error; err != nil {
	 return err
	}
	return nil
}

func GetAnswerDislikes(dislike *Dislike, id string) (err error) {
	if err = Config.DB.Raw("select answer_id, count(*) as total from likes where answer_id = ?", id).Scan(&dislike).Error; err != nil {
	 return err
	}
	return nil
}

func PostQuestionDislike(dislike *Dislike) (err error) {
	if err = Config.DB.Exec("insert into dislikes (user_id, question_id) values (?, ?)", dislike.User_id, dislike.Question_id).Error; err != nil {
	 return err
	}
	return nil
}

func PostAnswerDislike(dislike *Dislike) (err error) {
	if err = Config.DB.Exec("insert into dislikes (user_id, answer_id) values (?, ?)", dislike.User_id, dislike.Answer_id).Error; err != nil {
	 return err
 	}
 	return nil
}
	
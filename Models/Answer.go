package Models

import (
	"go/first-api/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAnswerByID(answer *Answer, id string) (err error) {
	if err = Config.DB.Where("answer_id = ?", id).First(&answer).Error; err != nil {
	 return err
	}
	return nil
}

func GetAnswerByQuestionID(answer *[]Answer, id string) (err error) {
	if err = Config.DB.Raw("select * from answer where question_id = ? and visible = true", id).Scan(&answer).Error; err != nil {
	 return err
	}
	return nil
}

func CreateAnswer(answer *Answer) (err error) {
	if err = Config.DB.Exec("insert into answer (text, question_id, user_id) values (?, ?, ?)", answer.Text, answer.Question_id, answer.User_id ).Error; err != nil {
	 return err
	}
	return nil
}

func UpdateAnswer(answer *Answer, id string) (err error) {
   	if err = Config.DB.Exec("update answer set text = ? where answer_id = ?", answer.Text, id).Error; err != nil {
	 return err
	}
	return nil
}

func DeleteAnswer(answer *Answer, id string) (err error) {
    if err = Config.DB.Exec("update answer set visible = false where answer_id = ?", id).Error; err != nil {
	 return err
	}
	return nil
}

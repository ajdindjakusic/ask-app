package Models

import (
	"fmt"
	"go/first-api/Config"
	//  "fmt"
	_ "github.com/go-sql-driver/mysql"
)

//GetAnswerByID ... Get answer for editing
func GetAnswerByID(answer *Answer, id string) (err error) {
	if err = Config.DB.Where("answer_id = ?", id).First(&answer).Error; err != nil {
	 return err
	}
	return nil
   }

//GetAnswerByQuestionID - Shows only answers for specific question
func GetAnswerByQuestionID(answer *[]Answer, id string) (err error) {
	if err = Config.DB.Raw("select * from answer where question_id = ?", id).Scan(&answer).Error; err != nil {
	 return err
	}
	return nil
   }

//CreateAnswer - Create new question
func CreateAnswer(answer *Answer) (err error) {
	//  if err = Config.DB.Create(user).Error; err != nil {
		if err = Config.DB.Exec("insert into answer (text, question_id, user_id) values (?, ?, ?)", answer.Text, answer.Question_id, answer.User_id ).Error; err != nil {
	
	  return err
	 }
	 return nil
	}


//UpdateUser ... Update answer by user
func UpdateAnswer(answer *Answer, id string) (err error) {
	fmt.Println(answer)
   //  Config.DB.Save(user)
   if err = Config.DB.Exec("update user set text = ? where answer_id = ?", answer.Text, id).Error; err != nil {
	   return err
	  }
	return nil
   }

   //DeleteAnswer ... just making it invisible
func DeleteAnswer(answer *Answer, id string) (err error) {
	fmt.Println(answer)
   if err = Config.DB.Exec("update user set visible = false where answer_id = ?", answer.Text, id).Error; err != nil {
	   return err
	  }
	return nil
   }

   //GetUsersWithMostAnswers ... Get Users with most answers
func GetUsersWithMostAnswers(answer *[]Answer) (err error) {
	if err = Config.DB.Raw("select answer.user_id, user.username, count(*) as total from answer INNER JOIN user on answer.user_id=user.user_id group by user_id order by total desc").Scan(&answer).Error; err != nil {
	 return err
	}
	return nil
   }
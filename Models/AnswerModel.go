package Models

type Answer struct {
	Answer_id  		uint   `json:"answer_id"`
	Text 		 	string `json:"text"`
	Question_id     uint   `json:"question_id"`
	User_id      	uint   `json:"user_id"`
	Username        string `json:"username"`
}

func (b *Answer) TableName() string {
	return "answer"
}
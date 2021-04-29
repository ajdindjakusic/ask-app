package Models

type Question struct {
	Question_id  uint   `json:"question_id"`
	Text 		 string `json:"text"`
	User_id      uint 	`json:"user_id"`
}

func (b *Question) TableName() string {
	return "question"
}
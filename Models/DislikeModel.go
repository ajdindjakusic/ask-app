package Models

type Dislike struct {
	User_id      uint   `json:"user_id"`
	Question_id  uint   `json:"question_id"`
	Answer_id    uint   `json:"answer_id"`
	Total        uint   `json:"total"`
}

func (b *Dislike) TableName() string {
	return "dislike"
}
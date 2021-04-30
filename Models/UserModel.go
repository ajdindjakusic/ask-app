package Models

type User struct {
	User_id  	uint   `json:"user_id"`
	Username 	string `json:"username"`
	Email    	string `json:"email"`
	Password 	string `json:"password"`
	Token 	 	string `json:"token"`
}

func (b *User) TableName() string {
	return "user"
}
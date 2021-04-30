package Models
import (
    "go/first-api/Config"
    _ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(user *[]User) (err error) {
    if err = Config.DB.Raw("select * from user").Scan(&user).Error; err != nil {
     return err
    }
    return nil
}

func GetUserByID(user *User, id string) (err error) {
    if err = Config.DB.Where("user_id = ?", id).First(user).Error; err != nil {
     return err
    }
    return nil
}

func GetUsersWithMostAnswers(user *[]User) (err error) {
	if err = Config.DB.Raw("select answer.user_id, user.username, count(*) as total from answer INNER JOIN user on answer.user_id=user.user_id group by user_id order by total desc").Scan(&user).Error; err != nil {
	 return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
    if err = Config.DB.Exec("insert into user (username, email, password) values (?, ?, ?)", user.Username, user.Email, user.Password).Error; err != nil {
   return err
  }
  return nil
}

func UpdateUser(user *User, id string) (err error) {
    if err = Config.DB.Exec("update user set username = ?, email = ?, password = ? where user_id = ?", user.Username, user.Email, user.Password, id).Error; err != nil {
	   return err
    }
    return nil
}

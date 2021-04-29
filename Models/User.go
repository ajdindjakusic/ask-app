package Models
import (
 "go/first-api/Config"
 "fmt"
_ "github.com/go-sql-driver/mysql"
)
//GetAllUsers Get all users - not needed for app
func GetAllUsers(user *[]User) (err error) {
 if err = Config.DB.Raw("select * from user").Scan(&user).Error; err != nil {
  return err
 }
 return nil
}
//CreateUser ... Creates new user on registration and creates a JWT token
func CreateUser(user *User) (err error) {
//  if err = Config.DB.Create(user).Error; err != nil {
	if err = Config.DB.Exec("insert into user (username, email, password) values (?, ?, ?)", user.Username, user.Email, user.Password).Error; err != nil {

  return err
 }
 return nil
}
//GetUserByID ... Used for login after JWT implementation
func GetUserByID(user *User, id string) (err error) {
  if err = Config.DB.Where("user_id = ?", id).First(user).Error; err != nil {
   return err
  }
  return nil
 }
//UpdateUser ... Update users username and email
func UpdateUser(user *User, id string) (err error) {
 fmt.Println(user)
//  Config.DB.Save(user)
if err = Config.DB.Exec("update user set username = ?, email = ? where user_id = ?", user.Username, user.Email, id).Error; err != nil {
	return err
   }
 return nil
}

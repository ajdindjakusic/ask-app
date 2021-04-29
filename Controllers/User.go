package Controllers

import (
	// "encoding/json"
	"fmt"
	"go/first-api/Models"
	"net/http"
	// "os/user"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// var jwtKey = []byte("secret_key")

// type Credentials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }


//GetUsers ... Get all users - not needed, just for testing
func GetUsers(c *gin.Context) {
 
//  var credentials Credentials
//  json.NewDecoder(c.Request.Body).Decode(&credentials)
 
//  expectedPassword, ok := user[credentials.Username]

//  if !ok || expectedPassword != credentials.Password {
// 	c.AbortWithStatus(http.StatusNotFound)
//  }

 var user []Models.User
 err := Models.GetAllUsers(&user)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, user)
 }
}

//GetUserByID ... Get the user by id - for profile page
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
	 c.AbortWithStatus(http.StatusNotFound)
	} else {
	 c.JSON(http.StatusOK, user)
	}
   }

//CreateUser ... Create new user apon register
func CreateUser(c *gin.Context) {
 var user Models.User
 c.BindJSON(&user)
 err := Models.CreateUser(&user)
 if err != nil {
  fmt.Println(err.Error())
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, user)
 }
}

//UpdateUser ... Update the user information - for profile editing
func UpdateUser(c *gin.Context) {
 var user Models.User
 id := c.Params.ByName("id")
 err := Models.GetUserByID(&user, id)
 if err != nil {
  c.JSON(http.StatusNotFound, user)
 }
 c.BindJSON(&user)
 err = Models.UpdateUser(&user, id)
 if err != nil {
  c.AbortWithStatus(http.StatusNotFound)
 } else {
  c.JSON(http.StatusOK, user)
 }
}

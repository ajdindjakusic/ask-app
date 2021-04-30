package main

import (
    "go/first-api/Config"
    "go/first-api/Models"
    "go/first-api/Routes"
    "fmt"
    "github.com/jinzhu/gorm"
)

var err error

func main() {
 Config.DB, err = gorm.Open("mysql", Config.DbURL())
if err != nil {
  fmt.Println("Status:", err)
 }
defer Config.DB.Close()
 Config.DB.AutoMigrate(&Models.User{})
 Config.DB.AutoMigrate(&Models.Question{})
 Config.DB.AutoMigrate(&Models.Answer{})
 Config.DB.AutoMigrate(&Models.Like{})
 Config.DB.AutoMigrate(&Models.Dislike{})

r := Routes.SetupRouter()
 r.Run()
}
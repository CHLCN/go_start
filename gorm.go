package main

// import (
// 	"fmt"
// 	"os"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct { // User users
// 	Id      int
// 	Keyword string `gorm:"column:keywords"` //keyword
// 	City    string //city
// }

// func (User) TableName() string {
// 	return "user"
// } // 重新定义表名

// func read(client *gorm.DB, city string) *User {
// 	var user User
// 	user.Id = 853985
// 	err := client.Select("id,city").Where("city=?").First(&user).Error //  id=853958
// 	checkError(err)
// 	return &user
// }

// func main() {
// 	dataSourceName := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&&parseTime=True"
// 	client, err := gorm.Open(mysql.Open(dataSourceName), nil)
// 	checkError(err)
// 	// user := User{
// 	// 	Id:      5858,
// 	// 	Keyword: "golang",
// 	// 	City:    "天津",
// 	// }
// 	// client.Create(user)
// 	client.Model(User{}).Where("id=?", 5858).Update("city", "郑州")
// 	client.Model(User{}).Where("id=?", 5858).Updates(
// 		map[string]interface{}{"city": "郑州", "gender": "true"},
// 	)
// 	client.Model(User{}).Where("id=?", 5858).Delete(User{})
// }

// func checkError(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

// gorm规范：表名转为蛇形复数 UserName->user_names
//			结构体中的成员变量转为蛇形（无需s）

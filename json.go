package main

// import (
// 	"fmt"

// 	"github.com/bytedance/sonic"
// )

// type Student struct {
// 	Name   string
// 	Age    int
// 	Gender bool
// }

// type Class struct {
// 	Id       string
// 	Students []Student
// }

// func main() {
// 	s := Student{"张三", 18, true}
// 	c := Class{
// 		Id:       "1(2)班",
// 		Students: []Student{s, s, s},
// 	}
// 	bytes, err := sonic.Marshal(c)
// 	if err != nil {
// 		fmt.Println("json序列化失败", err)
// 		return
// 	}
// 	str := string(bytes)
// 	fmt.Println(str)

// 	var c2 Class
// 	err = sonic.Unmarshal(bytes, &c2) // 反序列化
// 	if err != nil {
// 		fmt.Println("json反序列化失败", err)
// 		return
// 	}
// 	fmt.Printf("%+v\n", c2)
// }

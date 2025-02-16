package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"testing"

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

// var (
// 	s = Student{"张三", 18, true}
// 	c = Class{
// 		Id:       "1(2)班",
// 		Students: []Student{s, s, s},
// 	}
// )

// func TestJson(t *testing.T) {
// 	bytes, err := json.Marshal(c)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	var c2 Class
// 	err = json.Unmarshal(bytes, &c2)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
// 		t.Fail()
// 	}
// 	fmt.Println("json test success")
// }

// func TestSonic(t *testing.T) {
// 	bytes, err := sonic.Marshal(c)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	var c2 Class
// 	err = sonic.Unmarshal(bytes, &c2)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
// 		t.Fail()
// 	}
// 	fmt.Println("sonic test success")
// }

// 单元/基准测试的文件名必须以_test.go结尾

// func BenchmarkJson(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		bytes, _ := json.Marshal(c)
// 		var c2 Class
// 		json.Unmarshal(bytes, &c2)
// 	}
// }

// func BenchmarkSonic(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		bytes, _ := sonic.Marshal(c)
// 		var c2 Class
// 		sonic.Unmarshal(bytes, &c2)
// 	}
// }

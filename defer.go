package main

// import (
// 	"fmt"
// )

// func foo() int {
// 	a, b := 3, 5
// 	c := a + b
// 	defer fmt.Println("111", c) // defer的注册，在函数退出前倒序执行
// 	fmt.Println(c)
// 	defer fmt.Println("222", c)
// 	defer func() {
// 		fmt.Println("333", c)
// 	}()
// 	c = 100
// 	fmt.Println(c)
// 	return c
// }

// // 如果defer后直接跟go语句，则其中的变量在注册时已计算；如果跟匿名函数，函数体中的变量在执行时才计算
// func main() {
// 	foo()
// }

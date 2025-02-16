package main

// import (
// 	"fmt"
// 	"time"
// )

// func f2() {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("f2 finish")
// }

// func f1() {
// 	defer func() {
// 		err := recover() // 必须放在匿名函数内
// 		if err != nil {
// 			fmt.Printf("发生了panic: %s\n", err)
// 		}
// 	}()
// 	a, b := 3, 0
// 	fmt.Println(a, b)
// 	_ = a / b // 发生panic
// 	fmt.Println("f1 finish")
// }

// func foo() {
// 	fmt.Println("foo finish")
// }

// func main() { // 主协程/main协程
// 	// go foo() // 创建一个子协程以运行foo函数
// 	// wg := sync.WaitGroup{}
// 	// wg.Add(2)

// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	fmt.Println("foo finish1")

// 	// }()
// 	// go func() {
// 	// 	defer wg.Done()
// 	// 	fmt.Println("foo finish2")

// 	// }()
// 	// wg.Wait()
// 	go f1()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("main finish")

// }

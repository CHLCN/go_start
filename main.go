package main // 入口包文件

// import (
// 	"fmt"
// )

// func main() { // 入口函数
// 	// 变量的声明（指定变量的类型）
// 	var a int
// 	var b int
// 	var c int
// 	var age byte
// 	var sex bool
// 	var price float64

// 	// 变量的赋值
// 	a = 10
// 	b = 20
// 	c = a + b

// 	// := 声明和赋值在一行代码里完成
// 	f := 40.0  // 自动推断出f是float64
// 	g := c - a // 自动推断出g是int

// 	var h int = 4 // 声明并赋值
// 	var m = 4     // 声明和赋值一行完成，但是类型是自动推断出来的

// 	fmt.Println(c) // ln 每行末尾追加换行
// 	// 未显式赋值的变量默认为零值，bool为false，byte为0，float64为0.000000
// 	// fmt.Println(age)
// 	// fmt.Println(sex)
// 	// fmt.Println(price)
// 	fmt.Printf("%d %d %t %.2f\n", c, age, sex, price) // f format
// 	// fmt.Println("Hello, world!") // 函数要带上包名
// 	fmt.Printf("%T %T \n", f, g) // 输出变量的类型
// }

// 运行命令 go run main.go

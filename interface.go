package main

// import (
// 	"fmt"
// )

// // 接口是一组行为规范的集合
// type Human interface {
// 	Say(int, int) int
// }

// type None interface{}

// func foo(h Human) {
// 	c := h.Say(3, 6)
// 	fmt.Println("c=", c)
// }
// func main() {
// 	var a Human
// 	t := Tom{}
// 	a = t
// 	// j := Jim{}
// 	// a = j
// 	foo(a)

// 	var b interface{}
// 	j := Jim{}
// 	b = j

// 	var c string
// 	b = c
// 	// 任何类型都可以赋给空接口
// }

// type Tom struct {
// }

// func (Tom) Say(a int, b int) int {
// 	return a - b
// }

// type Jim struct {
// }

// func (Jim) Say(a int, b int) int {
// 	return a + b
// }

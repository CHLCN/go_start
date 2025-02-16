package main

// import (
// 	"errors"
// 	"fmt"
// )

// //	func foo(a int, b int) int {
// //		c := a + b
// //		return c
// //	}
// func foo(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	d := a / b
// 	return d, nil
// }

// func main() {
// 	m, n := 4, 0
// 	// q := foo(m, n)
// 	// fmt.Println(q)
// 	if p, err := foo(m, n); err == nil {
// 		fmt.Println("Result is ", p)
// 	} else {
// 		fmt.Println("Error is", err)
// 	}
// }

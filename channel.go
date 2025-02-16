package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	ch := make(chan int, 100)

// 	// 两个生产者，往channel里写入元素
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 		}
// 	}()

// 	// interface{}表示一种数据类型
// 	// 一个消费者
// 	mc := make(chan struct{}, 0)
// 	go func() {
// 		sum := 0
// 		for {
// 			a, ok := <-ch
// 			if !ok { // channel被关闭，且channel为空，此时ok才为false
// 				break
// 			} else {
// 				sum += a
// 			}
// 			sum += a
// 		}
// 		fmt.Printf("sum=%d\n", sum)
// 		mc <- struct{}{}
// 	}()

// 	wg.Wait()
// 	close(ch)

// 	<-mc
// }

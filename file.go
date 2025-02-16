package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// // Read file
// file, err := os.Open("a.txt")
// if err != nil {
// 	fmt.Println("打开文件失败", err)
// 	return
// }

// defer file.Close() // main函数退出之前执行

// content := make([]byte, 100)
// n, err := file.Read(content)
// if err != nil {
// 	fmt.Println("读文件发生错误", err)
// 	return
// }
// fmt.Println(string(content[0:n]))
// fmt.Println(content)

// Write file
// file, err := os.OpenFile("b.txt",
// 	os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
// if err != nil {
// 	fmt.Println("打开文件失败", err)
// 	return
// }
// defer file.Close()

// content := "hello world\n"
// n, err := file.Write([]byte(content))
// if err != nil {
// 	fmt.Println("写文件失败", err)
// 	return
// } else {
// 	fmt.Printf("成功向文件写入%d个字节\n", n)
// }

// file, err := os.OpenFile("b.txt")
// if err != nil {
// 	fmt.Println("打开文件失败", err)
// 	return
// }
// defer file.Close()

// reader := bufio.NewReader(file)
// for {
// 	line, err := reader.ReadString('\n')
// 	if err != nil {
// 		if err == io.EOF {
// 			fmt.Print(line)
// 			break
// 		} else {
// 			fmt.Println("读文件发生错误", err)
// 			return
// 		}
// 	} else {
// 		fmt.Print(line)
// 	}
// }

// }

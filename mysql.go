package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB

// func initMySQL() (err error) {
// 	// DSN: Data Source Name
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_demo"
// 	// 去初始化全局的db对象而不是新声明一个db变量
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 尝试与数据库建立连接（校验dsn是否正确）
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Printf("connect to db failed, err: %v\n", err)
// 		return
// 	}
// 	// 数值需要根据业务具体情况来确定
// 	db.SetConnMaxLifetime(time.Second * 10)
// 	db.SetMaxOpenConns(200) // 最大连接数
// 	db.SetMaxIdleConns(10)  // 最大空闲连接数
// 	return
// }

// func main() {
// 	if err := initMySQL(); err != nil {
// 		fmt.Printf("connect to db failed, err: %v\n", err)
// 	}
// 	// 做完错误检查之后，确保db不为nil
// 	// Close() 用来释放掉数据库连接相关的资源
// 	defer db.Close() //  要写在上面err判断的下面
// 	fmt.Println("connect to db success")
// }

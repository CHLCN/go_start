package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/redis/go-redis/v9"
// )

// func string(ctx context.Context, client *redis.Client) {
// 	key := "name"
// 	value := "cat"
// 	err := client.Set(ctx, key, value, 1*time.Second).Err()
// 	checkError(err)
// 	client.Expire(ctx, key, 3*time.Second)
// 	time.Sleep(2 * time.Second)

// 	v2, err := client.Get(ctx, key).Result()
// 	checkError(err)
// 	fmt.Println(v2)

// 	client.Del(ctx, key)
// }

// func list(ctx context.Context, client *redis.Client) {
// 	key := "ids"
// 	values := []interface{}{1, 2, 3, "中"}
// 	err := client.RPush(ctx, key, values...).Err()
// 	checkError(err)

// 	v2, err := client.LRange(ctx, key, 0, -1).Result()
// 	checkError(err)
// 	fmt.Println(v2)

// 	client.Del(ctx, key)
// }

// func hashtable(ctx context.Context, client *redis.Client) {
// 	err := client.HSet(ctx, "学生1", "Name", "张三", "Age", 18, "Height", 173.5).Err()
// 	checkError(err)
// 	err = client.HSet(ctx, "学生2", "Name", "李四", "Age", 20, "Height", 180.0).Err()
// 	checkError(err)
// 	age, err := client.HGet(ctx, "学生2", "Age").Result() // Field
// 	checkError(err)
// 	fmt.Println(age)

// 	for field, value := range client.HGetAll(ctx, "学生1").Val() {
// 		fmt.Println(field, value)
// 	}
// 	client.Del(ctx, "学生1")
// 	client.Del(ctx, "学生2")
// }

// func main() {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "127.0.0.1:6379",
// 		Password: "",
// 		DB:       0,
// 	})
// 	ctx := context.TODO()

// 	// string(ctx, client)
// 	// list(ctx, client)
// 	hashtable(ctx, client)

// }

// func checkError(err error) {
// 	if err != nil {
// 		if err == redis.Nil {
// 			fmt.Println("key不存在")
// 		} else {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 	}
// }

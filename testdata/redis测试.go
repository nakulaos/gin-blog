package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func get() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, //连接池大小
	})
	fmt.Println(rdb)

}
func main() {
	get()

	res := rdb.Keys(ctx, "*")
	keys, _ := res.Result()
	fmt.Println(keys)
	//fmt.Println(err)
}

package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client

func RedisInit() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("redis init err : ", pong, err)
	}
	Rdb = rdb
}

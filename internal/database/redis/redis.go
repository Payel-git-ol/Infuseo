package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func init() {
	if err := InitDbRedis; err != nil {
		log.Fatal("Ошибка инициализации REDI")
	}
}

func InitDbRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err, "redis connect error")
	}
	log.Println("Redis connect success!")
}

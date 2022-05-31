package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewRedis() {

}

var ctx = context.Background()
var RC *redis.Client

func ConnectRedis() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redis.Ping(ctx).Result()
	if err != nil {
		log.Panic(err)
	}
	return redis
}

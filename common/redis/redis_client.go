package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RC *redis.Client

func InitRedis() {
    RC = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    _, err := RC.Ping(context.Background()).Result()
    if err != nil {
        log.Fatal(err)
    }
}

func Get(key string) (string, error) {
    return RC.Get(context.Background(), key).Result()
}

func Set(key string, value interface{}) error {
    return RC.Set(context.Background(), key, value, 0).Err()
}